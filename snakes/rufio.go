package snakes

import (
	"fmt"
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/config"
	"github.com/JMTyler/battlesnake/_/utils"
	"github.com/JMTyler/battlesnake/snakes/rufio/tactics"
	"path/filepath"
	"runtime"
)

type Rufio struct{}

var standard = []tactics.Tactic{
	tactics.New("Only One Option", tactics.OnlyOneOption{}),
	tactics.New("Easy Kill", tactics.Aggrieve{Advantage: 1, Distance: 2}),
	tactics.New("Quick Snack", tactics.Eat{Distance: 2}),
	tactics.New("Abscond", tactics.Abscond{Disadvantage: 1, Distance: 3}),
	tactics.New("Hunt", tactics.Aggrieve{Advantage: 2}),
	tactics.New("Hungry", tactics.Eat{}),
	tactics.New("Go Centre", tactics.GoCentre{Width: 3, Height: 3}),
	tactics.New("Seek Tail", tactics.SeekTail{}),
	// TODO: Seek other snake's tail if available
}

var snackatron = []tactics.Tactic{
	tactics.New("Only One Option", tactics.OnlyOneOption{}),
	tactics.New("Abscond", tactics.Abscond{Disadvantage: 1, Distance: 2}),
	tactics.New("Go Centre", tactics.GoCentre{Width: 10, Height: 10}),
	tactics.New("Seek Tail", tactics.SeekTail{}),
}

func (me *Rufio) GetName() string {
	return "rufio"
}

func (me *Rufio) GetInfo() SnakeInfo {
	// TODO: It's a *lot* more work to load info from a file now than it was in Node... is it worth it?
	//	file, err := os.Open(currentDir() + "/rufio/info.yaml")
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	var info SnakeInfo
	//	decoder := yaml.NewDecoder(file)
	//	if err := decoder.Decode(&info); err != nil {
	//		panic(err)
	//	}
	//	return info

	return SnakeInfo{
		Color: "#DF0000",
		Head:  "shades",
		Tail:  "bolt",
	}
}

func (me *Rufio) Move(ctx *snek.Context) string {
	state := snek.GetState(ctx)
	neighbours := ctx.You.Head.Neighbours()

	state.UpdateSnakeHistory(ctx)

	strategy := standard
	if len(ctx.Board.Food) > (ctx.Board.Width * ctx.Board.Height / 2) {
		strategy = snackatron
	}

	move := ""
	riskyBusiness := ""
	for _, tactic := range strategy {
		result := tactic.Run(ctx, state)
		if result == "" {
			continue
		}

		utils.LogMove(ctx.Turn, result, tactic.Description())

		cell, withinBounds := neighbours[result]
		if !withinBounds {
			continue
		}

		if !cell.IsSafe(ctx) {
			if !cell.IsDeadly() && riskyBusiness == "" {
				riskyBusiness = result
			}
			continue
		}

		move = result
		break
	}

	if move != "" {
		state.Move = move
		return move
	}

	if riskyBusiness != "" {
		utils.LogMove(ctx.Turn, riskyBusiness, "Risky Business")
		state.Move = riskyBusiness
		return riskyBusiness
	}

	// TODO: Should still prefer to pick a random adjacent empty cell before fully welping out.
	nonWalls := make([]string, 0)
	emptyCells := make([]string, 0)
	isContinueOpen := false
	for dir, cell := range neighbours {
		nonWalls = append(nonWalls, dir)
		if !cell.IsDeadly() {
			emptyCells = append(emptyCells, dir)
		}
		if dir == state.Move {
			isContinueOpen = true
			break
		}
	}

	if !isContinueOpen {
		state.Move = nonWalls[0]
		if len(emptyCells) > 0 {
			state.Move = emptyCells[0]
		}
	}

	utils.LogMove(ctx.Turn, state.Move, "welp 👋")
	return state.Move
}

func (me *Rufio) StartGame(ctx *snek.Context) {
	snek.InitState(ctx, &snek.State{
		Move:   "right",
		Snakes: make(map[string]*snek.SnakeState),
	})

	if config.GetBool("debug") {
		fmt.Println("-----")
		fmt.Println()
	}
}

func (me *Rufio) EndGame(ctx *snek.Context) {
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

func currentDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}
