package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/movement"
)

type SeekTail struct{ Name string }

func (t *SeekTail) Description() string {
	return t.Name
}

func (tactic *SeekTail) Run(context snek.Context, state *snek.State) string {
	return movement.ApproachTarget(context.You.Body[len(context.You.Body)-1], context)
}
