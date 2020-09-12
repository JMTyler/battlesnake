package battlesnake

import "sync"

var states = &sync.Map{}

type State struct {
	Move   string
	Snakes map[string]*SnakeState
}

type SnakeState struct {
	Head *Cell
	Move string
}

func InitState(ctx *Context, value *State) {
	states.Store(ctx.Game.ID+"---"+ctx.You.ID, value)
}

func DeleteState(ctx *Context) {
	states.Delete(ctx.Game.ID+"---"+ctx.You.ID)
}

func GetState(ctx *Context) *State {
	state, ok := states.Load(ctx.Game.ID+"---"+ctx.You.ID)
	if !ok {
		return &State{Move: "right", Snakes: make(map[string]*SnakeState)}
	}
	return state.(*State)
}

func (state *State) UpdateSnakeHistory(ctx *Context) {
	// Figure out which move each snake took during the *last* turn, and toss it into state.
	for _, snake := range ctx.Board.Foes {
		prev, exists := state.Snakes[snake.ID]
		move := "up"
		if exists {
			move = prev.Head.ToDirection(snake.Head)
		}
		state.Snakes[snake.ID] = &SnakeState{snake.Head, move}
	}
}
