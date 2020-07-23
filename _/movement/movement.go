package movement

import (
	snek "github.com/JMTyler/battlesnake/_"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
	"math"
)

func InitPathfinder(context *snek.Context) {
	grid := simple.NewUndirectedGraph()
	for x := 0; x < context.Board.Width; x++ {
		for y := 0; y < context.Board.Height; y++ {
			node := snek.Position{x, y}
			if !node.IsDeadly(*context) {
				grid.AddNode(node)
			}
		}
	}
	for x := 0; x < context.Board.Width; x++ {
		for y := 0; y < context.Board.Height; y++ {
			node := snek.Position{x, y}
			if grid.Node(node.ID()) != nil {
				for _, cell := range node.GetAdjacentCells() {
					if grid.Node(cell.ID()) != nil {
						grid.SetEdge(grid.NewEdge(node, cell))
					}
				}
			}
		}
	}
	context.Board.Graph = grid
}

// TODO: Use pathfinding distance, not direct distance.
func GetDistance(origin snek.Position, target snek.Position) int {
	x := math.Abs(float64(target.X - origin.X))
	y := math.Abs(float64(target.Y - origin.Y))
	return int(x + y)
}
func GetDistances(origin snek.Position, targets []snek.Position) []int {
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

func GetVector(origin snek.Position, target snek.Position) Vector {
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

func ApproachTarget(target snek.Position, context snek.Context) string {
	shortest, _ := path.AStar(context.You.Head, target, context.Board.Graph, nil)
	nodes, _ := shortest.To(target.ID())
	if len(nodes) < 2 {
		return ""
	}
	if nodes[len(nodes)-1] != target {
		return ""
	}
	nextCell := nodes[1].(snek.Position)
	return context.You.Head.ToDirection(nextCell)
}

func FindClosestTarget(origin snek.Position, targets []snek.Position) snek.Position {
	if len(targets) == 1 {
		return targets[0]
	}

	if len(targets) == 0 {
		return snek.Position{}
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
