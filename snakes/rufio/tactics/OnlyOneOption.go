package tactics

import snek "github.com/JMTyler/battlesnake/_"

type OnlyOneOption struct{}

func (_ OnlyOneOption) Run(ctx *snek.Context, _ *snek.State) string {
	options := make([]string, 0)
	for dir, cell := range ctx.You.Head.GetAdjacentCells() {
		if !cell.IsDeadly() {
			options = append(options, dir)
		}
	}

	if len(options) == 1 {
		return options[0]
	}

	return ""
}
