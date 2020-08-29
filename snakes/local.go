package snakes

import (
	snek "github.com/JMTyler/battlesnake/_"
)

type Local struct{}

var rufio *Rufio = &Rufio{}

func (me *Local) GetName() string {
	return "local"
}

func (me *Local) GetInfo() SnakeInfo {
	return SnakeInfo{
		Color: "#008F00",
		Head:  "shac-workout",
		Tail:  "freckled",
	}
}

func (me *Local) StartGame(ctx *snek.Context) {
	rufio.StartGame(ctx)
}

func (me *Local) Move(ctx *snek.Context) string {
	return rufio.Move(ctx)
}

func (me *Local) EndGame(ctx *snek.Context) {
	rufio.EndGame(ctx)
}
