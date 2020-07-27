package movement

import (
	snek "github.com/JMTyler/battlesnake/_"
)

func ApproachTarget(target snek.Cell, context snek.Context) string {
	cells := context.You.Head.PathTo(target, context)
	if cells == nil {
		return ""
	}
	return context.You.Head.ToDirection(cells[0])
}

func FindClosestTarget(origin snek.Cell, targets []snek.Cell) snek.Cell {
	if len(targets) == 1 {
		return targets[0]
	}

	if len(targets) == 0 {
		return snek.Cell{}
	}

	distances := origin.GetDistances(targets)

	shortestIndex := 0
	for ix, distance := range distances {
		if distance < distances[shortestIndex] {
			shortestIndex = ix
		}
	}
	return targets[shortestIndex]
}

func FindClosestFood(context snek.Context) snek.Cell {
	return FindClosestTarget(context.You.Head, context.Board.Food)
}
