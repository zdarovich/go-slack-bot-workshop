package bot

import (
	"fmt"
	"github.com/zdarovich/go-slack-bot-workshop/slack"
	"strings"
	"time"
)

type bot struct {
	token string
}

func New(token string) *bot {
	return &bot{
		token:token,
	}
}

func (b *bot) Start() {

	msgChan := make(chan string, 0)
	defer close(msgChan)

	// start a websocket-based Real Time API session
	ws, id := slack.Connect(b.token)

	for {
		// read each incoming message
		m, err := slack.GetMessage(ws)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(m)

		// see if we're mentioned
		if m.Type == "message" && strings.HasPrefix(m.Text, "<@"+id+">") {
			// if so try to parse if
			parts := strings.Fields(m.Text)

			if len(parts) == 3 && parts[1] == "pic" && parts[2] == "seq"{
				start := time.Now()
				name, url := generateUserSequential()
				after := time.Since(start)
				fmt.Printf("time elapsed %s \n", after)


				msg := slack.NewBlock(name, url)
				msg.Channel = m.Channel
				if err := slack.PostBlockMessage(msg, b.token); err != nil {
					fmt.Println(err)
					return
				}

			} else if len(parts) == 3 && parts[1] == "pic" && parts[2] == "par" {
				start := time.Now()
				name, url := generateUserParallel()
				after := time.Since(start)
				fmt.Printf("time elapsed %s \n", after)
				msg := slack.NewBlock(name, url)
				msg.Channel = m.Channel
				if err := slack.PostBlockMessage(msg, b.token); err != nil {
					fmt.Println(err)
					return
				}

			} else {

				m.Text = fmt.Sprintf("sorry, that does not compute")
				if err := slack.PostMessage(ws, m); err != nil {
					fmt.Println(err)
					return
				}

			}
		}
	}
}

func generateUserSequential() (string, string) {
	name := getRandomName()
	pic := getRandomPictureUrl()
	return name, pic
}

func generateUserParallel() (string, string) {
	nameChan := make(chan string)
	picChan := make(chan string)

	go func() {
		name := getRandomName()
		nameChan <- name
	}()

	go func() {
		pic := getRandomPictureUrl()
		picChan <- pic
	}()

	name := <-nameChan
	pic := <-picChan

	return name, pic
}
