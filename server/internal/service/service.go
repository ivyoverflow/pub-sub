package service

// PublisherSubscriberI ...
type PublisherSubscriberI interface {
	Publish(topic string, message interface{})
	Subscribe(topic string) chan interface{}
}
