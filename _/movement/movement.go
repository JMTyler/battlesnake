package movement

import (
	snek "github.com/JMTyler/battlesnake/_"
	"gonum.org/v1/gonum/graph/path"
)

func ApproachTarget(target snek.Cell, context snek.Context) string {
	shortest, _ := path.AStar(context.You.Head, target, context.Board.Graph, nil)
	nodes, _ := shortest.To(target.ID())
	if len(nodes) < 2 {
		return ""
	}
	if nodes[len(nodes)-1] != target {
		return ""
	}
	nextCell := nodes[1].(snek.Cell)
	return context.You.Head.ToDirection(nextCell)
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
