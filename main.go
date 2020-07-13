package main

import (
	"fmt"
	"github.com/JMTyler/battlesnake/_/config"
	"net/http"
	"runtime"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	if r.Method == "OPTIONS" {
		return
	}

	fmt.Println("[http]", r.Method, r.RequestURI)

	res := fmt.Sprintf("🐍 %s %s (running %s)", r.Method, r.URL.Path, runtime.Version())
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(res))
}

func main() {
	http.HandleFunc("/", rootHandler)

	port := config.Get("PORT", 80).(string)
	fmt.Println("Battlesnake listening on port", port)
	fmt.Println("-----")

	http.ListenAndServe(":"+port, nil)
}
