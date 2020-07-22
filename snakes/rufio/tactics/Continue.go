package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
)

type Continue struct{}

func (tactic Continue) Run(_ snek.Context, state *snek.State) string {
	return state.Move
}
