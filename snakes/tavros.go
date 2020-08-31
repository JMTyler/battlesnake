package snakes

import (
	"fmt"
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/config"
)

type Tavros struct{}

func (me *Tavros) GetName() string {
	return "tavros"
}

func (me *Tavros) GetInfo() SnakeInfo {
	return SnakeInfo{
		Color: "#A15000",
		Head:  "shac-workout",
		Tail:  "fat-rattle",
	}
}

func (me *Tavros) Move(ctx *snek.Context) string {
	state := snek.GetState(ctx)
	state.UpdateSnakeHistory(ctx)

	move := ""

	if food := ctx.You.Head.FindClosestTarget(ctx.Board.Food); food != nil {
		if ctx.You.Health <= 20 || ctx.You.Head.GetDistance(food) <= 2 {
			move = ctx.You.Head.ApproachTarget(food, ctx)
		}
	}

	if move == "" {
		move = ctx.You.Head.ApproachTarget(ctx.You.Tail(), ctx)
	}

	if move == "" {
		move = state.Move
	}

	state.Move = move
	return move
}

func (me *Tavros) StartGame(ctx *snek.Context) {
	snek.InitState(ctx, &snek.State{
		Move:   "right",
		Snakes: make(map[string]*snek.SnakeState),
	})

	if config.GetBool("debug") {
		fmt.Println("-----")
		fmt.Println()
	}
}

func (me *Tavros) EndGame(ctx *snek.Context) {
	if config.GetBool("debug") {
		result := "LOSE"
		if len(ctx.Board.Snakes) > 0 && ctx.You.ID == ctx.Board.Snakes[0].ID {
			result = "WIN"
		}

		fmt.Println()
		fmt.Printf("* Game Over! %s *\n", result)
	}

	snek.DeleteState(ctx)
}
