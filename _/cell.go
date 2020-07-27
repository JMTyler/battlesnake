package battlesnake

import (
	"fmt"
	"math"
)

type Cell struct {
	X int `json:"x"`
	Y int `json:"y"`
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

func (cell Cell) ID() int64 {
	// HACK: Should replace this magic number with context.Board.Width somehow.
	return int64(cell.X + (cell.Y * 11))
}

func (cell Cell) String() string {
	return fmt.Sprintf("(%d,%d)", cell.X, cell.Y)
}

func (cell Cell) GetAdjacentCells() map[string]Cell {
	// HACK: Should replace these magic numbers with context.Board.Width somehow.
	cells := make(map[string]Cell)
	if cell.Y < 10 {
		cells["up"] = cell.Adjacent("up")
	}
	if cell.Y > 0 {
		cells["down"] = cell.Adjacent("down")
	}
	if cell.X > 0 {
		cells["left"] = cell.Adjacent("left")
	}
	if cell.X < 10 {
		cells["right"] = cell.Adjacent("right")
	}
	return cells
}

func (cell Cell) Adjacent(dir string) Cell {
	switch dir {
	case "up":
		return Cell{cell.X, cell.Y + 1}
	case "down":
		return Cell{cell.X, cell.Y - 1}
	case "left":
		return Cell{cell.X - 1, cell.Y}
	case "right":
		return Cell{cell.X + 1, cell.Y}
	}
	// TODO: error?
	return Cell{}
}

func (cell Cell) IsOutsideBoard(board Board) bool {
	return cell.X < 0 || cell.Y < 0 || cell.X >= board.Width || cell.Y >= board.Height
}

func (cell Cell) IsDeadly(context Context) bool {
	if cell.IsOutsideBoard(context.Board) {
		return true
	}

	for _, bodyPart := range context.You.Body()[1:] {
		if cell == bodyPart {
			// self collision
			return true
		}
	}

	for _, snake := range context.Board.Enemies {
		collision := false
		for _, bodyPart := range snake.Body() {
			if cell == bodyPart {
				collision = true
			}
		}

		if collision {
			if cell != snake.Head || context.You.Length <= snake.Length {
				return true
			}
		}
	}

	return false
}

func (cell Cell) IsRisky(context Context) bool {
	for _, snake := range context.Board.Enemies {
		// TODO: Should we use range and iterate over the adjacent map instead?
		gettinSpicy := cell == snake.Head.Adjacent("left") ||
			cell == snake.Head.Adjacent("right") ||
			cell == snake.Head.Adjacent("up") ||
			cell == snake.Head.Adjacent("down")

		if gettinSpicy {
			if context.You.Length <= snake.Length {
				return true
			}
		}
	}

	return false
}

func (cell Cell) IsSafe(context Context) bool {
	return !cell.IsDeadly(context) && !cell.IsRisky(context)
}

func (cellA Cell) Matches(cellB Cell) bool {
	return cellA == cellB
}

func (from Cell) ToDirection(to Cell) string {
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
func (origin Cell) GetDistance(target Cell) int {
	x := math.Abs(float64(target.X - origin.X))
	y := math.Abs(float64(target.Y - origin.Y))
	return int(x + y)
}

func (origin Cell) GetDistances(targets []Cell) []int {
	var distances []int
	for _, target := range targets {
		distances = append(distances, origin.GetDistance(target))
	}
	return distances
}

func (origin Cell) GetVector(target Cell) Vector {
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

	return Vector{
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
