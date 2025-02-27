package util

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

// Message contains the message's tag and its contained matched sentences
type Message struct {
	Tag      string   `json:"tag"`
	Messages []string `json:"messages"`
}

var messages = SerializeMessages()

// SerializeMessages serializes the content of `res/datasets/messages.json` in JSON
func SerializeMessages() (messages []Message) {
	err := json.Unmarshal(ReadFile("res/datasets/messages.json"), &messages)
	if err != nil {
		fmt.Println(err)
	}

	return messages
}

// GetMessage retrieves a message tag and returns a random message chose from res/datasets/messages.json
func GetMessage(tag string) string {
	for _, message := range messages {
		// Find the message with the right tag
		if message.Tag != tag {
			continue
		}

		// Returns the only element if there aren't more
		if len(message.Messages) == 1 {
			return message.Messages[0]
		}

		// Returns a random sentence
		return message.Messages[rand.Intn(len(message.Messages))]
	}

	return messages[0].Messages[0]
}
