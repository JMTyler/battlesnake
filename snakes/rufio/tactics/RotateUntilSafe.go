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

func (_ RotateUntilSafe) Run(ctx *snek.Context, state *snek.State) string {
	neighbours := ctx.You.Head.Neighbours()

	isSafe := false
	move := state.Move
	for turns := 0; !isSafe && turns < 4; turns += 1 {
		move = rotate[move]
		if cell, ok := neighbours[move]; ok {
			isSafe = cell.IsSafe(ctx)
		}
		//utils.LogMove(ctx.turn, move, 'Rotate Until Safe');
	}

	if isSafe {
		return move
	}

	// If there are no safe cells nearby, we have to be willing to move into risky cells.
	// Prioritising our current direction.
	if cell, ok := neighbours[state.Move]; ok && !cell.IsDeadly() {
		return state.Move
	}

	// But if our current direction is deadly, resort to any adjacent risky cell.
	for turns := 0; !isSafe && turns < 4; turns += 1 {
		move = rotate[move]
		if cell, ok := neighbours[move]; ok {
			isSafe = !cell.IsDeadly()
		}
		//utils.LogMove(ctx.turn, move, 'Rotate Until Risky');
	}

	return move
}
