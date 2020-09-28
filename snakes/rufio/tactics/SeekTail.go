package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
)

type SeekTail struct{}

func (_ SeekTail) Run(ctx *snek.Context, _ *snek.State) string {
	move := ctx.You.Head.Approach(ctx.You.Tail())
	if move != "" {
		return move
	}

	// Start 2 turns away since we already tried with the actual tail, effectively 1 turn away.
	for turnsAway := 2; turnsAway < len(ctx.You.FullBody); turnsAway++ {
		ix := len(ctx.You.FullBody) - turnsAway
		pathToTail := ctx.You.Head.GetTheoreticalPath(ctx.You.FullBody[ix])
		if pathToTail != nil && len(pathToTail) >= turnsAway {
			return ctx.You.Head.ToDirection(pathToTail[0])
		}
	}

	return ""
}
