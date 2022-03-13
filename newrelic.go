package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/config"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var newRelicApp *newrelic.Application

func getNewRelicApp() *newrelic.Application {
	if newRelicApp == nil {
		newRelicApp = initNewRelic()
	}
	return newRelicApp
}

func attachSnakeDataToNewRelic(ctx *snek.Context, tx *newrelic.Transaction) {
	food, err := json.Marshal(ctx.Board.Food)
	if err != nil {
		panic(err)
	}
	hazards, err := json.Marshal(ctx.Board.Hazards)
	if err != nil {
		panic(err)
	}
	you, err := json.Marshal(map[string]interface{}{"body": ctx.You.FullBody, "head": ctx.You.Head})
	if err != nil {
		panic(err)
	}

	tx.AddAttribute("snakeGameId", ctx.Game.ID)
	tx.AddAttribute("snakeRules", ctx.Game.Ruleset.Name)
	tx.AddAttribute("snakeTurn", ctx.Turn)
	tx.AddAttribute("snakeBoardHeight", ctx.Board.Height)
	tx.AddAttribute("snakeBoardWidth", ctx.Board.Width)
	tx.AddAttribute("snakeBoardFood", base64.StdEncoding.EncodeToString(food))
	tx.AddAttribute("snakeBoardHazards", base64.StdEncoding.EncodeToString(hazards))
	tx.AddAttribute("snakeName", ctx.You.Name)
	tx.AddAttribute("snakeId", ctx.You.ID)
	tx.AddAttribute("snakeHealth", ctx.You.Health)
	tx.AddAttribute("snakeLength", ctx.You.Length)
	tx.AddAttribute("snakeData", base64.StdEncoding.EncodeToString(you))
	for i, foe := range ctx.Board.Foes {
		tx.AddAttribute(fmt.Sprintf("snakeOpponent_%d_Name", i+1), foe.Name)
		tx.AddAttribute(fmt.Sprintf("snakeOpponent_%d_Id", i+1), foe.ID)
		tx.AddAttribute(fmt.Sprintf("snakeOpponent_%d_Health", i+1), foe.Health)
		tx.AddAttribute(fmt.Sprintf("snakeOpponent_%d_Length", i+1), foe.Length)
		foeData, err := json.Marshal(map[string]interface{}{"body": foe.FullBody, "head": foe.Head})
		if err != nil {
			panic(err)
		}
		tx.AddAttribute(fmt.Sprintf("snakeOpponent_%d_Data", i+1), base64.StdEncoding.EncodeToString(foeData))
	}
}

func initNewRelic() *newrelic.Application {
	license := config.Get("new_relic_license_key", "")
	if license == "" {
		return nil
	}

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(config.Get("new_relic_app_name", "")),
		newrelic.ConfigLicense(license),
	)
	if err != nil {
		panic(err)
	}
	return app
}
