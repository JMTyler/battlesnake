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

	closestSnake := ctx.You.Head.FindClosest(weaklings)

	if opts.Distance > 0 {
		distanceToSnake := ctx.You.Head.GetDistance(closestSnake)
		if distanceToSnake > opts.Distance {
			return ""
		}
	}

	var prey *snek.Snake
	for _, snake := range ctx.Board.Foes {
		if snake.Head == closestSnake {
			prey = snake
			break
		}
	}

	target := chooseAdjacentCell(prey, ctx, state)
	if target == nil {
		return ""
	}

	return ctx.You.Head.Approach(target)
}

func chooseAdjacentCell(prey *snek.Snake, ctx *snek.Context, state *snek.State) *snek.Cell {
	targetOptions := make(map[string]*snek.Cell)
	for dir, cell := range prey.Head.Neighbours() {
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

	return ctx.You.Head.FindClosest(targets)
}
