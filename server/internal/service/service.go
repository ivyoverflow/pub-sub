package service

import "sync"

// PublisherSubscriber implements Publish-Subscriber pattern methods.
type PublisherSubscriber struct {
	mutex  sync.RWMutex
	subs   map[string][]chan interface{}
	closed bool
}

// NewPublisherSubscriber returns a new PublishSubscriber object.
func NewPublisherSubscriber() *PublisherSubscriber {
	ps := &PublisherSubscriber{}
	ps.subs = make(map[string][]chan interface{})

	return ps
}

// Publish func writes a message to the transmitted topic.
func (ps *PublisherSubscriber) Publish(topic string, message interface{}) {
	ps.mutex.RLock()
	defer ps.mutex.RUnlock()

	if ps.closed {
		return
	}

	for _, channel := range ps.subs[topic] {
		go func(channel chan interface{}) {
			channel <- message
		}(channel)
	}
}

// Subscribe func adds a new subscriber to the transmitted topic.
func (ps *PublisherSubscriber) Subscribe(topic string) chan interface{} {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()

	channel := make(chan interface{}, 1)
	ps.subs[topic] = append(ps.subs[topic], channel)

	return channel
}
