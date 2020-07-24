package movement

import (
	snek "github.com/JMTyler/battlesnake/_"
	"gonum.org/v1/gonum/graph/path"
	"math"
)

// TODO: Use pathfinding distance, not direct distance.
func GetDistance(origin snek.Cell, target snek.Cell) int {
	x := math.Abs(float64(target.X - origin.X))
	y := math.Abs(float64(target.Y - origin.Y))
	return int(x + y)
}
func GetDistances(origin snek.Cell, targets []snek.Cell) []int {
	var distances []int
	for _, target := range targets {
		distances = append(distances, GetDistance(origin, target))
	}
	return distances
}

type Vector struct {
	Dir struct {
		X string
		Y string
	}
	Weight struct {
		X int
		Y int
	}
}

func GetVector(origin snek.Cell, target snek.Cell) Vector {
	x := target.X - origin.X
	y := target.Y - origin.Y

	xDir := "left"
	if x > 0 {
		xDir = "right"
	}

	yDir := "down"
	if y > 0 {
		yDir = "up"
	}

	return Vector{
		Dir: struct {
			X string
			Y string
		}{
			X: xDir,
			Y: yDir,
		},
		Weight: struct {
			X int
			Y int
		}{
			X: x,
			Y: y,
		},
	}
}

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

	distances := GetDistances(origin, targets)

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
