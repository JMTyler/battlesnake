package battlesnake

import (
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/traverse"
)

type Board struct {
	Width   int      `json:"width"`
	Height  int      `json:"height"`
	Snakes  []*Snake `json:"snakes"`
	Food    []*Cell  `json:"food"`
	Hazards []*Cell  `json:"hazards"`

	RiskyGraph  traverse.Graph `json:"-"`
	SafeGraph   traverse.Graph `json:"-"`
	FutureGraph traverse.Graph `json:"-"`

	Foes  []*Snake  `json:"-"`
	Cells [][]*Cell `json:"-"`
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

	// Replace hazards array with cell singletons.
	if board.Hazards == nil {
		board.Hazards = make([]*Cell, 0)
	}
	for ix, hazard := range board.Hazards {
		board.Hazards[ix] = board.CellAt(hazard.X, hazard.Y)
		board.Hazards[ix].AddTags("hazard")
	}

	board.loadEnemies(ctx)

	// everything needs to be prepared/loaded before this, so we can just check the tags on the cells
	board.loadGraphs(ctx)
}

func (board *Board) CellAt(x int, y int) *Cell {
	if x < 0 || y < 0 || x >= board.Width || y >= board.Height {
		return nil
	}
	return board.Cells[x][y]
}

func (board *Board) loadEnemies(ctx *Context) {
	// Remove `You` snake from the snakes array since we only ever want an array of enemies.
	for i, snake := range ctx.Board.Snakes {
		if snake.ID == ctx.You.ID {
			ctx.You = snake
			board.Foes = append(ctx.Board.Snakes[:i], ctx.Board.Snakes[i+1:]...)
			break
		}
	}
}

func (board *Board) loadGraphs(ctx *Context) {
	riskyGraph := simple.NewUndirectedGraph()
	safeGraph := simple.NewUndirectedGraph()
	noHeadsGraph := simple.NewUndirectedGraph()

	board.RiskyGraph = riskyGraph
	board.SafeGraph = safeGraph
	board.FutureGraph = noHeadsGraph

	for x := 0; x < board.Width; x++ {
		for y := 0; y < board.Height; y++ {
			node := board.CellAt(x, y)
			if !node.IsDeadly(ctx) {
				riskyGraph.AddNode(node)
				if !node.IsRisky(ctx) {
					safeGraph.AddNode(node)
					if !node.HasTags("head") {
						noHeadsGraph.AddNode(node)
					}
				}
			}
			if node.HasTags("tail") && noHeadsGraph.Node(node.ID()) == nil {
				noHeadsGraph.AddNode(node)
			}
		}
	}
	for x := 0; x < board.Width; x++ {
		for y := 0; y < board.Height; y++ {
			node := board.CellAt(x, y)
			if riskyGraph.Node(node.ID()) != nil {
				for _, cell := range node.GetAdjacentCells() {
					if riskyGraph.Node(cell.ID()) != nil {
						riskyGraph.SetEdge(riskyGraph.NewEdge(node, cell))
					}
				}
			}
			if safeGraph.Node(node.ID()) != nil {
				for _, cell := range node.GetAdjacentCells() {
					if safeGraph.Node(cell.ID()) != nil {
						safeGraph.SetEdge(safeGraph.NewEdge(node, cell))
					}
				}
			}
			if noHeadsGraph.Node(node.ID()) != nil {
				for _, cell := range node.GetAdjacentCells() {
					if noHeadsGraph.Node(cell.ID()) != nil {
						noHeadsGraph.SetEdge(noHeadsGraph.NewEdge(node, cell))
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
