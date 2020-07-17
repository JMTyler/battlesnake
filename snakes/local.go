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

func (me *Local) StartGame(context snek.Context) {
	rufio.StartGame(context)
}

func (me *Local) Move(context snek.Context) string {
	return rufio.Move(context)
}

func (me *Local) EndGame(context snek.Context) {
	rufio.EndGame(context)
}
