package main

import (
	"fmt"
	"github.com/JMTyler/battlesnake/_/config"
	"net/http"
)

func main() {
	RouteSnakes()

	port := config.Get("PORT", 80).(string)
	fmt.Println("🐍 listening on port", port)
	fmt.Println("-----")

	http.ListenAndServe(":"+port, nil)
}
