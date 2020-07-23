package battlesnake

var states = make(map[string]*State)

type State struct {
	Move   string
	Snakes map[string]SnakeState
}

type SnakeState struct {
	Head Position
	Move string
}

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

func (state *State) UpdateSnakeHistory(context Context) {
	// Figure out which move each snake took during the *last* turn, and toss it into state.
	for _, snake := range context.Board.Enemies {
		prev, exists := state.Snakes[snake.ID]
		move := "up"
		if exists {
			move = prev.Head.ToDirection(snake.Head)
		}
		state.Snakes[snake.ID] = SnakeState{snake.Head, move}
	}
}
