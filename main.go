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

func print_binary(s []byte) {
	fmt.Printf("Received b:")
	for n := 0; n < len(s); n++ {
		fmt.Printf("%d,", s[n])
	}
	fmt.Printf("\n")
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

		print_binary(p)

		err = conn.WriteMessage(messageType, []byte("Greetings browser"))
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
