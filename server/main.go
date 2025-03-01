package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func checkOrigin(r *http.Request) bool {
	origin := r.Header.Get("Origin")
	return origin == "http://localhost:5173"
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     checkOrigin,
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println()
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	if err := conn.WriteMessage(websocket.TextMessage, []byte("Yo!")); err != nil {
		log.Println(err)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.HandleFunc("GET /", handleIndex)

	log.Fatal(server.ListenAndServe())
}
