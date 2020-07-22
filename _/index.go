package battlesnake

import (
	"gonum.org/v1/gonum/graph/traverse"
)

type Context struct {
	Turn  int   `json:"turn"`
	Game  Game  `json:"game"`
	You   Snake `json:"you"`
	Board Board `json:"board"`
}

type Game struct {
	ID      string `json:"id"`
	Timeout int    `json:"timeout"`
	Dev     bool   `json:"dev"`
}

type Snake struct {
	ID     string     `json:"id"`
	Name   string     `json:"name"`
	Head   Position   `json:"head"`
	Body   []Position `json:"body"`
	Length int        `json:"length"`
	Health int        `json:"health"`
}

type Board struct {
	Width  int            `json:"width"`
	Height int            `json:"height"`
	Snakes []Snake        `json:"snakes"`
	Food   []Position     `json:"food"`
	Graph  traverse.Graph `json:", ignore"`
}
