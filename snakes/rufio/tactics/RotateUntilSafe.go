package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/position"
)

var rotate = map[string]string{
	"right": "down",
	"down":  "left",
	"left":  "up",
	"up":    "right",
}

type RotateUntilSafe struct{}

func (tactic RotateUntilSafe) Run(context snek.Context, state *snek.State) string {
	adjacent := position.GetAdjacentTiles(context.You.Head)

	isSafe := false
	move := state.Move
	for turns := 0; !isSafe && turns < 4; turns += 1 {
		move = rotate[move]
		isSafe = position.IsSafe(adjacent[move], context)
		//utils.LogMove(context.turn, move, 'Rotate Until Safe');
	}

	if isSafe {
		return move
	}

	// If there are no safe cells nearby, we have to be willing to move into risky cells.
	// Prioritising our current direction.
	if !position.IsDeadly(adjacent[state.Move], context) {
		return state.Move
	}

	// But if our current direction is deadly, resort to any adjacent risky cell.
	for turns := 0; !isSafe && turns < 4; turns += 1 {
		move = rotate[move]
		isSafe = !position.IsDeadly(adjacent[move], context)
		//utils.LogMove(context.turn, move, 'Rotate Until Risky');
	}

	return move
}
