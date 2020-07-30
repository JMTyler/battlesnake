package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
)

type Abscond struct {
	Disadvantage int
	Distance     int
}

func (opts Abscond) Run(context snek.Context, _ *snek.State) string {
	if opts.Disadvantage == 0 {
		opts.Disadvantage = 1
	}

	var predators []*snek.Cell
	for _, snake := range context.Board.Enemies {
		if context.You.Length <= snake.Length-opts.Disadvantage {
			predators = append(predators, snake.Head)
		}
	}

	if len(predators) == 0 {
		return ""
	}

	predator := context.You.Head.FindClosestTarget(predators)

	if opts.Distance > 0 {
		distanceToPredator := context.You.Head.GetDistance(predator)
		if distanceToPredator > opts.Distance {
			return ""
		}
	}

	vector := context.You.Head.GetVector(predator)
	xEscapeVector := -1 * vector.Weight.X
	yEscapeVector := -1 * vector.Weight.Y
	escapeTarget := context.Board.CellAt(
		clamp(xEscapeVector+context.You.Head.X, 0, context.Board.Width-1),
		clamp(yEscapeVector+context.You.Head.Y, 0, context.Board.Height-1),
	)

	return context.You.Head.ApproachTarget(escapeTarget, context)
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
