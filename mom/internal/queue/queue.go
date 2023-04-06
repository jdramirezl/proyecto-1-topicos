package queue;

import (
	"mom/internal/consumer"
	"mom/internal/linked_list"
)

type Queue struct {
	Messages        *linked_list.LinkedList
	Consumers       []consumer.Consumer
	CurrentConsumer int
	ConsumerMap     map[string] chan string // Why chan string? does it matter?
}

func NewQueue() *Queue {
	message_queue := linked_list.NewLinkedList()

	q := Queue{
        Messages: &message_queue,
        Consumers: []consumer.Consumer{},
        CurrentConsumer: -1,
        ConsumerMap: map[string] chan string {},
    }

	go q.consume()

	return &q
}

func (q *Queue) consume(){
	for {
		message := q.Messages.Pop()
		if message != nil {
			q.Send(*message)
		}
	}
}

func (q *Queue) Send(message string) {
	q.roundRobin()
	q.sendMessage(message)
}

func (q *Queue) roundRobin() {
	for !q.Consumers[q.CurrentConsumer].Available {
		q.CurrentConsumer = (q.CurrentConsumer + 1) % len(q.Consumers)
	}
}

func (q *Queue) sendMessage(message string) {
	q.Consumers[q.CurrentConsumer].Available = false
}

func (q *Queue) AddConsumer(address string) chan string {
	channel := make(chan string)
	cons := consumer.NewConsumer(address)
	q.ConsumerMap[cons.ID] = channel
	return channel
}

func (q *Queue) RemoveConsumer(address string) {
	delete(q.ConsumerMap, address)
}

func (q *Queue) AddMessage(message string){
	q.Messages.Add(message)
}
