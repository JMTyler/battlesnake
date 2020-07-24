package battlesnake

import (
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/traverse"
)

type Board struct {
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Snakes []Snake `json:"snakes"`
	Food   []Cell  `json:"food"`

	Graph   traverse.Graph `json:"-"`
	Enemies []Snake        `json:"-"`
}

func (board *Board) LoadEnemies(context Context) {
	// Remove `You` snake from the snakes array since we only ever want an array of enemies.
	for i, snake := range context.Board.Snakes {
		if snake.ID == context.You.ID {
			board.Enemies = append(context.Board.Snakes[:i], context.Board.Snakes[i+1:]...)
			break
		}
	}
}

func (board *Board) LoadGraph(context Context) {
	grid := simple.NewUndirectedGraph()
	for x := 0; x < board.Width; x++ {
		for y := 0; y < board.Height; y++ {
			node := Cell{x, y}
			if !node.IsDeadly(context) {
				grid.AddNode(node)
			}
		}
	}
	for x := 0; x < board.Width; x++ {
		for y := 0; y < board.Height; y++ {
			node := Cell{x, y}
			if grid.Node(node.ID()) != nil {
				for _, cell := range node.GetAdjacentCells() {
					if grid.Node(cell.ID()) != nil {
						grid.SetEdge(grid.NewEdge(node, cell))
					}
				}
			}
		}
	}
	board.Graph = grid
}
