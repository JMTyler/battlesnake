package snakes

import (
	snek "github.com/JMTyler/battlesnake/_"
)

type SnakeInfo struct {
	APIVersion string `json:"apiversion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
}

type SnakeService interface {
	GetName() string
	GetInfo() SnakeInfo
	StartGame(snek.Context)
	Move(snek.Context) string
	EndGame(snek.Context)
}
