package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/movement"
)

func SeekTail() func(snek.Context, snek.State) string {
	return func(context snek.Context, state snek.State) string {
		return movement.ApproachTarget(context.You.Body[len(context.You.Body)-1], context)
	}
}
