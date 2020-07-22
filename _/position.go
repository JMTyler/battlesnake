package battlesnake

import (
	"fmt"
)

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (pos Position) ID() int64 {
	// HACK: Should replace this magic number with context.Board.Width somehow.
	return int64(pos.X + (pos.Y * 11))
}

func (pos Position) String() string {
	return fmt.Sprintf("(%d,%d)", pos.X, pos.Y)
}

func (pos Position) GetAdjacentTiles() map[string]Position {
	// HACK: Should replace these magic numbers with context.Board.Width somehow.
	cells := make(map[string]Position)
	if pos.Y < 10 {
		cells["up"] = Position{pos.X, pos.Y + 1}
	}
	if pos.Y > 0 {
		cells["down"] = Position{pos.X, pos.Y - 1}
	}
	if pos.X > 0 {
		cells["left"] = Position{pos.X - 1, pos.Y}
	}
	if pos.X < 10 {
		cells["right"] = Position{pos.X + 1, pos.Y}
	}
	return cells
}

func (pos Position) IsOutsideBoard(board Board) bool {
	return pos.X < 0 || pos.Y < 0 || pos.X >= board.Width || pos.Y >= board.Height
}

func (pos Position) IsDeadly(context Context) bool {
	if pos.IsOutsideBoard(context.Board) {
		return true
	}

	for _, cell := range context.You.Body[1 : len(context.You.Body)-1] {
		if cell == pos {
			// self collision
			return true
		}
	}

	for _, snake := range context.Board.Snakes {
		collision := false
		for _, cell := range snake.Body[:len(snake.Body)-1] {
			if cell == pos {
				collision = true
			}
		}

		if collision {
			if pos != snake.Head || context.You.Length <= snake.Length {
				return true
			}
		}
	}

	return false
}

func (pos Position) IsRisky(context Context) bool {
	for _, snake := range context.Board.Snakes {
		adjacent := snake.Head.GetAdjacentTiles()
		// TODO: should be able to use range to iterate over the adjacent map instead
		gettinSpicy := pos == adjacent["left"] ||
			pos == adjacent["right"] ||
			pos == adjacent["up"] ||
			pos == adjacent["down"]

		if gettinSpicy {
			if context.You.Length <= snake.Length {
				return true
			}
		}
	}

	return false
}

func (pos Position) IsSafe(context Context) bool {
	return !pos.IsDeadly(context) && !pos.IsRisky(context)
}

func (posA Position) Matches(posB Position) bool {
	return posA == posB
}

func (from Position) ToDirection(to Position) string {
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
