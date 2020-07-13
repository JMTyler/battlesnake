package main

import (
	"fmt"
	"net/http"
	"os"
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

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "80"
	}
	fmt.Println("Battlesnake listening on port", port)
	fmt.Println("-----")

	http.ListenAndServe(":"+port, nil)
}
