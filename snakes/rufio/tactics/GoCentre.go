package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
	"math/rand"
)

type GoCentre struct {
	Width  int
	Height int
}

func (opts GoCentre) Run(ctx *snek.Context, _ *snek.State) string {
	if opts.Width == 0 {
		opts.Width = 1
	}
	if opts.Height == 0 {
		opts.Height = 1
	}

	leftEdge := (ctx.Board.Width - opts.Width) / 2
	bottomEdge := (ctx.Board.Height - opts.Height) / 2

	// If you're already within the centre area, move on.
	you := ctx.You.Head
	if you.X >= leftEdge && you.X < leftEdge+opts.Width {
		if you.Y >= bottomEdge && you.Y < bottomEdge+opts.Height {
			return ""
		}
	}

	var centreCells []*snek.Cell
	for x := leftEdge; x < leftEdge+opts.Width; x++ {
		for y := bottomEdge; y < bottomEdge+opts.Height; y++ {
			cell := ctx.Board.CellAt(x, y)
			if cell.IsSafe(ctx) {
				centreCells = append(centreCells, cell)
			}
		}
	}

	if len(centreCells) == 0 {
		return ""
	}

	index := rand.Intn(len(centreCells))
	target := centreCells[index]
	return ctx.You.Head.ApproachTarget(target)
}
