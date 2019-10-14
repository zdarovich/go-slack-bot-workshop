package bot

import (
	"encoding/csv"
	"fmt"
	"github.com/zdarovich/go-slack-bot-workshop/slack"
	"log"
	"net/http"
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
			log.Fatal(err)
		}
		fmt.Print(m)

		// see if we're mentioned
		if m.Type == "message" && strings.HasPrefix(m.Text, "<@"+id+">") {
			// if so try to parse if
			parts := strings.Fields(m.Text)

			if len(parts) == 3 && parts[1] == "stock" {

				// looks good, get the quote and reply with the result
				go func(m slack.Message) {
					m.Text = getQuote(parts[2])
					if err := slack.PostMessage(ws, m); err != nil {
						log.Fatal(err)
					}
				}(m)

			} else {

				m.Text = fmt.Sprintf("sorry, that does not compute\n")

				if err := slack.PostMessage(ws, m); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

// Get the quote via Yahoo. You should replace this method to something
// relevant to your team!
func getQuote(sym string) string {
	sym = strings.ToUpper(sym)
	url := fmt.Sprintf("http://download.finance.yahoo.com/d/quotes.csv?s=%s&f=nsl1op&e=.csv", sym)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	rows, err := csv.NewReader(resp.Body).ReadAll()
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	if len(rows) >= 1 && len(rows[0]) == 5 {
		return fmt.Sprintf("%s (%s) is trading at $%s", rows[0][0], rows[0][1], rows[0][2])
	}
	return fmt.Sprintf("unknown response format (symbol was \"%s\")", sym)
}
