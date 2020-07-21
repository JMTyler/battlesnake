package tactics

import (
	"fmt"
	"gonum.org/v1/gonum/graph/traverse"
)

type TacticOptions struct {
	Health       int
	Distance     int
	Advantage    int
	Disadvantage int
}

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

type State struct {
	Move   string
	Snakes map[string]SnakeState
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

type SnakeState struct {
	Head Position
	Move string
}

func (pos Position) ID() int64 {
	// HACK: Should replace this magic number with context.Board.Width somehow.
	return int64(pos.X + (pos.Y * 11))
}

func (pos Position) String() string {
	return fmt.Sprintf("(%d,%d)", pos.X, pos.Y)
}

var states = make(map[string]*State)

func InitState(context Context, value State) {
	states[context.Game.ID+"---"+context.You.ID] = &value
}

func DeleteState(context Context) {
	delete(states, context.Game.ID+"---"+context.You.ID)
}

func GetState(context Context) *State {
	state, ok := states[context.Game.ID+"---"+context.You.ID]
	if !ok {
		return &State{Move: "right", Snakes: make(map[string]SnakeState)}
	}
	return state
}
