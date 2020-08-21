package main

import (
	"fmt"
	"github.com/JMTyler/battlesnake/_/config"
	"github.com/JMTyler/battlesnake/_/db"
	"net/http"
)

func main() {
	initSentry()
	defer recoverWithSentry()

	db.InitDatabase()
	defer db.CloseDatabase()

	RouteSnakes()

	port := config.Get("port", "80")
	go http.ListenAndServe(":"+port, nil)

	fmt.Println("🐍 listening on port", port)
	fmt.Println()

	WaitForKillSignal()
}
