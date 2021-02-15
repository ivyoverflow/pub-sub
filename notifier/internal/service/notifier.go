// Package service cotains Publisher-Subscriber pattern implementation.
package service

import "sync"

// Notifier implements Publish-Subscriber pattern methods.
type Notifier struct {
	mutex sync.RWMutex
	subs  map[string][]chan interface{}
}

// NewNotifier returns a new PublishSubscriber object.
func NewNotifier() *Notifier {
	n := &Notifier{}
	n.subs = make(map[string][]chan interface{})

	return n
}

// Publish func writes a message to the transmitted topic.
func (n *Notifier) Publish(topic string, message interface{}) {
	n.mutex.RLock()
	defer n.mutex.RUnlock()

	for _, channel := range n.subs[topic] {
		go func(channel chan interface{}) {
			channel <- message
		}(channel)
	}
}

// Subscribe func adds a new subscriber to the transmitted topic.
func (n *Notifier) Subscribe(topic string) chan interface{} {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	channel := make(chan interface{}, 1)
	n.subs[topic] = append(n.subs[topic], channel)

	return channel
}
