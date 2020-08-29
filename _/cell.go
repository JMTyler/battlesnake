package battlesnake

import (
	"fmt"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/traverse"
	"math"
)

// They should already know if anything resides on them, and whether they're deadly or risky, such that we don't have to
// have the same loops checking for the same qualities numerous times each request.
// We can then fix the ID() and GetAdjacentCells() calculations.  Which will allow us to handle the August challenge.
type Cell struct {
	X int `json:"x"`
	Y int `json:"y"`

	board *Board   `json:"-"`
	tags  []string `json:"-"`
}

func (cell *Cell) AddTags(tags ...string) {
	cell.tags = append(cell.tags, tags...)
}

func (cell *Cell) HasTags(tags ...string) bool {
	for _, requiredTag := range tags {
		if !cell.hasTag(requiredTag) {
			return false
		}
	}
	return true
}

func (cell *Cell) hasTag(requiredTag string) bool {
	for _, cellTag := range cell.tags {
		if requiredTag == cellTag {
			return true
		}
	}
	return false
}

func (cell *Cell) Prepare(ctx *Context) {
	cell.board = ctx.Board
	cell.tags = make([]string, 0)
}

type Vector struct {
	Dir struct {
		X string
		Y string
	}
	Weight struct {
		X int
		Y int
	}
}

func (cell *Cell) ID() int64 {
	return int64(cell.X + (cell.Y * cell.board.Width))
}

func (cell *Cell) String() string {
	return fmt.Sprintf("(%d,%d)", cell.X, cell.Y)
}

func (cell *Cell) GetAdjacentCells() map[string]*Cell {
	cells := make(map[string]*Cell)
	if cell.Y < cell.board.Height-1 {
		cells["up"] = cell.Adjacent("up")
	}
	if cell.Y > 0 {
		cells["down"] = cell.Adjacent("down")
	}
	if cell.X > 0 {
		cells["left"] = cell.Adjacent("left")
	}
	if cell.X < cell.board.Width-1 {
		cells["right"] = cell.Adjacent("right")
	}
	return cells
}

func (origin *Cell) Adjacent(dir string) *Cell {
	switch dir {
	case "up":
		return origin.board.CellAt(origin.X, origin.Y+1)
	case "down":
		return origin.board.CellAt(origin.X, origin.Y-1)
	case "left":
		return origin.board.CellAt(origin.X-1, origin.Y)
	case "right":
		return origin.board.CellAt(origin.X+1, origin.Y)
	}
	// TODO: error?
	return nil
}

func (cell *Cell) IsOutsideBoard(board *Board) bool {
	return cell.X < 0 || cell.Y < 0 || cell.X >= board.Width || cell.Y >= board.Height
}

func (cell *Cell) IsDeadly(context *Context) bool {
	// TODO: Might not need this check anymore now that the board only provides cells within the bounds.
	if cell.IsOutsideBoard(context.Board) {
		return true
	}

	if cell.HasTags("you", "head") {
		return false
	}
	if cell.HasTags("you", "body") {
		return true
	}

	if cell.HasTags("enemy", "body") {
		return true
	}
	if cell.HasTags("enemy", "head", "enemy-longer") {
		return true
	}
	if cell.HasTags("enemy", "head", "enemy-equal") {
		return true
	}

	return false
}

func (cell *Cell) IsRisky(context *Context) bool {
	if cell.HasTags("enemy-adjacent") {
		if cell.HasTags("enemy-longer") || cell.HasTags("enemy-equal") {
			return true
		}
	}

	if cell.HasTags("hazard") {
		return true
	}

	return false
}

func (cell *Cell) CanReachTail(context *Context) bool {
	// TODO: Can't seem to path to my tail from right next to it.
	pathToTail := cell.GetFuturePath(context.You.Tail(), context)
	// TODO: We still don't want to follow a path if it funnels us through only one spot, especially next to a head.
	return pathToTail != nil
}

// TODO: Cell should know its own context, and not have to pass it around everywhere.
func (cell *Cell) IsSafe(context *Context) bool {
	return !cell.IsDeadly(context) && !cell.IsRisky(context) && cell.CanReachTail(context)
}

func (a *Cell) Matches(b *Cell) bool {
	return a.X == b.X && a.Y == b.Y
}

func (from *Cell) ToDirection(to *Cell) string {
	x := to.X - from.X
	y := to.Y - from.Y

	if x != 0 && y != 0 {
		// TODO: error: non-lateral
		return ""
	}

	if x > 0 {
		return "right"
	}

	if x < 0 {
		return "left"
	}

	if y > 0 {
		return "up"
	}

	return "down"
}

// TODO: Use pathfinding distance, not direct distance.
func (origin *Cell) GetDistance(target *Cell) int {
	x := math.Abs(float64(target.X - origin.X))
	y := math.Abs(float64(target.Y - origin.Y))
	return int(x + y)
}

func (origin *Cell) GetDistances(targets []*Cell) []int {
	var distances []int
	for _, target := range targets {
		distances = append(distances, origin.GetDistance(target))
	}
	return distances
}

func (origin *Cell) GetVector(target *Cell) *Vector {
	x := target.X - origin.X
	y := target.Y - origin.Y

	xDir := "left"
	if x > 0 {
		xDir = "right"
	}

	yDir := "down"
	if y > 0 {
		yDir = "up"
	}

	return &Vector{
		Dir: struct {
			X string
			Y string
		}{
			xDir,
			yDir,
		},
		Weight: struct {
			X int
			Y int
		}{
			x,
			y,
		},
	}
}

func (origin *Cell) PathTo(target *Cell, graph traverse.Graph) []*Cell {
	shortest, _ := path.AStar(origin, target, graph, nil)
	nodes, _ := shortest.To(target.ID())
	if len(nodes) < 2 {
		return nil
	}

	cells := make([]*Cell, len(nodes)-1)
	for ix, node := range nodes[1:] {
		cells[ix] = node.(*Cell)
	}
	return cells
}

func (origin *Cell) GetRiskyPath(target *Cell, context *Context) []*Cell {
	return origin.PathTo(target, context.Board.RiskyGraph)
}

func (origin *Cell) GetSafePath(target *Cell, context *Context) []*Cell {
	return origin.PathTo(target, context.Board.SafeGraph)
}

func (origin *Cell) GetFuturePath(target *Cell, context *Context) []*Cell {
	return origin.PathTo(target, context.Board.FutureGraph)
}

func (you *Cell) ApproachTarget(target *Cell, context *Context) string {
	cells := you.GetRiskyPath(target, context)
	if cells == nil {
		return ""
	}
	return you.ToDirection(cells[0])
}

func (origin *Cell) FindClosestTarget(targets []*Cell) *Cell {
	if len(targets) == 1 {
		return targets[0]
	}

	if len(targets) == 0 {
		return nil
	}

	distances := origin.GetDistances(targets)

	shortestIndex := 0
	for ix, distance := range distances {
		if distance < distances[shortestIndex] {
			shortestIndex = ix
		}
	}
	return targets[shortestIndex]
}
