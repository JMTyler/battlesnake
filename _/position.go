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
