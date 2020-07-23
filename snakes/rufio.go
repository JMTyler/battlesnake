package snakes

import (
	"fmt"
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/movement"
	"github.com/JMTyler/battlesnake/_/utils"
	"github.com/JMTyler/battlesnake/snakes/rufio/tactics"
	"path/filepath"
	"runtime"
)

type Rufio struct{}

var strategy = []tactics.Tactic{
	tactics.New("Easy Kill", tactics.Aggrieve{Advantage: 1, Distance: 1}),
	tactics.New("Quick Snack", tactics.Eat{Distance: 2}),
	tactics.New("Abscond", tactics.Abscond{Disadvantage: 1, Distance: 3}),
	tactics.New("Hunt", tactics.Aggrieve{Advantage: 2}),
	tactics.New("Hungry", tactics.Eat{}),
	tactics.New("Go Centre", tactics.GoCentre{Width: 3, Height: 3}),
	tactics.New("Continue", tactics.Continue{}),
	tactics.New("Seek Tail", tactics.SeekTail{}),
	tactics.New("Rotate Until Safe", tactics.RotateUntilSafe{}),
}

func (me *Rufio) Move(context snek.Context) string {
	state := snek.GetState(context)
	adjacent := context.You.Head.GetAdjacentCells()

	state.UpdateSnakeHistory(context)

	// Remove `You` snake from the `Snakes` array since we only ever want an array of enemies.
	for i, snake := range context.Board.Snakes {
		if snake.ID == context.You.ID {
			context.Board.Snakes = append(context.Board.Snakes[:i], context.Board.Snakes[i+1:]...)
			break
		}
	}

	movement.InitPathfinder(&context)

	move := ""
	for _, tactic := range strategy {
		result := tactic.Run(context, state)
		if result == "" {
			continue
		}

		utils.LogMove(context.Turn, result, tactic.Description())

		pos, withinBounds := adjacent[result]
		if !withinBounds {
			continue
		}

		isSafe := pos.IsSafe(context)
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

func (me *Rufio) StartGame(context snek.Context) {
	snek.InitState(context, snek.State{
		Move:   "right",
		Snakes: make(map[string]snek.SnakeState),
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

	snek.DeleteState(context)
}

func currentDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}
