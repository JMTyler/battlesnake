package main

import (
	"encoding/json"
	"fmt"
	snek "github.com/JMTyler/battlesnake/_"
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
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			return
		}
		fmt.Println("[http]", r.Method, r.RequestURI)
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

			var ctx snek.Context
			if err := json.Unmarshal(bytes, &ctx); err != nil {
				panic(err)
			}

			snake.StartGame(ctx)
		})

		handleRoute(prefix+"/move", snake, func(snake snakes.SnakeService, w http.ResponseWriter, r *http.Request) {
			bytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}

			var ctx snek.Context
			if err := json.Unmarshal(bytes, &ctx); err != nil {
				panic(err)
			}

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
			fmt.Printf("Move took %vμs.\n", duration)

			payload, err := json.Marshal(map[string]string{
				"move": move,
			})
			if err != nil {
				panic(err)
			}
			w.Write(payload)
		})

		handleRoute(prefix+"/end", snake, func(snake snakes.SnakeService, w http.ResponseWriter, r *http.Request) {
			bytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}

			var ctx snek.Context
			if err := json.Unmarshal(bytes, &ctx); err != nil {
				panic(err)
			}

			snake.EndGame(ctx)
		})
	}
}
