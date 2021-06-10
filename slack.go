package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Slack struct {
	HookURL string
}

type SlackMessage struct {
	Text string `json:"text"`
}

func (s Slack) sendMessage(text string) {
	message := SlackMessage{
		Text: text,
	}
	requestBody, err := json.Marshal(message)
	if err != nil {
		log.Println("Failed to marshal slack message")
		return
	}

	req, err := http.NewRequest("POST", s.HookURL, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println("Failed to create slack request")
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Failed to send slack message")
	}
	defer resp.Body.Close()
}
