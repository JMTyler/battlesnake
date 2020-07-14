package main

import (
	"fmt"
	"github.com/JMTyler/battlesnake/_/config"
	"net/http"
)

func main() {
	RouteSnakes()

	port := config.Get("port", "80")
	fmt.Println("🐍 listening on port", port)
	fmt.Println("-----")

	http.ListenAndServe(":"+port, nil)
}
