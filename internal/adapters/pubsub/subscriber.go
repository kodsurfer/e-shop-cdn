package pubsub

import (
	"fmt"
	"sync"
)

var _ ISubscriber = (*Subscriber)(nil)

// OnNotifyFn callback func for notifications
type OnNotifyFn = func(conn *SubscriberConnectionOpts, msg *Message)

// SubscriberConnectionOpts represent connection options
type SubscriberConnectionOpts struct {
	CID string // ws connection uuid
	UID string // associated user id
}

// GetSubID return sub id like uuid:uid
func (s *SubscriberConnectionOpts) GetSubID() string {
	return fmt.Sprintf("%s:%s", s.CID, s.UID)
}

// Subscriber wrapper
type Subscriber struct {
	/* conn represent connection id */
	conn *SubscriberConnectionOpts
	/* messages receive changes */
	messages chan *Message
	/* online indicates active sub */
	online bool
	/* onNotify handler for notifications */
	onNotify OnNotifyFn

	mu sync.RWMutex
	/* topics keep tracking subscribed topics */
	topics map[string]struct{}
}

// NewSubscriber create new default subscriber
func NewSubscriber(conn *SubscriberConnectionOpts) *Subscriber {
	return &Subscriber{
		conn:     conn,
		messages: make(chan *Message),
		topics:   make(map[string]struct{}),
		online:   true,
		onNotify: func(conn *SubscriberConnectionOpts, msg *Message) {},
	}
}

// RegisterNotifyHandler attach notification handler
func (s *Subscriber) RegisterNotifyHandler(h OnNotifyFn) {
	s.onNotify = h
}

// GetSubID return sub id
func (s *Subscriber) GetSubID() string {
	return s.conn.GetSubID()
}

// GetUID return user id
func (s *Subscriber) GetUID() string {
	return s.conn.UID
}

// GetID return conn id
func (s *Subscriber) GetID() string {
	return s.conn.GetSubID()
}

// AddTopic sub to topic
func (s *Subscriber) AddTopic(topic string) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	s.topics[topic] = struct{}{}
}

// AddTopics sub to topics
func (s *Subscriber) AddTopics(topics []string) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, topic := range topics {
		s.topics[topic] = struct{}{}
	}
}

// RemoveTopic unsub topic
func (s *Subscriber) RemoveTopic(topic string) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	delete(s.topics, topic)
}

// GetTopics return all topics
func (s *Subscriber) GetTopics() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	topics := make([]string, 0)
	for topic, _ := range s.topics {
		topics = append(topics, topic)
	}

	return topics
}

// Destroy clear sub (topics, close channels)
func (s *Subscriber) Destroy() {
	s.mu.RLock()
	defer s.mu.RUnlock()

	s.online = false
}

// Notify sub with message
func (s *Subscriber) Notify(msg *Message) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.online {
		s.messages <- msg
	}
}

// Listen incoming messages in loop
func (s *Subscriber) Listen() {
	for {
		select {
		case msg, ok := <-s.messages:
			if ok {
				fmt.Printf("Subscriber conn id %s / uid %s, received: %s from topic: %s\n", s.conn.CID, s.conn.UID, msg.GetMessagePayload(), msg.GetTopic())
				s.onNotify(s.conn, msg)
			} else {
				break
			}
		}
	}
}
