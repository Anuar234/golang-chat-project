package main

import (
	"fmt"
	"net/http"

	"github.com/Anuar234/golang-chat-project/pkg/websocket"
)

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})
}

func main() {
	fmt.Println("Anuar full stack chat project")
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}