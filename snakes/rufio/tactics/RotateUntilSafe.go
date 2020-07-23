package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
)

var rotate = map[string]string{
	"right": "down",
	"down":  "left",
	"left":  "up",
	"up":    "right",
}

type RotateUntilSafe struct{}

func (_ RotateUntilSafe) Run(context snek.Context, state *snek.State) string {
	adjacent := context.You.Head.GetAdjacentCells()

	isSafe := false
	move := state.Move
	for turns := 0; !isSafe && turns < 4; turns += 1 {
		move = rotate[move]
		isSafe = adjacent[move].IsSafe(context)
		//utils.LogMove(context.turn, move, 'Rotate Until Safe');
	}

	if isSafe {
		return move
	}

	// If there are no safe cells nearby, we have to be willing to move into risky cells.
	// Prioritising our current direction.
	if !adjacent[state.Move].IsDeadly(context) {
		return state.Move
	}

	// But if our current direction is deadly, resort to any adjacent risky cell.
	for turns := 0; !isSafe && turns < 4; turns += 1 {
		move = rotate[move]
		isSafe = !adjacent[move].IsDeadly(context)
		//utils.LogMove(context.turn, move, 'Rotate Until Risky');
	}

	return move
}
