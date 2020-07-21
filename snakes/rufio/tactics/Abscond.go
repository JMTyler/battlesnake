package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/movement"
)

func clamp(val int, min int, max int) int {
	if val <= min {
		return min
	}
	if val >= max {
		return max
	}
	return val
}

func Abscond(options snek.TacticOptions) func(snek.Context, *snek.State) string {
	if options.Disadvantage == 0 {
		options.Disadvantage = 1
	}

	return func(context snek.Context, state *snek.State) string {
		var predators []snek.Position
		for _, snake := range context.Board.Snakes {
			if context.You.Length <= snake.Length-options.Disadvantage {
				predators = append(predators, snake.Head)
			}
		}

		if len(predators) == 0 {
			return ""
		}

		predator := movement.FindClosestTarget(context.You.Head, predators)

		if options.Distance > 0 {
			distanceToPredator := movement.GetDistance(context.You.Head, predator)
			if distanceToPredator > options.Distance {
				return ""
			}
		}

		vector := movement.GetVector(context.You.Head, predator)
		xEscapeVector := -1 * vector.Weight.X
		yEscapeVector := -1 * vector.Weight.Y
		escapeTarget := snek.Position{
			X: clamp(xEscapeVector+context.You.Head.X, 0, context.Board.Width-1),
			Y: clamp(yEscapeVector+context.You.Head.Y, 0, context.Board.Height-1),
		}

		return movement.ApproachTarget(escapeTarget, context)
	}
}
