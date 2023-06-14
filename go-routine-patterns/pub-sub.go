package main

import "fmt"

type Message struct {
	Topic   string
	Payload string
}

type Subscriber struct {
	Id       int
	Topic    string
	Messages chan Message
}

type Broker struct {
	subscribers map[string][]Subscriber
}

func (b *Broker) Subscribe(subscriber Subscriber) {
	b.subscribers[subscriber.Topic] = append(b.subscribers[subscriber.Topic], subscriber)
}

func (b *Broker) Publish(message Message) {
	subscribers := b.subscribers[message.Topic]
	for _, sub := range subscribers {
		sub.Messages <- message
	}
}

func initPubSub() {
	fmt.Println("initializing pub-sub")

	broker := &Broker{
		subscribers: make(map[string][]Subscriber),
	}

	sub1 := Subscriber{
		Id:       1,
		Topic:    "t1",
		Messages: make(chan Message),
	}

	sub2 := Subscriber{
		Id:       2,
		Topic:    "t2",
		Messages: make(chan Message),
	}

	broker.Subscribe(sub1)
	broker.Subscribe(sub2)

	go func() {
		for {
			select {
			case msg := <-sub1.Messages:
				fmt.Printf("Subscriber %d received message on topic %s: %s\n", sub1.Id, msg.Topic, msg.Payload)
			case msg := <-sub2.Messages:
				fmt.Printf("Subscriber %d received message on topic %s: %s\n", sub2.Id, msg.Topic, msg.Payload)
			}
		}
	}()

	broker.Publish(Message{
		Topic: "t1", Payload: "Hi from 1",
	})
	broker.Publish(Message{
		Topic: "t2", Payload: "Hi from 2",
	})
	return
}
