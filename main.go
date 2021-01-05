package main

import (
	"fmt"
	"sync"
)

type publisherSubscriber struct {
	mutex  sync.RWMutex
	subs   map[string][]chan interface{}
	closed bool
}

func newPublisherSubscriber() *publisherSubscriber {
	ps := &publisherSubscriber{}
	ps.subs = make(map[string][]chan interface{})

	return ps
}

// publish func writes a message to the transmitted topic.
func (ps *publisherSubscriber) publish(topic string, message interface{}) {
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

// subscribe func adds a new subscriber to the transmitted topic.
func (ps *publisherSubscriber) subscribe(topic string) chan interface{} {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()

	channel := make(chan interface{}, 1)
	ps.subs[topic] = append(ps.subs[topic], channel)

	return channel
}

// close func closes all channels for all subscribers.
func (ps *publisherSubscriber) close() {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()

	if !ps.closed {
		ps.closed = true
		for _, subs := range ps.subs {
			for _, channel := range subs {
				close(channel)
			}
		}
	}
}

func main() {
	pubSub := newPublisherSubscriber()

	channel := make(chan interface{})
	channel = pubSub.subscribe("news")
	pubSub.publish("news", "Ohhhh...")
	fmt.Println(<-channel)
	pubSub.publish("news", "Uhhhh...")
	fmt.Println(<-channel)
	pubSub.publish("news", "Ghhhh...")
	fmt.Println(<-channel)

	channel = pubSub.subscribe("games")
	pubSub.publish("games", "Jhhhh...")
	fmt.Println(<-channel)
	pubSub.publish("games", "Ihhhh...")
	fmt.Println(<-channel)
	pubSub.publish("games", "Lhhhh...")
	fmt.Println(<-channel)
}
