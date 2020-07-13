package movement

import (
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/position"
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
	//	const path = pathfinder.findPath(you.head.x, you.head.y, target.x, target.y, grid)
	//	const nextCell = path[1]
	//	if !nextCell {
	//		return "up"
	//	}
	//	const pos = snek.Position{X: nextCell[0], Y: nextCell[1]}
	//	return position.ToDirection(pos, you.head)

	vector := GetVector(context.You.Head, target)
	adjacent := position.GetAdjacentTiles(context.You.Head)

	// TODO: Support target being a straight line away, making left/right or up/down equal choices.
	moveX := adjacent[vector.Dir.X]
	if !position.IsSafe(moveX, context) {
		return vector.Dir.Y
	}

	moveY := adjacent[vector.Dir.Y]
	if !position.IsSafe(moveY, context) {
		return vector.Dir.X
	}

	if context.You.Head.X != target.X {
		return vector.Dir.X
	}

	return vector.Dir.Y
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
