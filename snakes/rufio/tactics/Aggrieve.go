package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
)

type Aggrieve struct {
	Advantage int
	Distance  int
}

func (opts Aggrieve) Run(ctx *snek.Context, state *snek.State) string {
	if opts.Advantage == 0 {
		opts.Advantage = 1
	}

	var weaklings []*snek.Cell
	for _, snake := range ctx.Board.Foes {
		if ctx.You.Length >= snake.Length+opts.Advantage {
			weaklings = append(weaklings, snake.Head)
		}
	}
	if len(weaklings) == 0 {
		return ""
	}

	closestSnake := ctx.You.Head.FindClosestTarget(weaklings)

	if opts.Distance > 0 {
		distanceToSnake := ctx.You.Head.GetDistance(closestSnake)
		if distanceToSnake > opts.Distance {
			return ""
		}
	}

	var prey *snek.Snake
	// TODO: Would be nice if there's a way to break out of this when we find it?  I can't remember.
	for _, snake := range ctx.Board.Foes {
		if snake.Head == closestSnake {
			prey = snake
		}
	}

	target := chooseAdjacentCell(prey, ctx, state)
	if target == nil {
		return ""
	}

	return ctx.You.Head.ApproachTarget(target, ctx)
}

func chooseAdjacentCell(prey *snek.Snake, ctx *snek.Context, state *snek.State) *snek.Cell {
	targetOptions := make(map[string]*snek.Cell)
	for dir, cell := range prey.Head.GetAdjacentCells() {
		if cell.IsSafe(ctx) {
			targetOptions[dir] = cell
		}
	}
	if len(targetOptions) == 0 {
		return nil
	}

	preysLastMove := state.Snakes[prey.ID].Move
	forwardCell, valid := targetOptions[preysLastMove]
	if valid {
		return forwardCell
	}

	var targets []*snek.Cell
	for _, cell := range targetOptions {
		targets = append(targets, cell)
	}

	return ctx.You.Head.FindClosestTarget(targets)
}
