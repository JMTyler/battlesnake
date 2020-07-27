package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
)

type SeekTail struct{}

func (_ SeekTail) Run(context snek.Context, _ *snek.State) string {
	return context.You.Head.ApproachTarget(context.You.Tail(), context)
}
