package utils

import (
	snek "github.com/JMTyler/battlesnake/_"
	//	"github.com/JMTyler/battlesnake/_/config"
	"fmt"
	//	"time"
)

//let db = require("./db")
//db.Connect(config.get("database_url"))
//	.then((_db) => {
//		db = _db
//	})

func Leftpad(str string, pad int) string {
	gap := pad - len(str)
	prefix := ""
	for i := 0; i < gap; i++ {
		prefix += " "
	}
	return prefix + str
}

var previousTurn int

func LogMove(turn int, move string, comment string) {
	if turn > previousTurn+1 {
		fmt.Println(" [ ... ]")
	}

	moveTag := Leftpad(move, 5)
	turnTag := fmt.Sprintf("[%s]", Leftpad(string(turn), 5))
	if turn == previousTurn {
		turnTag = fmt.Sprintf(" %s ", Leftpad("↳", 5))
	}

	previousTurn = turn

	fmt.Printf(" %s %s :  %s\n", turnTag, moveTag, comment)
}

// TODO: Merge LogMove and RecordFrame, once we can be sure which move was the final choice.
func RecordFrame(context snek.Context, update struct{} /*= null*/) {
	if context.Game.Dev {
		return
	}

	//	NOW := time.Now()

	//	if (update) {
	//		return await db.Frames.update({
	//			game_id  : context.game.id,
	//			snake_id : context.you.id,
	//			turn     : context.turn,
	//		}, {
	//			...update,
	//			updated_at : NOW,
	//		})
	//	}
	//
	//	return await db.Frames.insert({
	//		context,
	//		game_id    : context.game.id,
	//		snake_id   : context.you.id,
	//		name       : context.you.name,
	//		turn       : context.turn,
	//		created_at : NOW,
	//		updated_at : NOW,
	//	}, { onConflict: { action: "ignore" } })
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
