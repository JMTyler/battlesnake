package tactics

import snek "github.com/JMTyler/battlesnake/_"

type OnlyOneOption struct{}

func (_ OnlyOneOption) Run(context snek.Context, _ *snek.State) string {
	adjacent := context.You.Head.GetAdjacentCells()

	options := make([]string, 0)
	for dir, cell := range adjacent {
		if !cell.IsDeadly(context) {
			options = append(options, dir)
		}
	}

	if len(options) == 1 {
		return options[0]
	}

	return ""
}
