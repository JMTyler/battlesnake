package battlesnake

// TODO: Make a Cell-like interface that works with Snake, which just asks Head.
// TODO: Maybe Cell is not even a struct of its own? Snake.Head/Food/etc. just implement it?
type Snake struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Head     *Cell   `json:"head"`
	FullBody []*Cell `json:"body"`
	Length   int     `json:"length"`
	Health   int     `json:"health"`
	Squad    string  `json:"squad"`
}

func (snake *Snake) Prepare(ctx *Context) {
	// Replace head and body with cell singletons.
	snake.Head = ctx.Board.CellAt(snake.Head.X, snake.Head.Y)
	for ix, part := range snake.FullBody {
		snake.FullBody[ix] = ctx.Board.CellAt(part.X, part.Y)
	}

	snake.Head.AddTags("head")
	snake.Tail().AddTags("tail")
	for _, part := range snake.Body() {
		part.AddTags("body")
	}

	if snake.ID == ctx.You.ID {
		for _, part := range snake.FullBody {
			part.AddTags("you")
		}
	} else if snake.Squad == ctx.You.Squad && snake.Squad != "" {
		for _, part := range snake.FullBody {
			part.AddTags("friend")
		}

		neighbours := snake.Head.Neighbours()
		for _, cell := range neighbours {
			cell.AddTags("friend-adjacent")
		}
	} else {
		lengthTag := "enemy-equal"
		switch {
		case snake.Length > ctx.You.Length:
			lengthTag = "enemy-longer"
		case snake.Length < ctx.You.Length:
			lengthTag = "enemy-shorter"
		}

		for _, part := range snake.FullBody {
			part.AddTags("enemy", lengthTag)
		}

		neighbours := snake.Head.Neighbours()
		for _, cell := range neighbours {
			cell.AddTags("enemy-adjacent", lengthTag)
		}
	}
}

// TODO: Body() and Tail() should both be fields that are set during Prepare().
func (snake *Snake) Body() []*Cell {
	return snake.FullBody[:len(snake.FullBody)-1]
}

// TODO: We could set Body to a slice of the original body, then grab the tail here using cap() instead of len()?
func (snake *Snake) Tail() *Cell {
	return snake.FullBody[len(snake.FullBody)-1]
}
