package snakes

import (
	snek "github.com/JMTyler/battlesnake/_"
)

type SnakeService interface {
	GetName() string
	GetInfo() map[string]string
	StartGame(snek.Context)
	Move(snek.Context) string
	EndGame(snek.Context)
}
