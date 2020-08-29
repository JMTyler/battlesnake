package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
)

type SeekTail struct{}

func (_ SeekTail) Run(ctx *snek.Context, _ *snek.State) string {
	// TODO: If can't reach tail, try to approach the closest part to it instead.
	return ctx.You.Head.ApproachTarget(ctx.You.Tail(), ctx)
}
