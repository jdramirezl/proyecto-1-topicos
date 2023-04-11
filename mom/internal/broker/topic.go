package broker

import (
	"github.com/jdramirezl/proyecto-1-topicos/mom/internal/consumer"
	"github.com/jdramirezl/proyecto-1-topicos/mom/internal/linked_list"
)

type Topic struct {
	Messages        *linked_list.LinkedList
	Consumers       []consumer.Consumer
	CurrentConsumer int
	ConsumerMap     map[string]chan string // Why chan string? does it matter?
	Creator         string
}

func NewTopic(creator_ip string) *Topic {
	messageList := linked_list.NewLinkedList()

	t := Topic{
		Messages:    &messageList,
		Consumers:   []consumer.Consumer{},
		ConsumerMap: map[string]chan string{},
		Creator:     creator_ip,
	}

	return &t
}

func (t *Topic) GetCreator() string {
	return t.Creator
}

func (t *Topic) GetMessages() *linked_list.LinkedList {
	return t.Messages
}

func (t *Topic) GetConsumers() *[]consumer.Consumer {
	return &t.Consumers
}

func (q *Topic) Consume() {
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

func (q *Topic) Send(message string) {
	q.sendMessage(message)
}

func (q *Topic) sendMessage(message string) {
	for _, consumer := range q.Consumers {
		consumer.Available = false
		q.ConsumerMap[consumer.IP] <- message
	}
}

func (q *Topic) AddConsumer(address string) chan string {
	channel := make(chan string)
	cons := consumer.NewConsumer(address)
	q.ConsumerMap[cons.IP] = channel
	q.Consumers = append(q.Consumers, cons)

	return channel
}

func (q *Topic) RemoveConsumer(address string) {
	delete(q.ConsumerMap, address)
	if len(q.ConsumerMap) == 0 {
		q.CurrentConsumer = -1
	}
	var newConsumers []consumer.Consumer
	for _, consumer := range q.Consumers {
		if consumer.IP == address {
			continue
		}
		newConsumers = append(newConsumers, consumer)
	}
	q.Consumers = newConsumers
}

func (q *Topic) AddMessage(message string) {
	q.Messages.Add(message)
}

func (q *Topic) PopMessage() {
	q.Messages.Pop()
}

func (q *Topic) EnableConsumer(consumerIP string) {
	for _, consumer := range q.Consumers {
		if consumer.IP == consumerIP {
			consumer.Available = true
			return
		}
	}
}

func (q *Topic) GetConsumerChannel(consumerIP string) chan string {
	return q.ConsumerMap[consumerIP]
}
