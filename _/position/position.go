package position

import (
	snek "github.com/JMTyler/battlesnake/_"
)

func GetAdjacentTiles(pos snek.Position) map[string]snek.Position {
	// HACK: Should replace these magic numbers with context.Board.Width somehow.
	cells := make(map[string]snek.Position)
	if pos.Y < 10 {
		cells["up"] = snek.Position{pos.X, pos.Y + 1}
	}
	if pos.Y > 0 {
		cells["down"] = snek.Position{pos.X, pos.Y - 1}
	}
	if pos.X > 0 {
		cells["left"] = snek.Position{pos.X - 1, pos.Y}
	}
	if pos.X < 10 {
		cells["right"] = snek.Position{pos.X + 1, pos.Y}
	}
	return cells
}

func IsOutsideBoard(pos snek.Position, board snek.Board) bool {
	return pos.X < 0 || pos.Y < 0 || pos.X >= board.Width || pos.Y >= board.Height
}

func IsDeadly(pos snek.Position, context snek.Context) bool {
	if IsOutsideBoard(pos, context.Board) {
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

func IsRisky(pos snek.Position, context snek.Context) bool {
	for _, snake := range context.Board.Snakes {
		adjacent := GetAdjacentTiles(snake.Head)
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

func IsSafe(pos snek.Position, context snek.Context) bool {
	return !IsDeadly(pos, context) && !IsRisky(pos, context)
}

func Matches(posA snek.Position, posB snek.Position) bool {
	return posA == posB
}

func ToDirection(to snek.Position, from snek.Position) string {
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
