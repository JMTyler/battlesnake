package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/movement"
	"github.com/JMTyler/battlesnake/_/position"
	"math/rand"
)

const centreWidth = 3
const centreHeight = 3

type GoCentre struct{ Name string }

func (t *GoCentre) Description() string {
	return t.Name
}

func (tactic *GoCentre) Run(context snek.Context, state *snek.State) string {
	leftEdge := (context.Board.Width - centreWidth) / 2
	bottomEdge := (context.Board.Height - centreHeight) / 2

	var centreCells []snek.Position
	for x := leftEdge; x < leftEdge+centreWidth; x++ {
		for y := bottomEdge; y < bottomEdge+centreHeight; y++ {
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
