package bot

import (
	"fmt"
	"github.com/zdarovich/go-slack-bot-workshop/slack"
	"log"
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
		//read each incoming message
		m, err := slack.GetMessage(ws)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(m)

		parts := strings.Fields(m.Text)
		botQuote := fmt.Sprintf("<@%s>", id)

		// see if we're mentioned && message has exactly 3 elements && first element is pic
		if strings.HasPrefix(m.Text, botQuote) && len(parts) == 3 && parts[1] == "pic" {

			switch parts[2] {

			case "seq":

				name, url := generateUserSequential()
				msg := slack.NewBlock(name, url)
				msg.Channel = m.Channel
				if err := slack.PostBlockMessage(msg, b.token); err != nil {
					log.Fatal(err)
					return
				}
				break


			case "par":

				name, url := generateUserParallel()
				msg := slack.NewBlock(name, url)
				msg.Channel = m.Channel
				if err := slack.PostBlockMessage(msg, b.token); err != nil {
					log.Fatal(err)
					return
				}
				break
			default:
				m.Text = fmt.Sprintf("sorry, that does not compute")
				if err := slack.PostMessage(ws, m); err != nil {
					log.Fatal(err)
					return
				}

			}
		}
	}
}

