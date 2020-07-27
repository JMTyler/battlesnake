package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
)

type Aggrieve struct {
	Advantage int
	Distance  int
}

func (opts Aggrieve) Run(context snek.Context, state *snek.State) string {
	if opts.Advantage == 0 {
		opts.Advantage = 1
	}

	var weaklings []snek.Cell
	for _, snake := range context.Board.Enemies {
		if context.You.Length >= snake.Length+opts.Advantage {
			weaklings = append(weaklings, snake.Head)
		}
	}
	if len(weaklings) == 0 {
		return ""
	}

	closestSnake := context.You.Head.FindClosestTarget(weaklings)

	if opts.Distance > 0 {
		distanceToSnake := context.You.Head.GetDistance(closestSnake)
		if distanceToSnake > opts.Distance {
			return ""
		}
	}

	var prey snek.Snake
	// TODO: Would be nice if there's a way to break out of this when we find it?  I can't remember.
	for _, snake := range context.Board.Enemies {
		if snake.Head == closestSnake {
			prey = snake
		}
	}

	target := chooseAdjacentCell(prey, context, state)
	if target == (snek.Cell{}) {
		return ""
	}

	return context.You.Head.ApproachTarget(target, context)
}

func chooseAdjacentCell(prey snek.Snake, context snek.Context, state *snek.State) snek.Cell {
	targetOptions := make(map[string]snek.Cell)
	for dir, cell := range prey.Head.GetAdjacentCells() {
		if cell.IsSafe(context) {
			targetOptions[dir] = cell
		}
	}
	if len(targetOptions) == 0 {
		return snek.Cell{}
	}

	preysLastMove := state.Snakes[prey.ID].Move
	forwardCell, valid := targetOptions[preysLastMove]
	if valid {
		return forwardCell
	}

	var targets []snek.Cell
	for _, cell := range targetOptions {
		targets = append(targets, cell)
	}

	return context.You.Head.FindClosestTarget(targets)
}
