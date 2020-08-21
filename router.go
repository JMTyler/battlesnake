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
}

func handleRoute(route string, snake snakes.SnakeService, f func(snakes.SnakeService, http.ResponseWriter, *http.Request)) {
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

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			return
		}
		if config.GetBool("debug") {
			fmt.Println("[http]", r.Method, r.RequestURI)
		}
		w.Header().Add("Content-Type", "application/json")
		f(snake, w, r)
	})
}

// TODO: Setup root paths to default to local snake.
func RouteSnakes() {
	for _, snake := range the_snakes {
		prefix := "/" + snake.GetName()
		handleRoute(prefix+"/", snake, func(snake snakes.SnakeService, w http.ResponseWriter, r *http.Request) {
			info := snake.GetInfo()
			info.APIVersion = "1"
			info.Author = "JMTyler"

			payload, _ := json.Marshal(info)
			w.Write(payload)
		})

		handleRoute(prefix+"/start", snake, func(snake snakes.SnakeService, w http.ResponseWriter, r *http.Request) {
			//			var bytes []byte
			//			n, err := r.Body.Read(bytes)
			bytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}

			var ctx *snek.Context
			if err := json.Unmarshal(bytes, &ctx); err != nil {
				panic(err)
			}
			ctx.Prepare()

			sentry.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetRequest(r)
				scope.SetUser(sentry.User{ID: ctx.Game.ID})
			})

			snake.StartGame(ctx)

			sentry.ConfigureScope(func(scope *sentry.Scope) {
				scope.Clear()
			})
		})

		handleRoute(prefix+"/move", snake, func(snake snakes.SnakeService, w http.ResponseWriter, r *http.Request) {
			bytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}

			var ctx *snek.Context
			if err := json.Unmarshal(bytes, &ctx); err != nil {
				panic(err)
			}
			ctx.Prepare()

			sentry.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetRequest(r)
				scope.SetUser(sentry.User{ID: ctx.Game.ID})
			})

			frame := db.NewFrame(ctx)
			if !ctx.Game.Dev {
				frame.Insert()
			}

			start := time.Now()
			move := snake.Move(ctx)
			duration := time.Now().Sub(start).Microseconds()

			// If move takes longer than 400ms, something is wrong.
			if duration >= 400000 {
				sentry.WithScope(func(scope *sentry.Scope) {
					scope.SetLevel(sentry.LevelWarning)
					scope.SetTag("game", fmt.Sprintf("https://play.battlesnake.com/g/%s", frame.GameID))
					scope.SetTag("turn", fmt.Sprintf("%v", frame.Turn))

					sentry.CaptureMessage("Move took unusually long to calculate.")
				})
			}

			if !ctx.Game.Dev {
				frame.Update(move, duration)
			}

			payload, err := json.Marshal(map[string]string{
				"move": move,
			})
			if err != nil {
				panic(err)
			}
			w.Write(payload)

			sentry.ConfigureScope(func(scope *sentry.Scope) {
				scope.Clear()
			})
		})

		handleRoute(prefix+"/end", snake, func(snake snakes.SnakeService, w http.ResponseWriter, r *http.Request) {
			bytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}

			var ctx *snek.Context
			if err := json.Unmarshal(bytes, &ctx); err != nil {
				panic(err)
			}
			ctx.Prepare()

			sentry.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetRequest(r)
				scope.SetUser(sentry.User{ID: ctx.Game.ID})
			})

			snake.EndGame(ctx)

			if !ctx.Game.Dev {
				db.PruneGames()
			}

			sentry.ConfigureScope(func(scope *sentry.Scope) {
				scope.Clear()
			})
		})
	}
}
