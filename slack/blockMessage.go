package slack

type BlockMessage struct {
	Blocks []Block  `json:"blocks"`
	Channel string `json:"channel"`
}

type Block struct {
	Type string `json:"type"`
	Text Text `json:"text"`
	Accessory Accessory  `json:"accessory"`
}

type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Accessory struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url"`
	AltText  string `json:"alt_text"`
}