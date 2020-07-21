package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
	//"utils"
)

func Continue() func(snek.Context, *snek.State) string {
	return func(context snek.Context, state *snek.State) string {
		//if (context.turn === 0) utils.LogMove(context.turn, state.move, 'Initial Move')
		//else utils.LogMove(context.turn, state.move, 'Continue')

		return state.Move
	}
}
