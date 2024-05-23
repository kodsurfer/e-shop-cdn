package pubsub

import (
	"sync"
)

// HINT: concurrency safe maps

// safeTopics topics pool
type safeTopics struct {
	sync.RWMutex
	// List of topics
	topics map[string]struct{}
}

func (p *safeTopics) set(topic string) {
	p.Lock()
	defer p.Unlock()
	p.topics[topic] = struct{}{}
}

func (p *safeTopics) all() []string {
	p.RLock()
	defer p.RUnlock()
	ret := make([]string, 0)
	for topic, _ := range p.topics {
		ret = append(ret, topic)
	}
	return ret
}

func (p *safeTopics) contains(topic string) bool {
	p.RLock()
	_, ok := p.topics[topic]
	p.RUnlock()
	return ok
}

func (p *safeTopics) delete(topic string) {
	p.Lock()
	defer p.Unlock()
	delete(p.topics, topic)
}

//nolint:all
func (p *safeTopics) reset() {
	p.Lock()
	defer p.Unlock()
	p.topics = make(map[string]struct{})
}

// Subscribers map of subs like map[id:Subscriber]
type Subscribers map[string]*Subscriber

// Topics map of sub's topics like map[topic:map[id:Subscriber]]
type Topics map[string]Subscribers

// safePubSubs subs pool
type safePubSubs struct {
	sync.RWMutex
	// List of subs
	subs Subscribers
}

func (p *safePubSubs) set(s *Subscriber) {
	p.Lock()
	defer p.Unlock()
	p.subs[s.GetSubID()] = s
}

func (p *safePubSubs) get(sid string) *Subscriber {
	p.Lock()
	defer p.Unlock()

	s, ok := p.subs[sid]
	if !ok {
		return nil
	}

	return s
}

func (p *safePubSubs) delete(s *Subscriber) {
	p.Lock()
	defer p.Unlock()
	delete(p.subs, s.GetSubID())
}

func (p *safePubSubs) contains(sid string) bool {
	p.RLock()
	defer p.RUnlock()
	_, ok := p.subs[sid]
	return ok
}

// safePubTopics topic subs pool
type safePubTopics struct {
	sync.RWMutex
	// List of topics
	topics Topics
}

func (p *safePubTopics) addSubIfNotExists(topic string, s *Subscriber) {
	p.Lock()
	defer p.Unlock()

	if p.topics[topic] == nil {
		p.topics[topic] = Subscribers{}
	}

	p.topics[topic][s.GetSubID()] = s
}

func (p *safePubTopics) count(topic string) int {
	p.Lock()
	defer p.Unlock()

	return len(p.topics[topic])
}

func (p *safePubTopics) removeSub(topic string, s *Subscriber) {
	p.Lock()
	defer p.Unlock()
	delete(p.topics[topic], s.GetSubID())
}

func (p *safePubTopics) subs(topic string) []*Subscriber {
	p.RLock()
	defer p.RUnlock()

	ret := make([]*Subscriber, 0)

	subs, ok := p.topics[topic]
	if !ok {
		return []*Subscriber{}
	}

	for _, subscriber := range subs {
		ret = append(ret, subscriber)
	}

	return ret
}

func (p *safePubTopics) all() []*Subscriber {
	p.RLock()
	defer p.RUnlock()

	ret := make([]*Subscriber, 0)
	for topic, _ := range p.topics {
		for _, s := range p.topics[topic] {
			ret = append(ret, s)
		}
	}

	return ret
}
