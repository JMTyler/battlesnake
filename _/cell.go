package battlesnake

import (
	"fmt"
	"github.com/JMTyler/battlesnake/_/utils"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/traverse"
	"math"
)

type Cell struct {
	X int `json:"x"`
	Y int `json:"y"`

	board *Board   `json:"-"`
	tags  []string `json:"-"`
}

// TODO: Merge with Cell when it becomes an interface.  Only difference is Position isn't limited to board boundaries.
type Position struct {
	X int
	Y int
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
	Weight *Position
}

func (cell *Cell) ID() int64 {
	return int64(cell.X + (cell.Y * cell.board.Width))
}

func (cell *Cell) String() string {
	return fmt.Sprintf("(%d,%d)", cell.X, cell.Y)
}

func (cell *Cell) Neighbours() map[string]*Cell {
	cells := make(map[string]*Cell)
	if cell.Y < cell.board.Height-1 {
		cells["up"] = cell.Neighbour("up")
	}
	if cell.Y > 0 {
		cells["down"] = cell.Neighbour("down")
	}
	if cell.X > 0 {
		cells["left"] = cell.Neighbour("left")
	}
	if cell.X < cell.board.Width-1 {
		cells["right"] = cell.Neighbour("right")
	}
	return cells
}

func (origin *Cell) Neighbour(dir string) *Cell {
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

func (cell *Cell) IsOutsideBoard() bool {
	return cell.X < 0 || cell.Y < 0 || cell.X >= cell.board.Width || cell.Y >= cell.board.Height
}

func (cell *Cell) IsDeadly() bool {
	// TODO: Might not need this check anymore now that the board only provides cells within the bounds.
	if cell.IsOutsideBoard() {
		return true
	}

	if cell.HasTags("you", "head") {
		return false
	}
	if cell.HasTags("you", "body") {
		return true
	}

	if cell.HasTags("friend", "head") {
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

	if cell.HasTags("food", "friend-adjacent") {
		return true
	}
	if cell.HasTags("food", "enemy-adjacent", "enemy-longer") {
		return true
	}
	if cell.HasTags("food", "enemy-adjacent", "enemy-equal") {
		return true
	}

	return false
}

func (cell *Cell) IsRisky() bool {
	if cell.HasTags("friend-adjacent") {
		return true
	}
	if cell.HasTags("enemy-adjacent", "enemy-longer") {
		return true
	}
	if cell.HasTags("enemy-adjacent", "enemy-equal") {
		return true
	}

	if cell.HasTags("hazard") {
		return true
	}
	if cell.HasTags("edge") {
		return true
	}

	return false
}

func (cell *Cell) IsEdge() bool {
	return cell.HasTags("edge")
}

func (origin *Cell) CanReachTail(snake *Snake) bool {
	if origin == snake.Tail() {
		return true
	}

	for turnsAway := 1; turnsAway < len(snake.FullBody); turnsAway++ {
		ix := len(snake.FullBody) - turnsAway
		pathToTail := origin.GetTheoreticalPath(snake.FullBody[ix])
		// This is actually one extra turn away, just in case we run into food.
		if pathToTail != nil && len(pathToTail) > turnsAway {
			return true
		}
	}

	return false

	// TODO: This doesn't really work. Doesn't handle corners well, and might be too simple-minded anyway.
	// TODO: What if we add a "funnel" tag to cells and are scared of *many* of them or avoid funnels when seeking tail?
	//for _, cell := range pathToTail {
	//	neighbours := cell.Neighbours()
	//	if len(neighbours) == 2 {
	//		// This cell is a funnel - its only neighbours are the entry and exit.  It's too risky.
	//		return false
	//	}
	//}
	//
	//return true
}

// TODO: Cell should know its own context, and not have to pass it around everywhere.
func (cell *Cell) IsSafe(ctx *Context) bool {
	return !cell.IsDeadly() && !cell.IsRisky() && cell.CanReachTail(ctx.You)
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

func (origin *Cell) GetDistance(target *Cell) int {
	path := origin.PathTo(target)
	if path == nil {
		return math.MaxInt32
	}
	return len(path)
}

func (origin *Cell) GetDistances(targets []*Cell) []int {
	var distances []int
	for _, target := range targets {
		distances = append(distances, origin.GetDistance(target))
	}
	return distances
}

// TODO: Implement methods for various common matrix operations.
func (origin *Cell) Translate(delta *Position) *Cell {
	x := utils.Clamp(origin.X+delta.X, 0, origin.board.Width-1)
	y := utils.Clamp(origin.Y+delta.Y, 0, origin.board.Height-1)
	return origin.board.CellAt(x, y)
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
		Weight: &Position{
			x,
			y,
		},
	}
}

func (origin *Cell) getPath(target *Cell, graph traverse.Graph) []*Cell {
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

func (origin *Cell) PathTo(target *Cell) []*Cell {
	if superSafePath := origin.getPath(target, origin.board.SuperSafeGraph); superSafePath != nil {
		return superSafePath
	}
	if safePath := origin.getPath(target, origin.board.SafeGraph); safePath != nil {
		return safePath
	}
	return origin.getPath(target, origin.board.RiskyGraph)
}

func (origin *Cell) GetTheoreticalPath(target *Cell) []*Cell {
	// Clone the graph so we can add the target to it, just this once.
	graph := *origin.board.FutureGraph

	if graph.Node(target.ID()) == nil {
		graph.AddNode(target)
		for _, neighbour := range target.Neighbours() {
			if graph.Node(neighbour.ID()) != nil {
				graph.SetEdge(graph.NewEdge(target, neighbour))
			}
		}
	}

	return origin.getPath(target, &graph)
}

func (origin *Cell) Approach(target *Cell) string {
	path := origin.PathTo(target)
	if path == nil {
		return ""
	}
	return origin.ToDirection(path[0])
}

func (origin *Cell) FindClosest(targets []*Cell) *Cell {
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

func FilterCells(cells []*Cell, predicate func(*Cell) bool) []*Cell {
	result := make([]*Cell, 0)
	for _, cell := range cells {
		if predicate(cell) {
			result = append(result, cell)
		}
	}
	return result
}
