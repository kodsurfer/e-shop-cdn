package pubsub

// Message represent text notification for topic
type Message struct {
	topic   string // topics kind a like /a/b/* or a/b/file.txt
	payload string // base64 encoded data or json string (probably)
}

// NewMessage create new text message for topic
func NewMessage(topic, payload string) *Message {
	return &Message{
		topic:   topic,
		payload: payload,
	}
}

// GetTopic return topic name
func (m Message) GetTopic() string {
	return m.topic
}

// GetMessagePayload return message payload
func (m Message) GetMessagePayload() string {
	return m.payload
}
