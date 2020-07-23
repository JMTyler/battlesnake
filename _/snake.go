package battlesnake

type Snake struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Head     Position   `json:"head"`
	fullBody []Position `json:"body"`
	Length   int        `json:"length"`
	Health   int        `json:"health"`
}

func (snake Snake) Body() []Position {
	return snake.fullBody[:len(snake.fullBody)-1]
}

// TODO: We could set Body to a slice of the original body, then grab the tail here using cap() instead of len()?
func (snake Snake) Tail() Position {
	return snake.fullBody[len(snake.fullBody)-1]
}
