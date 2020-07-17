package board

import (
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/movement"
)

func FindClosestFood(context snek.Context) snek.Position {
	return movement.FindClosestTarget(context.You.Head, context.Board.Food)
}
