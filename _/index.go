package tactics

type TacticOptions struct {
	Health       int
	Distance     int
	Advantage    int
	Disadvantage int
}

type Position struct {
	X int
	Y int
}

type Context struct {
	Turn  int
	Game  Game
	You   Snake
	Board Board
}

type State struct {
	Move   string
	Snakes map[string]SnakeState
}

type Game struct {
	ID      string
	Timeout int
	Dev     bool
}

type Snake struct {
	ID     string
	Head   Position
	Body   []Position
	Length int
	Health int
}

type Board struct {
	Width  int
	Height int
	Snakes []Snake
	Food   []Position
}

type SnakeState struct {
	Head Position
	Move string
}

var states = make(map[string]State)

func InitState(context Context, value State) {
	states[context.Game.ID+"---"+context.You.ID] = value
}

func GetState(context Context) State {
	state, ok := states[context.Game.ID+"---"+context.You.ID]
	if !ok {
		return State{Move: "right", Snakes: make(map[string]SnakeState)}
	}
	return state
}
