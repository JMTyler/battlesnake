package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
)

type Eat struct {
	Health   int
	Distance int
}

// TODO: Is it realistic for us to figure out how to use Infinity / -Infinity as default tactic instead of zero values?
// Yes, I think this problem solves itself once we switch to tactics being structs, since the constructor can set defaults.

func (opts Eat) Run(ctx *snek.Context, _ *snek.State) string {
	if opts.Health > 0 {
		if ctx.You.Health > opts.Health {
			return ""
		}
	}

	food := ctx.You.Head.FindClosestTarget(ctx.Board.Food)
	if food == nil {
		return ""
	}

	if food.HasTags("hazard") {
		return ""
	}

	if opts.Distance > 0 {
		distanceToFood := ctx.You.Head.GetDistance(food)
		if distanceToFood > opts.Distance {
			return ""
		}
	}

	return ctx.You.Head.ApproachTarget(food, ctx)
}
