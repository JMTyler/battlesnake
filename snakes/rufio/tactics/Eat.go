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

func (opts Eat) Run(context snek.Context, _ *snek.State) string {
	if opts.Health > 0 {
		if context.You.Health > opts.Health {
			return ""
		}
	}

	food := context.You.Head.FindClosestTarget(context.Board.Food)
	if *food == (snek.Cell{}) {
		return ""
	}

	if opts.Distance > 0 {
		distanceToFood := context.You.Head.GetDistance(food)
		if distanceToFood > opts.Distance {
			return ""
		}
	}

	return context.You.Head.ApproachTarget(food, context)
}
