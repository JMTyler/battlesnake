package battlesnake

type Snake struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Head     Cell   `json:"head"`
	fullBody []Cell `json:"body"`
	Length   int    `json:"length"`
	Health   int    `json:"health"`
}

func (snake Snake) Body() []Cell {
	return snake.fullBody[:len(snake.fullBody)-1]
}

// TODO: We could set Body to a slice of the original body, then grab the tail here using cap() instead of len()?
func (snake Snake) Tail() Cell {
	return snake.fullBody[len(snake.fullBody)-1]
}
