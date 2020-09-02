package utils

import (
	"fmt"
	"github.com/JMTyler/battlesnake/_/config"
	"strings"
)

func Leftpad(str string, pad int) string {
	prefix := strings.Repeat(" ", pad-len(str))
	return prefix + str
}

var previousTurn int = -1

// TODO: Merge LogMove() and Frame.Insert()/.Update(), once we can be sure which move was the final choice.
func LogMove(turn int, move string, comment string) {
	if !config.GetBool("debug") {
		return
	}

	if turn > previousTurn+1 {
		fmt.Println(" [ ... ]")
	}

	moveTag := Leftpad(move, 5)
	turnTag := fmt.Sprintf("[%s]", Leftpad(fmt.Sprintf("%v", turn), 5))
	if turn == previousTurn {
		turnTag = fmt.Sprintf(" %s ", Leftpad("↳", 5))
	}

	previousTurn = turn

	fmt.Printf(" %s %s :  %s\n", turnTag, moveTag, comment)
}

func Clamp(val int, min int, max int) int {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}
