package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/board"
	"github.com/JMTyler/battlesnake/_/movement"
)

type Eat struct {
	Name string

	Health   int
	Distance int
}

func (t *Eat) Description() string {
	return t.Name
}

// TODO: Is it realistic for us to figure out how to use Infinity / -Infinity as default tactic instead of zero values?
// Yes, I think this problem solves itself once we switch to tactics being structs, since the constructor can set defaults.

func (tactic *Eat) Run(context snek.Context, state *snek.State) string {
	if tactic.Health > 0 {
		if context.You.Health > tactic.Health {
			return ""
		}
	}

	food := board.FindClosestFood(context)
	if food == (snek.Position{}) {
		return ""
	}

	if tactic.Distance > 0 {
		distanceToFood := movement.GetDistance(context.You.Head, food)
		if distanceToFood > tactic.Distance {
			return ""
		}
	}

	return movement.ApproachTarget(food, context)
}
