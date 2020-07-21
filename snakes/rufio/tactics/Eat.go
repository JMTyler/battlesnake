package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/board"
	"github.com/JMTyler/battlesnake/_/movement"
)

// TODO: Is it realistic for us to figure out how to use Infinity / -Infinity as default options instead of zero values?
// Yes, I think this problem solves itself once we switch to tactics being structs, since the constructor can set defaults.

func Eat(options snek.TacticOptions) func(snek.Context, *snek.State) string {
	return func(context snek.Context, state *snek.State) string {
		if options.Health > 0 {
			if context.You.Health > options.Health {
				return ""
			}
		}

		food := board.FindClosestFood(context)
		if food == (snek.Position{}) {
			return ""
		}

		if options.Distance > 0 {
			distanceToFood := movement.GetDistance(context.You.Head, food)
			if distanceToFood > options.Distance {
				return ""
			}
		}

		return movement.ApproachTarget(food, context)
	}
}
