package main

import (
	"encoding/json"
	"fmt"
	snek "github.com/JMTyler/battlesnake/_"
	"net/http"
	"time"
	//	"github.com/JMTyler/battlesnake/_/utils"
	//	"github.com/JMTyler/battlesnake/snakes"
)

type SnakeService interface {
	GetName() string
	GetInfo() map[string]string
	StartGame(snek.Context)
	Move(snek.Context) string
	EndGame(snek.Context)
}

var snakes = []SnakeService{
	//	snakes.Local,
	//	snakes.Rufio,
	//	snakes.Proxy,
}

func handleRoute(route string, f func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			return
		}
		fmt.Println("[http]", r.Method, r.RequestURI)
		w.Header().Add("Content-Type", "application/json")
		f(w, r)
	})
}

// TODO: Setup root paths to default to local snake.
func RouteSnakes() {
	for _, snake := range snakes {
		prefix := "/" + snake.GetName()
		handleRoute(prefix, func(w http.ResponseWriter, r *http.Request) {
			info := snake.GetInfo()
			info["apiversion"] = "1"
			info["author"] = "JMTyler"

			payload, _ := json.Marshal(info)
			w.Write(payload)
		})

		handleRoute(prefix+"/start", func(w http.ResponseWriter, r *http.Request) {
			var bytes []byte
			if _, err := r.Body.Read(bytes); err != nil {
				panic(err)
			}

			var ctx snek.Context
			if err := json.Unmarshal(bytes, &ctx); err != nil {
				panic(err)
			}

			snake.StartGame(ctx)
		})

		handleRoute(prefix+"/move", func(w http.ResponseWriter, r *http.Request) {
			var bytes []byte
			if _, err := r.Body.Read(bytes); err != nil {
				panic(err)
			}
			var ctx snek.Context
			if err := json.Unmarshal(bytes, &ctx); err != nil {
				panic(err)
			}

			start := time.Now()
			//			await utils.RecordFrame(req.body);
			move := snake.Move(ctx)
			duration := time.Now().Sub(start).Milliseconds()
			//			await utils.RecordFrame(req.body, { move, duration });
			fmt.Printf("Move took %vms.\n", duration)

			payload, err := json.Marshal(map[string]string{
				"move": move,
			})
			if err != nil {
				panic(err)
			}
			w.Write(payload)
		})

		handleRoute(prefix+"/end", func(w http.ResponseWriter, r *http.Request) {
			var bytes []byte
			if _, err := r.Body.Read(bytes); err != nil {
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
