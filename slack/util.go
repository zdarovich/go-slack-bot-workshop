package slack

// block message type in slack
func NewBlock(name, url string) BlockMessage {
	blocks := make([]Block, 1)
	blocks[0] = Block{
		Type: "section",
		Text: Text{
			Type: "mrkdwn",
			Text: name,
		},
		Accessory:Accessory{
			Type:     "image",
			ImageURL: url,
			AltText:  "palm tree",
		},
	}
	return BlockMessage{Blocks: blocks}
}
