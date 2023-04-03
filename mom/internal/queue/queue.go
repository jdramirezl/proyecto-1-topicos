package queue;

import (
	"mom/internal/consumer"
	"mom/internal/linked_list"
)

type Queue struct {
	Messages        *linked_list.LinkedList
	Consumers       []consumer.Consumer
	CurrentConsumer int
	ConsumerMap     map[string]chan string
}

func (q *Queue) consume(){
	for {
		// Leer mensaje y si hay llamar al "send()"
	}
}

func (q *Queue) roundRobin() {
	for !q.Consumers[q.CurrentConsumer].Available {
		q.CurrentConsumer = (q.CurrentConsumer + 1) % len(q.Consumers)
	}
}

func (q *Queue) sendMessage() {
	message, err := q.Messages.Peek()
	if err != nil {
		return
	}
	//q.Consumers[q.CurrentConsumer].ReceiveMessage(message)
	q.Messages.Pop()
	q.Consumers[q.CurrentConsumer].Available = false
}

func (q *Queue) unicast() {
	q.roundRobin()
	q.sendMessage()
}

func (q *Queue) addConsumer() chan string {
	channel := make(chan string)
	cons := consumer.NewConsumer()
	q.ConsumerMap[cons.ID] = channel
}

func (q *Queue) RemoveConsumer(id string) {
	delete(q.ConsumerMap, id)
}

func (q *Queue) addMessage(){
	// Receive message from brojer and add to list
}

func (q *Queue) Send() {
	q.unicast()
}