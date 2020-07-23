package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/movement"
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

	food := movement.FindClosestFood(context)
	if food == (snek.Position{}) {
		return ""
	}

	if opts.Distance > 0 {
		distanceToFood := movement.GetDistance(context.You.Head, food)
		if distanceToFood > opts.Distance {
			return ""
		}
	}

	return movement.ApproachTarget(food, context)
}
