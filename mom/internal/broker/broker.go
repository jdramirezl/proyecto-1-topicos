package broker

type Broker interface {
	Consume()
	Send(message string)
	sendMessage(message string)
	AddConsumer(address string) chan string
	RemoveConsumer(address string)
	AddMessage(message string)
	PopMessage()
	EnableConsumer(consumerIP string)
}
