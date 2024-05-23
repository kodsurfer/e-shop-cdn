package pubsub

import (
	"github.com/WildEgor/e-shop-gopack/pkg/libs/logger/models"
	"log/slog"
)

var _ IPubSub = (*PubSub)(nil)

// PubSub implement abstraction for notifications
type PubSub struct {
	subs   safePubSubs
	topics safePubTopics
}

// NewPubSub create default PubSub
func NewPubSub() *PubSub {
	return &PubSub{
		subs: safePubSubs{
			subs: make(Subscribers),
		},
		topics: safePubTopics{
			topics: make(Topics),
		},
	}
}

// AddSubscriber by unique id (for example, ws for connection could be "conn_id:user_id")
func (b *PubSub) AddSubscriber(conn *SubscriberConnectionOpts) *Subscriber {
	s := NewSubscriber(conn)
	b.subs.set(s)
	return s
}

// RemoveSubscriber remove by unique id
func (b *PubSub) RemoveSubscriber(conn *SubscriberConnectionOpts) *Subscriber {
	s := b.subs.get(conn.GetSubID())
	if s == nil {
		return nil
	}

	b.subs.delete(s)

	b.Unsub(s.topics.all(), &SubscriberConnectionOpts{
		CID: s.GetID(),
		UID: s.GetUID(),
	})

	s.Destroy()

	return s
}

// CountTopicSubscribers count of active topic subs
func (b *PubSub) CountTopicSubscribers(topic string) int {
	return b.topics.count(topic)
}

// Sub user to topic
func (b *PubSub) Sub(topics []string, conn *SubscriberConnectionOpts) {
	s := b.subs.get(conn.GetSubID())
	if s == nil {
		return
	}

	s.AddTopics(topics)

	for _, topic := range topics {
		slog.Debug("subscriber subs to topic", models.LogEntryAttr(&models.LogEntry{
			Props: map[string]interface{}{
				"sid":   s.GetSubID(),
				"topic": topic,
			},
		}))

		b.topics.addSubIfNotExists(topic, s)
	}
}

// Unsub user to topic
func (b *PubSub) Unsub(topics []string, conn *SubscriberConnectionOpts) {
	s := b.subs.get(conn.GetSubID())
	if s == nil {
		return
	}

	for _, topic := range topics {
		slog.Debug("unsub sub for topic", models.LogEntryAttr(&models.LogEntry{
			Props: map[string]interface{}{
				"sid":   s.GetSubID(),
				"topic": topic,
			},
		}))

		b.topics.removeSub(topic, s)
		s.RemoveTopic(topic)
	}
}

// Publish text message to topic
func (b *PubSub) Publish(topics []string, msg string) {
	for _, topic := range topics {
		for _, s := range b.topics.subs(topic) {
			m := NewMessage(msg, topic)

			go (func(s *Subscriber) {
				s.Notify(m)
			})(s)
		}
	}
}

// Broadcast publish msg to all
func (b *PubSub) Broadcast(msg string) {
	// TODO
}
