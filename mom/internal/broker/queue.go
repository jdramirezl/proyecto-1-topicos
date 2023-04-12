package broker

import (
	"fmt"

	"github.com/jdramirezl/proyecto-1-topicos/mom/internal/consumer"
	"github.com/jdramirezl/proyecto-1-topicos/mom/internal/linked_list"
)

type Queue struct {
	Messages        *linked_list.LinkedList
	Consumers       []*consumer.Consumer
	CurrentConsumer int
	ConsumerMap     map[string]chan string // Why chan string? does it matter?
	Creator         string
}

func NewQueue(creator_ip string) *Queue {
	messageList := linked_list.NewLinkedList()

	q := Queue{
		Messages:        &messageList,
		Consumers:       []*consumer.Consumer{},
		CurrentConsumer: -1,
		ConsumerMap:     map[string]chan string{},
		Creator:         creator_ip,
	}

	return &q
}

func (q *Queue) GetCreator() string {
	return q.Creator
}

func (q *Queue) GetMessages() *linked_list.LinkedList {
	return q.Messages
}

func (q *Queue) GetConsumers() []*consumer.Consumer {
	return q.Consumers
}

func (q *Queue) Consume() {
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
	fmt.Println("i was called")
	if q.CurrentConsumer == -1 && len(q.ConsumerMap) > 0 {
		q.CurrentConsumer = 0
	}

	for !q.Consumers[q.CurrentConsumer].Available {
		// fmt.Println("im stuck step daddy")
		q.CurrentConsumer = (q.CurrentConsumer + 1) % len(q.Consumers)
	}
}

func (q *Queue) sendMessage(message string) {
	q.Consumers[q.CurrentConsumer].Available = false
	fmt.Printf("i havent sent the message: %s\n", message)
	fmt.Println(q.Consumers[q.CurrentConsumer].IP)
	q.ConsumerMap[q.Consumers[q.CurrentConsumer].IP] <- message
	fmt.Println("i sent the message")
}

func (q *Queue) AddConsumer(address string) chan string {
	channel := make(chan string)
	cons := consumer.NewConsumer(address)
	q.ConsumerMap[cons.IP] = channel

	q.Consumers = append(q.Consumers, &cons)
	q.roundRobin() // When current consumer is -1

	return channel
}

func (q *Queue) RemoveConsumer(address string) {
	delete(q.ConsumerMap, address)
	if len(q.ConsumerMap) == 0 {
		q.CurrentConsumer = -1
	}
	var newConsumers []*consumer.Consumer
	for _, consumer := range q.Consumers {
		if consumer.IP == address {
			continue
		}
		newConsumers = append(newConsumers, consumer)
	}
	q.Consumers = newConsumers
}

func (q *Queue) AddMessage(message string) {
	q.Messages.Add(message)
}

func (q *Queue) PopMessage() {
	q.Messages.Pop()
}

func (q *Queue) EnableConsumer(consumerIP string) {
	fmt.Println("Arrived consumer: " + consumerIP)
	for _, consumer := range q.Consumers {
		fmt.Println("Seeing consumer: " + consumer.IP)
		if consumer.IP == consumerIP {
			fmt.Println(consumer.IP)
			fmt.Println(consumerIP)
			fmt.Println("i entered the important condition")
			consumer.Available = true
			return
		}
	}
}

func (q *Queue) GetConsumerChannel(consumerIP string) chan string {
	return q.ConsumerMap[consumerIP]
}
