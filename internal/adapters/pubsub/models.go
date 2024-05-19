package pubsub

// IPubSub represents abstract PubSub for topic notifications
type IPubSub interface {
	AddSubscriber(conn *SubscriberConnectionOpts) *Subscriber
	RemoveSubscriber(conn *SubscriberConnectionOpts) *Subscriber
	CountTopicSubscribers(topic string) int
	Sub(topics []string, conn *SubscriberConnectionOpts)
	Unsub(topics []string, conn *SubscriberConnectionOpts)
	Publish(topic string, msg string)
	Broadcast(topics []string, msg string)
}

// ISubscriber represents topic(-s) subscriber
type ISubscriber interface {
	GetSubID() string
	GetID() string
	GetUID() string
	AddTopic(topic string)
	AddTopics(topics []string)
	RemoveTopic(topic string)
	GetTopics() []string
	Destroy()
	Notify(msg *Message)
	Listen()
}
