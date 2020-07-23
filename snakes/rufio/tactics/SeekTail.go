package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/movement"
)

type SeekTail struct{}

func (_ SeekTail) Run(context snek.Context, _ *snek.State) string {
	return movement.ApproachTarget(context.You.Tail(), context)
}
