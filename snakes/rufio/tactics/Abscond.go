package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
)

type Abscond struct {
	Disadvantage int
	Distance     int
}

func (opts Abscond) Run(ctx *snek.Context, _ *snek.State) string {
	if opts.Disadvantage == 0 {
		opts.Disadvantage = 1
	}

	var predators []*snek.Cell
	for _, snake := range ctx.Board.Foes {
		if ctx.You.Length <= snake.Length-opts.Disadvantage {
			predators = append(predators, snake.Head)
		}
	}

	if len(predators) == 0 {
		return ""
	}

	predator := ctx.You.Head.FindClosestTarget(predators)

	if opts.Distance > 0 {
		distanceToPredator := ctx.You.Head.GetDistance(predator)
		if distanceToPredator > opts.Distance {
			return ""
		}
	}

	vector := ctx.You.Head.GetVector(predator)
	xEscapeVector := -1 * vector.Weight.X
	yEscapeVector := -1 * vector.Weight.Y
	escapeTarget := ctx.Board.CellAt(
		clamp(xEscapeVector+ctx.You.Head.X, 0, ctx.Board.Width-1),
		clamp(yEscapeVector+ctx.You.Head.Y, 0, ctx.Board.Height-1),
	)

	/* Sometimes escapeTarget gets clamped such that there's nothing to do (clamped to your current position if you're
	   next to the wall, for example), causing this tactic to get skipped.  This is a problem when there are still
	   valid ways to abscond and you really should be taking them.
	   TODO: Make the escape target/vector smarter so you still abscond when you need to abscond. */
	return ctx.You.Head.ApproachTarget(escapeTarget, ctx)
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
