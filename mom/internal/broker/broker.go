package broker

import (
	"github.com/jdramirezl/proyecto-1-topicos/mom/internal/consumer"
	"github.com/jdramirezl/proyecto-1-topicos/mom/internal/linked_list"
)

type Broker interface {
	Consume()
	Send(message string)
	sendMessage(message string)
	GetMessages() *linked_list.LinkedList
	AddConsumer(address string) chan string
	RemoveConsumer(address string)
	GetConsumers() *[]consumer.Consumer
	AddMessage(message string)
	PopMessage()
	EnableConsumer(consumerIP string)
	GetConsumerChannel(consumerIP string) chan string
	GetCreator() string
}
