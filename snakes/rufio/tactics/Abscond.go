package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/movement"
)

type Abscond struct {
	Disadvantage int
	Distance     int
}

func (tactic Abscond) Run(context snek.Context, state *snek.State) string {

	if tactic.Disadvantage == 0 {
		tactic.Disadvantage = 1
	}

	var predators []snek.Position
	for _, snake := range context.Board.Snakes {
		if context.You.Length <= snake.Length-tactic.Disadvantage {
			predators = append(predators, snake.Head)
		}
	}

	if len(predators) == 0 {
		return ""
	}

	predator := movement.FindClosestTarget(context.You.Head, predators)

	if tactic.Distance > 0 {
		distanceToPredator := movement.GetDistance(context.You.Head, predator)
		if distanceToPredator > tactic.Distance {
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

func clamp(val int, min int, max int) int {
	if val <= min {
		return min
	}
	if val >= max {
		return max
	}
	return val
}
