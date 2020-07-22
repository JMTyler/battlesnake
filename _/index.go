package battlesnake

import (
	"fmt"
	"gonum.org/v1/gonum/graph/traverse"
)

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

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

func (pos Position) ID() int64 {
	// HACK: Should replace this magic number with context.Board.Width somehow.
	return int64(pos.X + (pos.Y * 11))
}

func (pos Position) String() string {
	return fmt.Sprintf("(%d,%d)", pos.X, pos.Y)
}
