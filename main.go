package main

import (
	"fmt"
	"github.com/JMTyler/battlesnake/_/config"
	"github.com/JMTyler/battlesnake/_/db"
	"net/http"
)

func main() {
	db.InitDatabase()
	defer db.CloseDatabase()

	RouteSnakes()

	port := config.Get("port", "80")
	fmt.Println("🐍 listening on port", port)
	fmt.Println()

	http.ListenAndServe(":"+port, nil)
}
