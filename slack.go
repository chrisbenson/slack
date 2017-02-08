package slack

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Message struct {
	Channel   string `json:"channel, omitempty"`
	Username  string `json:"username, omitempty"`
	Text      string `json:"text, omitempty"`
	IconEmoji string `json:"icon_emoji, omitempty"`
}

func Send(url string, m Message) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(m)
	_, err := http.Post(url, "application/json; charset=utf-8", b)
	if err != nil {
		return err
	}
	return nil
}
