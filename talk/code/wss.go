package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
)

type Msg struct {
	Body string  `json:"body"`
}

func main() {
	ws, err := websocket.Dial("wss://echo.websocket.org", "", "")
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	send := Msg{Body:"Hello"}
	if err := websocket.JSON.Send(ws, &send); err != nil {
		log.Fatal(err)
	}

	var m Msg
	if err := websocket.JSON.Receive(ws, &m); err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
}
