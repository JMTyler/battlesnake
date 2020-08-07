package battlesnake

import (
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/traverse"
)

type Board struct {
	Width  int      `json:"width"`
	Height int      `json:"height"`
	Snakes []*Snake `json:"snakes"`
	Food   []*Cell  `json:"food"`

	Graph   traverse.Graph `json:"-"`
	Enemies []*Snake       `json:"-"`
	Cells   [][]*Cell      `json:"-"`
}

func (board *Board) Prepare(ctx *Context) {
	board.loadCells(ctx)

	for _, snake := range board.Snakes {
		snake.Prepare(ctx)
	}

	// Replace food array with cell singletons.
	// TODO: Can we just update the pointer to `food` instead of caring about the index?
	for ix, food := range board.Food {
		board.Food[ix] = board.CellAt(food.X, food.Y)
		board.Food[ix].AddTags("food")
	}

	board.loadEnemies(ctx)

	// everything needs to be prepared/loaded before this, so we can just check the tags on the cells
	board.loadGraph(ctx)
}

func (board *Board) CellAt(x int, y int) *Cell {
	if x < 0 || y < 0 || x >= board.Width || y >= board.Height {
		return nil
	}
	return board.Cells[x][y]
}

func (board *Board) loadEnemies(context *Context) {
	// Remove `You` snake from the snakes array since we only ever want an array of enemies.
	for i, snake := range context.Board.Snakes {
		if snake.ID == context.You.ID {
			context.You = snake
			board.Enemies = append(context.Board.Snakes[:i], context.Board.Snakes[i+1:]...)
			break
		}
	}
}

func (board *Board) loadGraph(context *Context) {
	grid := simple.NewUndirectedGraph()
	board.Graph = grid
	for x := 0; x < board.Width; x++ {
		for y := 0; y < board.Height; y++ {
			node := board.CellAt(x, y)
			if !node.IsDeadly(context) {
				grid.AddNode(node)
			}
		}
	}
	for x := 0; x < board.Width; x++ {
		for y := 0; y < board.Height; y++ {
			node := board.CellAt(x, y)
			if grid.Node(node.ID()) != nil {
				for _, cell := range node.GetAdjacentCells() {
					if grid.Node(cell.ID()) != nil {
						grid.SetEdge(grid.NewEdge(node, cell))
					}
				}
			}
		}
	}
}

func (board *Board) loadCells(ctx *Context) {
	board.Cells = make([][]*Cell, board.Width)
	for x := 0; x < board.Width; x++ {
		board.Cells[x] = make([]*Cell, board.Height)
		for y := 0; y < board.Height; y++ {
			board.Cells[x][y] = &Cell{X: x, Y: y}
			board.Cells[x][y].Prepare(ctx)
		}
	}
}
