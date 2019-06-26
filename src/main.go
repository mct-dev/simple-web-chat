package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func IfError(err error, handleError func()) {
	if err != nil {
		log.Println(err.Error())
		if handleError != nil {
			handleError()
		}
	}
}

var clients = make(map[*websocket.Conn]bool) // connected client sockets
var broadcast = make(chan Message)           // broadcast channel

var upgrader = websocket.Upgrader{} // used for upgrading a regular HTTP conn to a websocket

type Message struct {
	// back tick strings are "metadata" which help GO
	// serialize and deserialize the Message obj to and from JSON
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	IfError(err, nil)
	defer ws.Close()

	clients[ws] = true

	for {
		var msg Message

		// read a message from the socket and put it in `msg` object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf(" error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	fs := http.FileServer(http.Dir("public"))

	// main route
	http.Handle("/", fs)

	// web socket route
	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	log.Println("http server started on port 8000...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
