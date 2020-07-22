package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
	//"utils"
)

type Continue struct{ Name string }

func (t *Continue) Description() string {
	return t.Name
}

func (tactic *Continue) Run(context snek.Context, state *snek.State) string {
	//if (context.turn === 0) utils.LogMove(context.turn, state.move, 'Initial Move')
	//else utils.LogMove(context.turn, state.move, 'Continue')

	return state.Move
}
