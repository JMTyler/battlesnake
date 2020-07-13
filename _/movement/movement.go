package movement

import (
	snek "github.com/JMTyler/battlesnake/_"
	//	"github.com/JMTyler/battlesnake/_/position"
	"math"
)

//var pathfinding struct{}
//const pathfinding = require('pathfinding')
//const pathfinder = new pathfinding.JumpPointFinder({ diagonalMovement: pathfinding.DiagonalMovement.Never })

//func InitPathfinder(context snek.Context) {
//	const grid = new pathfinding.Grid(context.board.width, context.board.height)
//
//	_.each(_.initial(context.you.body), ({ x, y }) => {
//		grid.setWalkableAt(x, y, false)
//	})
//
//	// TODO: Consider adding a safeGrid (this is a riskyGrid) that also avoids risky cells.
//	_.each(context.board.snakes, (snake) => {
//		_.each(_.initial(snake.body), ({ x, y }) => {
//			grid.setWalkableAt(x, y, false)
//		})
//	})
//
//	// Make context.grid an accessor that always returns a clone.
//	Object.defineProperty(context, 'grid', { get: () => grid.clone() })
//}

// TODO: Use pathfinding distance, not direct distance.
func GetDistance(origin snek.Position, target snek.Position) int {
	x := math.Abs(float64(target.X - origin.X))
	y := math.Abs(float64(target.Y - origin.Y))
	return int(x + y)
}
func GetDistances(origin snek.Position, targets []snek.Position) []int {
	distances := make([]int, len(targets))
	for i, target := range targets {
		distances[i] = GetDistance(origin, target)
	}
	return distances
}

func GetVector(origin snek.Position, target snek.Position) map[string]interface{} {
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

	return map[string]interface{}{
		"dir": map[string]string{
			"x": xDir,
			"y": yDir,
		},
		"weight": map[string]int{"x": x, "y": y},
	}
}

//func ApproachTarget(target snek.Position, context snek.Context) string {
//	const path = pathfinder.findPath(you.head.x, you.head.y, target.x, target.y, grid)
//	const nextCell = path[1]
//	if !nextCell {
//		return "up"
//	}
//	const pos = snek.Position{X: nextCell[0], Y: nextCell[1]}
//	return position.ToDirection(pos, you.head)
//}

func FindClosestTarget(origin snek.Position, targets []snek.Position) snek.Position {
	if len(targets) == 1 {
		return targets[0]
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
