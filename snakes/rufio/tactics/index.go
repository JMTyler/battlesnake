package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
)

// HACK: This stuff feels weird & janky, but it keeps both ends clean & simple (creating tactics and consuming them).
// Should take another look at this file later and think if there's a more idiomatic Go way to handle it.

type Tactic interface {
	Description() string
	Run(*snek.Context, *snek.State) string
}

type runnable interface {
	Run(*snek.Context, *snek.State) string
}

type wrapper struct {
	description string
	tactic      runnable
}

func (w *wrapper) Description() string {
	return w.description
}

func (w *wrapper) Run(ctx *snek.Context, state *snek.State) string {
	return w.tactic.Run(ctx, state)
}

func New(description string, tactic runnable) Tactic {
	return &wrapper{
		description: description,
		tactic:      tactic,
	}
}
