package pubsub

import (
	"fmt"
	"sync"
)

var _ IPubSub = (*PubSub)(nil)

// Subscribers map of subs like map[id:Subscriber]
type Subscribers map[string]*Subscriber

// Topics map of sub's topics like map[topic:map[id:Subscriber]]
type Topics map[string]Subscribers

// PubSub implement abstraction for notifications
type PubSub struct {
	subs   Subscribers
	topics Topics

	mu sync.RWMutex
}

// NewPubSub create default PubSub
func NewPubSub() *PubSub {
	return &PubSub{
		subs:   make(Subscribers),
		topics: make(Topics),
	}
}

// AddSubscriber by unique id (for example, ws for connection could be "conn_id:user_id")
func (b *PubSub) AddSubscriber(conn *SubscriberConnectionOpts) *Subscriber {
	b.mu.Lock()
	defer b.mu.Unlock()

	s := NewSubscriber(conn)

	b.subs[s.GetSubID()] = s

	return s
}

// RemoveSubscriber remove by unique id
func (b *PubSub) RemoveSubscriber(conn *SubscriberConnectionOpts) *Subscriber {
	s, ok := b.subs[conn.GetSubID()]
	if !ok {
		return nil
	}

	for topic := range s.topics {
		b.Unsub([]string{topic}, &SubscriberConnectionOpts{
			CID: s.GetID(),
			UID: s.GetUID(),
		})
	}

	b.mu.Lock()
	delete(b.subs, s.GetSubID())
	b.mu.Unlock()

	s.Destroy()

	return s
}

// CountTopicSubscribers count of active topic subs
func (b *PubSub) CountTopicSubscribers(topic string) int {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return len(b.topics[topic])
}

// Sub user to topic
func (b *PubSub) Sub(topics []string, conn *SubscriberConnectionOpts) {
	b.mu.Lock()
	defer b.mu.Unlock()

	s, ok := b.subs[conn.GetSubID()]
	if !ok {
		return
	}

	s.AddTopics(topics)

	for _, topic := range topics {
		if b.topics[topic] == nil {
			b.topics[topic] = Subscribers{}
		}

		b.topics[topic][conn.GetSubID()] = s

		fmt.Printf("%s Subscribed for topic: %s\n", s.GetSubID(), topic)
	}
}

// Unsub user to topic
func (b *PubSub) Unsub(topics []string, conn *SubscriberConnectionOpts) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	s, ok := b.subs[conn.GetSubID()]
	if !ok {
		return
	}

	for _, topic := range topics {
		delete(b.topics[topic], s.GetSubID())
		s.RemoveTopic(topic)
		fmt.Printf("%s Unsubscribed for topic: %s\n", s.GetSubID(), topic)
	}
}

// Publish text message to topic
func (b *PubSub) Publish(topic string, msg string) {
	b.mu.RLock()
	bTopics := b.topics[topic]
	b.mu.RUnlock()

	for _, s := range bTopics {
		m := NewMessage(msg, topic)

		go (func(s *Subscriber) {
			s.Notify(m)
		})(s)
	}
}

// Broadcast publish msg to topics
func (b *PubSub) Broadcast(topics []string, msg string) {
	for _, topic := range topics {
		for _, s := range b.topics[topic] {
			m := NewMessage(msg, topic)

			go (func(s *Subscriber) {
				s.Notify(m)
			})(s)
		}
	}
}
