package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/movement"
	"github.com/JMTyler/battlesnake/_/position"
)

func chooseAdjacentCell(prey snek.Snake, context snek.Context, state snek.State) snek.Position {
	targetOptions := make(map[string]snek.Position)
	for dir, pos := range position.GetAdjacentTiles(prey.Head) {
		if position.IsSafe(pos, context) {
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

func Aggrieve(options snek.TacticOptions) func(snek.Context, snek.State) string {
	if options.Advantage == 0 {
		options.Advantage = 1
	}

	return func(context snek.Context, state snek.State) string {
		var weaklings []snek.Position
		for _, snake := range context.Board.Snakes {
			if context.You.Length >= snake.Length+options.Advantage {
				weaklings = append(weaklings, snake.Head)
			}
		}
		if len(weaklings) == 0 {
			return ""
		}

		closestSnake := movement.FindClosestTarget(context.You.Head, weaklings)

		if options.Distance > 0 {
			distanceToSnake := movement.GetDistance(context.You.Head, closestSnake)
			if distanceToSnake > options.Distance {
				return ""
			}
		}

		var prey snek.Snake
		// TODO: Would be nice if there's a way to break out of this when we find it?  I can't remember.
		for _, snake := range context.Board.Snakes {
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
}