package consumer

type Consumer struct {
	Available bool
	IP        string
}

func NewConsumer(address string) Consumer {
	return Consumer{Available: true, IP: address}
}
