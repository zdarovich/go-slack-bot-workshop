package bot

import (
	"fmt"
	"github.com/zdarovich/go-slack-bot-workshop/slack"
	"strings"
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

			if len(parts) == 2 && parts[1] == "pic" {

				go func(m slack.Message) {
					name, url := generateUser()
					msg := slack.NewBlock(name, url)
					msg.Channel = m.Channel
					if err := slack.PostBlockMessage(msg, b.token); err != nil {
						fmt.Println(err)
						return
					}

				}(m)

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

func generateUserSequentially() (string, string) {
	name := getRandomName()
	url := getRandomPictureUrl()
	return name, url
}

func generateUser() (string, string) {

}
