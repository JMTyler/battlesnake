package battlesnake

type Snake struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Head     *Cell   `json:"head"`
	FullBody []*Cell `json:"body"`
	Length   int     `json:"length"`
	Health   int     `json:"health"`
}

func (snake *Snake) Prepare(ctx *Context) {
	// Replace head and body with cell singletons.
	snake.Head = ctx.Board.CellAt(snake.Head.X, snake.Head.Y)
	for ix, part := range snake.FullBody {
		snake.FullBody[ix] = ctx.Board.CellAt(part.X, part.Y)
	}
}

func (snake *Snake) Body() []*Cell {
	return snake.FullBody[:len(snake.FullBody)-1]
}

// TODO: We could set Body to a slice of the original body, then grab the tail here using cap() instead of len()?
func (snake *Snake) Tail() *Cell {
	return snake.FullBody[len(snake.FullBody)-1]
}
