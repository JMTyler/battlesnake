package tactics

import (
	snek "github.com/JMTyler/battlesnake/_"
)

type Abscond struct {
	Disadvantage int
	Distance     int
}

func (opts Abscond) Run(ctx *snek.Context, state *snek.State) string {
	if opts.Disadvantage == 0 {
		opts.Disadvantage = 1
	}

	var predators []*snek.Cell
	for _, snake := range ctx.Board.Foes {
		if ctx.You.Length <= snake.Length-opts.Disadvantage {
			predators = append(predators, snake.Head)
		}
	}

	if len(predators) == 0 {
		return ""
	}

	predator := ctx.You.Head.FindClosest(predators)

	if opts.Distance > 0 {
		distanceToPredator := ctx.You.Head.GetDistance(predator)
		if distanceToPredator > opts.Distance {
			return ""
		}
	}

	youContinue := ctx.You.Head.Neighbour(state.Move)
	theyContinue := predator.Neighbour(state.Snakes[getSnakeByHead(ctx.Board.Foes, predator).ID].Move)

	moves := make([]string, 0)
	for dir, n := range ctx.You.Head.Neighbours() {
		if !n.IsDeadly() {
			if n != youContinue || youContinue != theyContinue {
				// Only valid moves are safe, and are NOT the continuation of *both* your path and the predator's path.
				moves = append(moves, dir)
			}
		}
	}

	if len(moves) == 1 {
		// (If all options take you into a risky cell,) prioritise the move that either doesn't continue *your own* path, or doesn't take you into the cell that continues the predator's path.
		return moves[0]
	}

	escapeVector := ctx.You.Head.GetVector(predator)
	escapeVector.Weight.X *= -1
	escapeVector.Weight.Y *= -1
	escapeTarget := ctx.You.Head.Translate(escapeVector.Weight)


	/* Sometimes escapeTarget can't be directly satisfied (clamped to your current position if you're next to the wall;
	   target is on your own body; etc.), causing this tactic to get skipped.  This is a problem when there are still
	   valid ways to abscond and you really should be taking them.
	   TODO: Make the escape target/vector smarter so you still abscond when you need to abscond. */
	return ctx.You.Head.Approach(escapeTarget)
}

func getSnakeByHead(snakes []*snek.Snake, head *snek.Cell) *snek.Snake {
	for _, snake := range snakes {
		if snake.Head == head {
			return snake
		}
	}
	return nil
}
