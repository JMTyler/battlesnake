package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/movement"
)

type Aggrieve struct {
	Advantage int
	Distance  int
}

func (opts Aggrieve) Run(context snek.Context, state *snek.State) string {
	if opts.Advantage == 0 {
		opts.Advantage = 1
	}

	var weaklings []snek.Position
	for _, snake := range context.Board.Enemies {
		if context.You.Length >= snake.Length+opts.Advantage {
			weaklings = append(weaklings, snake.Head)
		}
	}
	if len(weaklings) == 0 {
		return ""
	}

	closestSnake := movement.FindClosestTarget(context.You.Head, weaklings)

	if opts.Distance > 0 {
		distanceToSnake := movement.GetDistance(context.You.Head, closestSnake)
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
	if target == (snek.Position{}) {
		return ""
	}

	return movement.ApproachTarget(target, context)
}

func chooseAdjacentCell(prey snek.Snake, context snek.Context, state *snek.State) snek.Position {
	targetOptions := make(map[string]snek.Position)
	for dir, pos := range prey.Head.GetAdjacentCells() {
		if pos.IsSafe(context) {
			targetOptions[dir] = pos
		}
	}
	if len(targetOptions) == 0 {
		return snek.Position{}
	}

	preysLastMove := state.Snakes[prey.ID].Move
	forwardCell, valid := targetOptions[preysLastMove]
	if valid {
		return forwardCell
	}

	var targets []snek.Position
	for _, pos := range targetOptions {
		targets = append(targets, pos)
	}

	return movement.FindClosestTarget(context.You.Head, targets)
}
