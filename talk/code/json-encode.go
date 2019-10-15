package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Response struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Region  string `json:"region"`
}

func main() {
	msg := Response{Name:"Nicolaie", Surname:"Timofti", Gender:"male", Region:"Romania"}

	if b, err := json.Marshal(msg); err != nil {
		log.Fatal(err)

	} else {
		fmt.Println(string(b))
	}
}
