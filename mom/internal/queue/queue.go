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
	Creator 	   string
}

func NewQueue(creator_ip string) *Queue {
	message_queue := linked_list.NewLinkedList()

	q := Queue{
        Messages: &message_queue,
        Consumers: []consumer.Consumer{},
        CurrentConsumer: -1,
        ConsumerMap: map[string] chan string {},
		Creator: creator_ip,
    }

	go q.consume()

	return &q
}

func (q *Queue) consume(){
	for {
		if q.CurrentConsumer == -1 {
			continue
		}

		message := q.Messages.Pop()
		
		if message != nil {
			go q.Send(*message)
		}
	}
}

func (q *Queue) Send(message string) {
	q.roundRobin()
	q.sendMessage(message)
}

func (q *Queue) roundRobin() {
	if q.CurrentConsumer == -1 && len(q.ConsumerMap) > 0 {
		q.CurrentConsumer = 0
	}

	for !q.Consumers[q.CurrentConsumer].Available {
		q.CurrentConsumer = (q.CurrentConsumer + 1) % len(q.Consumers)
	}
}

func (q *Queue) sendMessage(message string) {
	q.Consumers[q.CurrentConsumer].Available = false
	// TODO: SEND THE MESSAGE!!
}

func (q *Queue) AddConsumer(address string) chan string {
	channel := make(chan string)
	cons := consumer.NewConsumer(address)
	q.ConsumerMap[cons.ID] = channel

	q.roundRobin() // When current consumer is -1

	return channel
}

func (q *Queue) RemoveConsumer(address string) {
	delete(q.ConsumerMap, address)
	if len(q.ConsumerMap) == 0 {
		q.CurrentConsumer = -1
	}
}

func (q *Queue) AddMessage(message string){
	q.Messages.Add(message)
}
