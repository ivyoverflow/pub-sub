package main

import (
	"sync"
)

// publisherI interface describes all methods for publisher.
type publisherI interface {
	publish(topic string, message interface{})
	close()
}

// subscriberI interface describes all methods for subscriber.
type subscriberI interface {
	subscribe(topic string)
}

// publisher struct implements all methods from publisherI interface.
type publisher struct {
	mutex  sync.RWMutex
	subs   map[string][]chan interface{}
	closed bool
}

// subscriber struct implements all methods from subscriber interface.
type subscriber struct {
	pub *publisher
}

// pubSub struct contains implementations for both publisherI and subsriberI interfaces.
type pubSub struct {
	publisherI
	subscriberI
}

// newPublisher func returns configured publisher object.
func newPublisher() *publisher {
	pub := &publisher{}
	pub.subs = make(map[string][]chan interface{})

	return pub
}

// publish func writes a message to the transmitted topic.
func (pub *publisher) publish(topic string, message interface{}) {
	pub.mutex.RLock()
	defer pub.mutex.RUnlock()

	if pub.closed {
		return
	}

	for _, channel := range pub.subs[topic] {
		go func(channel chan interface{}) {
			channel <- message
		}(channel)
	}
}

// close func closes all channels for all subscribers.
func (pub *publisher) close() {
	pub.mutex.Lock()
	defer pub.mutex.Unlock()

	if !pub.closed {
		pub.closed = true
		for _, subs := range pub.subs {
			for _, channel := range subs {
				close(channel)
			}
		}
	}
}

// newSubscriber func returns configured subcsriber object.
func newSubscriber() *subscriber {
	pub := newPublisher()

	return &subscriber{pub}
}

// subscribe func adds a new subscriber to the transmitted topic.
func (sub *subscriber) subscribe(topic string) {
	sub.pub.mutex.Lock()
	defer sub.pub.mutex.Unlock()

	channel := make(chan interface{}, 1)
	sub.pub.subs[topic] = append(sub.pub.subs[topic], channel)
}

// newPubSub returns configured pubSub object.
func newPubSub() *pubSub {
	return &pubSub{
		publisherI:  newPublisher(),
		subscriberI: newSubscriber(),
	}
}

func main() {
	pubSub := newPubSub()

	pubSub.subscribe("news")
	pubSub.subscribe("news")
	pubSub.subscribe("news")

	pubSub.subscribe("games")
	pubSub.subscribe("games")
	pubSub.subscribe("games")

	pubSub.publish("news", "...")
	pubSub.publish("sports", "...")
	pubSub.publish("games", "...")
	pubSub.publish("movies", "...")
}
