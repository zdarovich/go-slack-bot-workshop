package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Message struct {
	Id      uint64 `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}
const stream = `
	{
	  "id": 0,
	  "type": "message",
	  "channel": "dasdas",
	  "text": "palm tree"
	}
`
func main() {
	var response Message
	err := json.Unmarshal([]byte(stream), response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}
