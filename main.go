package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
		return
	}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}

		fmt.Println(string(p))

		err = conn.WriteMessage(messageType, p)
		if err != nil {
			return
		}
	}
}

func main() {
	http.HandleFunc("/ws", echoHandler)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	err := http.ListenAndServe(":8008", nil)
	if err != nil {
		panic("Error: " + err.Error())
	}
}
