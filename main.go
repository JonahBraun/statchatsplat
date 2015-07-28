package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var broadcastChat = make(chan []byte)
var register = make(chan chan []byte)

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		Warn("Upgrading to websocket error:", err)
		return
	}

	name := conn.RemoteAddr().String()
	rec := make(chan []byte)

	register <- rec

	go func() {
		for {
			// TODO: creating a new msg var everytime is bad!
			msg := <-rec
			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				Talk("Websocket read error:", err)
				return
			}
		}
	}()

	// infinitely read messages
	for {
		messageType, p, err := conn.ReadMessage()

		if messageType == websocket.BinaryMessage {
			Err("Websocket sent binary message!", p)
			return
		}
		if err != nil {
			Talk("Websocket read error:", err)
			return
		}
		broadcastChat <- []byte(name + ": " + string(p))
	}
}

func main() {
	clients := make(map[chan []byte]bool)

	// broadcast to all clients
	go func() {
		for {
			msg := <-broadcastChat
			for c, _ := range clients {
				c <- msg
			}
		}
	}()

	// handle client registration
	go func() {
		for {
			c := <-register
			clients[c] = true
		}
	}()

	// TODO unregister clients

	http.HandleFunc("/ws", handleWebsocket)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	err := http.ListenAndServe(":8008", nil)
	if err != nil {
		panic("Error: " + err.Error())
	}
}

/* Logging utilities */

const (
	// terminal reset
	TR         = "\x1b[0m"
	Bright     = "\x1b[1m"
	Dim        = "\x1b[2m"
	Underscore = "\x1b[4m"
	Blink      = "\x1b[5m"
	Reverse    = "\x1b[7m"
	Hidden     = "\x1b[8m"

	FgBlack   = "\x1b[30m"
	FgRed     = "\x1b[31m"
	FgGreen   = "\x1b[32m"
	FgYellow  = "\x1b[33m"
	FgBlue    = "\x1b[34m"
	FgMagenta = "\x1b[35m"
	FgCyan    = "\x1b[36m"
	FgWhite   = "\x1b[37m"

	BgBlack   = "\x1b[40m"
	BgRed     = "\x1b[41m"
	BgGreen   = "\x1b[42m"
	BgYellow  = "\x1b[43m"
	BgBlue    = "\x1b[44m"
	BgMagenta = "\x1b[45m"
	BgCyan    = "\x1b[46m"
	BgWhite   = "\x1b[47m"
)

func Talk(v ...interface{}) {
	Log(FgCyan, "→ ", v...)
}

func Note(v ...interface{}) {
	Log(FgGreen, "✏  ", v...)
}

func Warn(v ...interface{}) {
	Log(FgYellow, "⚠  ", v...)
}

func Err(v ...interface{}) {
	Log(FgRed, "✖  ", v...)
}

func Fatal(v ...interface{}) {
	Log(FgRed+Bright, "☠  ", v...)
	os.Exit(1)
}

func Log(color string, icon string, v ...interface{}) {
	v[0] = fmt.Sprintf("%v%v%v%v", color, icon, v[0], TR)

	fmt.Println(v...)
}
