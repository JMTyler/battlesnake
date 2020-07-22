package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/movement"
	"github.com/JMTyler/battlesnake/_/position"
	"math/rand"
)

type GoCentre struct {
	Width  int
	Height int
}

func (tactic GoCentre) Run(context snek.Context, state *snek.State) string {
	if tactic.Width == 0 {
		tactic.Width = 1
	}
	if tactic.Height == 0 {
		tactic.Height = 1
	}

	leftEdge := (context.Board.Width - tactic.Width) / 2
	bottomEdge := (context.Board.Height - tactic.Height) / 2

	var centreCells []snek.Position
	for x := leftEdge; x < leftEdge+tactic.Width; x++ {
		for y := bottomEdge; y < bottomEdge+tactic.Height; y++ {
			pos := snek.Position{X: x, Y: y}
			if position.IsSafe(pos, context) {
				centreCells = append(centreCells, pos)
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
