package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/movement"
)

type SeekTail struct{}

func (tactic *SeekTail) Run(context snek.Context, state *snek.State) string {
	return movement.ApproachTarget(context.You.Body[len(context.You.Body)-1], context)
}
