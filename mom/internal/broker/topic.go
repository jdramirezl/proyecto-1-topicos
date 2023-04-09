package broker

import (
	"mom/internal/consumer"
	"mom/internal/linked_list"
)

type Topic struct {
	Messages        *linked_list.LinkedList
	Consumers       []consumer.Consumer
	CurrentConsumer int
	ConsumerMap     map[string]chan string // Why chan string? does it matter?
	Creator         string
}

func NewTopic(creator_ip string) Broker {
	messageList := linked_list.NewLinkedList()

	q := Topic{
		Messages:        &messageList,
		Consumers:       []consumer.Consumer{},
		ConsumerMap:     map[string]chan string{},
		Creator:         creator_ip,
	}

	return &q
}

func (q *Topic) Consume(){
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
		q.Consumers[q.CurrentConsumer].Available = false
		q.ConsumerMap[q.Consumers[q.CurrentConsumer].IP] <- message
	}
}

func (q *Topic) AddConsumer(address string) chan string {
	channel := make(chan string)
	cons := consumer.NewConsumer(address)
	q.ConsumerMap[cons.IP] = channel

	q.roundRobin() // When current consumer is -1

	return channel
}

func (q *Topic) RemoveConsumer(address string) {
	delete(q.ConsumerMap, address)
	if len(q.ConsumerMap) == 0 {
		q.CurrentConsumer = -1
	}
}

func (q *Topic) AddMessage(message string){
	q.Messages.Add(message)
}

func (q *Topic) PopMessage(){
	q.Messages.Pop()
}

func (q *Topic) EnableConsumer(consumerIP string){
	for _, consumer := q.Consumers {
		if consumer.IP == consumerIP {
			consumer.Available = true
			return
		}
	}
}
