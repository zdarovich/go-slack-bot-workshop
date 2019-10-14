package main

import (
	"fmt"
	"github.com/zdarovich/go-slack-bot-workshop/bot"
	"github.com/zdarovich/go-slack-bot-workshop/config"
)

func main() {
	c, err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	token := c.Token
	if len(token) == 0 {
		fmt.Println("token is empty")
	}
	bot := bot.New(token)
	bot.Start()
}