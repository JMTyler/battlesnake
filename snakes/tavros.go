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

	// Grab food if you're close to it, or if you're super hungry.
	if food := ctx.You.Head.FindClosestTarget(ctx.Board.Food); food != nil {
		if ctx.You.Health <= 20 || ctx.You.Head.GetDistance(food) <= 2 {
			move = ctx.You.Head.ApproachTarget(food, ctx)
		}
	}

	// Chase your tail.
	if move == "" {
		move = ctx.You.Head.ApproachTarget(ctx.You.Tail(), ctx)
	}

	if move == "" {
		move = state.Move
	}

	// Don't ever move into a deadly space, even if the alternative is random.
	target := ctx.You.Head.Adjacent(move)
	if target == nil || target.IsDeadly(ctx) {
		for dir, cell := range ctx.You.Head.GetAdjacentCells() {
			if !cell.IsDeadly(ctx) {
				move = dir
				break
			}
		}
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
