package slack

func NewModel() BlockMessage {
	blocks := make([]Block, 1)
	blocks[0] = Block{
		Type: "section",
		Text: Text{
			Type: "mrkdwn",
			Text: "Take a look at this image.",
		},
		Accessory:Accessory{
			Type:     "image",
			ImageURL: "https://api.slack.com/img/blocks/bkb_template_images/palmtree.png",
			AltText:  "palm tree",
		},
	}
	return BlockMessage{Blocks: blocks}
}

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
