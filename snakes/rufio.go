package snakes

import (
	"fmt"
	snek "github.com/JMTyler/battlesnake/_"
	//	"github.com/JMTyler/battlesnake/_/movement"
	"github.com/JMTyler/battlesnake/_/position"
	"github.com/JMTyler/battlesnake/_/utils"
	"github.com/JMTyler/battlesnake/snakes/rufio/tactics"
)

type Rufio struct{}

var strategy = map[string]func(snek.Context, snek.State) string{
	"Easy Kill":         tactics.Aggrieve(snek.TacticOptions{Advantage: 1, Distance: 1}),
	"Quick Snack":       tactics.Eat(snek.TacticOptions{Distance: 2}),
	"Abscond":           tactics.Abscond(snek.TacticOptions{Disadvantage: 1, Distance: 3}),
	"Aggrieve":          tactics.Aggrieve(snek.TacticOptions{Advantage: 2}),
	"Hungry":            tactics.Eat(snek.TacticOptions{}),
	"Go Centre":         tactics.GoCentre(),
	"Continue":          tactics.Continue(),
	"Seek Tail":         tactics.SeekTail(),
	"Rotate Until Safe": tactics.RotateUntilSafe(),
}

func (me *Rufio) Move(context snek.Context) string {
	state := snek.GetState(context)
	adjacent := position.GetAdjacentTiles(context.You.Head)

	//	movement.InitPathfinder(context)

	// Figure out which move each snake took during the *last* turn, and toss it into state.
	for _, snake := range context.Board.Snakes {
		prev, exists := state.Snakes[snake.ID]
		move := "up"
		if exists {
			move = position.ToDirection(snake.Head, prev.Head)
		}
		state.Snakes[snake.ID] = snek.SnakeState{Head: snake.Head, Move: move}
	}

	// Remove `You` snake from the `Snakes` array since we only ever want an array of enemies.
	for i, snake := range context.Board.Snakes {
		if snake.ID == context.You.ID {
			context.Board.Snakes = append(context.Board.Snakes[:i], context.Board.Snakes[i+1:]...)
			break
		}
	}

	move := ""
	for description, tactic := range strategy {
		result := tactic(context, state)
		if result == "" {
			continue
		}

		utils.LogMove(context.Turn, result, description)
		isSafe := position.IsSafe(adjacent[result], context)
		if !isSafe {
			continue
		}

		move = result
		break
	}

	if move != "" {
		state.Move = move
		return move
	}

	utils.LogMove(context.Turn, state.Move, "welp 👋")
	return state.Move
}

func (me *Rufio) GetName() string {
	return "rufio"
}

func (me *Rufio) GetInfo() map[string]string {
	return map[string]string{}
}

func (me *Rufio) StartGame(context snek.Context) {
	snek.InitState(context, snek.State{
		Move: "right",
	})

	fmt.Println("-----")
	fmt.Println()
}

func (me *Rufio) EndGame(context snek.Context) {
	result := "LOSE"
	// dangerous array access
	if context.You.ID == context.Board.Snakes[0].ID {
		result = "WIN"
	}

	fmt.Println()
	fmt.Printf("* Game Over! %s *\n", result)

	utils.PruneGames(context)
}
