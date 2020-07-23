package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/movement"
	"math/rand"
)

type GoCentre struct {
	Width  int
	Height int
}

func (opts GoCentre) Run(context snek.Context, _ *snek.State) string {
	if opts.Width == 0 {
		opts.Width = 1
	}
	if opts.Height == 0 {
		opts.Height = 1
	}

	leftEdge := (context.Board.Width - opts.Width) / 2
	bottomEdge := (context.Board.Height - opts.Height) / 2

	var centreCells []snek.Cell
	for x := leftEdge; x < leftEdge+opts.Width; x++ {
		for y := bottomEdge; y < bottomEdge+opts.Height; y++ {
			cell := snek.Cell{x, y}
			if cell.IsSafe(context) {
				centreCells = append(centreCells, cell)
			}
		}
	}

	if len(centreCells) == 0 {
		return ""
	}

	index := rand.Intn(len(centreCells))
	target := centreCells[index]
	return movement.ApproachTarget(target, context)
}
