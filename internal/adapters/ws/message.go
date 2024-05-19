package ws

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

// MessagePayload from websocket client
type MessagePayload struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

// MessageData from websocket client
type MessageData struct {
	Payload   *MessagePayload `json:"payload"`
	CreatedAt time.Time       `json:"created_at"`
}

// NewMessageData create new text message
func NewMessageData(message string) (MessageData, error) {
	payload := &MessagePayload{}
	if err := json.Unmarshal([]byte(message), payload); err != nil {
		return MessageData{}, err
	}

	return MessageData{
		Payload:   payload,
		CreatedAt: time.Now(),
	}, nil
}

// Message for clients
type Message struct {
	ID          string      `json:"id"`
	From        string      `json:"from"`
	To          string      `json:"to"`
	Content     MessageData `json:"content"`
	IsBroadcast bool        `json:"-"`
}

// NewMessage for subs
func NewMessage(from string, content MessageData) Message {
	return Message{
		ID:      uuid.New().String(),
		From:    from,
		Content: content,
	}
}

// NewBroadcastMessage for subs
func NewBroadcastMessage(content MessageData) Message {
	return Message{
		ID:          uuid.New().String(),
		Content:     content,
		IsBroadcast: true,
	}
}
