package pubsub

import (
	"fmt"
	"github.com/WildEgor/e-shop-gopack/pkg/libs/logger/models"
	"log/slog"
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
	/* topics keep tracking subscribed topics */
	topics safeTopics
}

// NewSubscriber create new default subscriber
func NewSubscriber(conn *SubscriberConnectionOpts) *Subscriber {
	return &Subscriber{
		conn:     conn,
		messages: make(chan *Message),
		topics: safeTopics{
			topics: make(map[string]struct{}),
		},
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

// GetID return topics id
func (s *Subscriber) GetID() string {
	return s.conn.GetSubID()
}

// AddTopic sub to topic
func (s *Subscriber) AddTopic(topic string) {
	s.topics.set(topic)
}

// AddTopics sub to topics
func (s *Subscriber) AddTopics(topics []string) {
	for _, topic := range topics {
		s.topics.set(topic)
	}
}

// RemoveTopic unsub topic
func (s *Subscriber) RemoveTopic(topic string) {
	s.topics.delete(topic)
}

// GetTopics return all topics
func (s *Subscriber) GetTopics() []string {
	return s.topics.all()
}

// Destroy clear sub (topics, close channels)
func (s *Subscriber) Destroy() {
	s.topics.reset()
	s.online = false
}

// Notify sub with message
func (s *Subscriber) Notify(msg *Message) {
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
				slog.Debug("subscriber message", models.LogEntryAttr(&models.LogEntry{
					Props: map[string]interface{}{
						"cid":     s.conn.CID,
						"uid":     s.conn.UID,
						"topic":   msg.topic,
						"payload": msg.payload,
					},
				}))
				
				s.onNotify(s.conn, msg)
			} else {
				break
			}
		}
	}
}
