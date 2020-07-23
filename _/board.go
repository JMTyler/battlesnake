package battlesnake

import (
	"gonum.org/v1/gonum/graph/traverse"
)

type Board struct {
	Width  int            `json:"width"`
	Height int            `json:"height"`
	Snakes []Snake        `json:"snakes"`
	Food   []Cell         `json:"food"`
	Graph  traverse.Graph `json:", ignore"`

	Enemies []Snake
}

func (board Board) LoadEnemies(context Context) {
	// Remove `You` snake from the snakes array since we only ever want an array of enemies.
	for i, snake := range context.Board.Snakes {
		if snake.ID == context.You.ID {
			board.Enemies = append(context.Board.Snakes[:i], context.Board.Snakes[i+1:]...)
			break
		}
	}
}
