package utils

import (
	snek "github.com/JMTyler/battlesnake/_"
	//	"github.com/JMTyler/battlesnake/_/config"
	"fmt"
	"strings"
)

func Leftpad(str string, pad int) string {
	prefix := strings.Repeat(" ", pad-len(str))
	return prefix + str
}

var previousTurn int = -1

// TODO: Merge LogMove() and Frame.Insert()/.Update(), once we can be sure which move was the final choice.
func LogMove(turn int, move string, comment string) {
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

func PruneGames(context snek.Context) {
	if context.Game.Dev {
		return
	}

	//	const numRows = _.toNumber(await db.Frames.count())
	//	if (numRows < 9500) {
	//		return
	//	}
	//
	//	// Find the oldest game in the database.
	//	const { game_id } = await db.Frames.findOne({ important: false }, {
	//		fields : ["game_id"],
	//		order  : [{ field: "created_at", direction: "asc" }],
	//	})
	//
	//	// And delete it.
	//	await db.Frames.destroy({ game_id })
	//
	//	return await PruneGames()
}
