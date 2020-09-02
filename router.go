package main

import (
	"encoding/json"
	"fmt"
	snek "github.com/JMTyler/battlesnake/_"
	"github.com/JMTyler/battlesnake/_/config"
	"github.com/getsentry/sentry-go"
	"net/http"
	"time"
	//	"github.com/JMTyler/battlesnake/_/utils"
	"github.com/JMTyler/battlesnake/_/db"
	"github.com/JMTyler/battlesnake/snakes"
	"io/ioutil"
)

var the_snakes = []snakes.SnakeService{
	&snakes.Local{},
	&snakes.Rufio{},
	&snakes.Proxy{},
	&snakes.Tavros{},
}

func handleRoute(route string, snake snakes.SnakeService, f func(snakes.SnakeService, *snek.Context) []byte) {
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// HACK: Would be simpler to use sentry.Recover() but it doesn't seem to work as expected.
			err := recover()
			if err != nil {
				w.WriteHeader(500)

				if exception, ok := err.(error); ok {
					sentry.CaptureException(exception)
				} else if str, ok := err.(string); ok {
					sentry.CaptureMessage(str)
				} else {
					sentry.CaptureMessage(fmt.Sprintf("Object: %#v", err))
				}

				fmt.Println("Flushing Sentry...")
				sentry.Flush(time.Second)
			}
		}()

		start := time.Now()

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			return
		}
		if config.GetBool("logging.router") {
			fmt.Println("[http]", r.Method, r.RequestURI)
		}
		w.Header().Add("Content-Type", "application/json")

		var ctx *snek.Context
		if r.Method == "POST" {
			bytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}

			if err := json.Unmarshal(bytes, &ctx); err != nil {
				panic(err)
			}
			ctx.Prepare()

			sentry.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetRequest(r)
				scope.SetUser(sentry.User{ID: ctx.Game.ID})
				scope.SetTag("game", fmt.Sprintf("https://play.battlesnake.com/g/%s/?turn=%v", ctx.Game.ID, ctx.Turn))
				scope.SetTag("turn", fmt.Sprintf("%v", ctx.Turn))
			})
		}

		if response := f(snake, ctx); response != nil {
			w.Write(response)
		}

		duration := time.Now().Sub(start).Milliseconds()
		if duration >= 350 {
			// If request takes longer than 400ms, something is wrong.
			sentry.WithScope(func(scope *sentry.Scope) {
				scope.SetLevel(sentry.LevelWarning)
				scope.SetTag("duration", fmt.Sprintf("%v", duration))
				sentry.CaptureMessage("Move took unusually long to calculate.")
			})
		}

		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.Clear()
		})
	})
}

// TODO: Setup root paths to default to local snake.
func RouteSnakes() {
	for _, snake := range the_snakes {
		prefix := "/" + snake.GetName()
		handleRoute(prefix+"/", snake, func(snake snakes.SnakeService, ctx *snek.Context) []byte {
			info := snake.GetInfo()
			info.APIVersion = "1"
			info.Author = "JMTyler"

			payload, _ := json.Marshal(info)
			return payload
		})

		handleRoute(prefix+"/start", snake, func(snake snakes.SnakeService, ctx *snek.Context) []byte {
			snake.StartGame(ctx)
			return nil
		})

		handleRoute(prefix+"/move", snake, func(snake snakes.SnakeService, ctx *snek.Context) []byte {
			frame := db.NewFrame(ctx)
			if !ctx.Game.Dev {
				frame.Insert()
			}

			start := time.Now()
			move := snake.Move(ctx)
			duration := time.Now().Sub(start).Microseconds()

			if !ctx.Game.Dev {
				frame.Update(move, duration)
			}

			payload, err := json.Marshal(map[string]string{
				"move": move,
			})
			if err != nil {
				panic(err)
			}
			return payload
		})

		handleRoute(prefix+"/end", snake, func(snake snakes.SnakeService, ctx *snek.Context) []byte {
			snake.EndGame(ctx)

			if !ctx.Game.Dev {
				db.PruneGames()
			}

			return nil
		})
	}
}
