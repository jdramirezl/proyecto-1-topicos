package mom

import (
	"errors"
	"mom/internal/broker"

	"mom/internal/cluster"
	"mom/internal/proto/message"
)

var (
	ErrBrokerNotFound = errors.New("broker not found")
)

type MomService interface {
	StartConsumption()
	CreateConnection(address string)
	DeleteConnection(address string)
	CreateTopic(name string, clientIP string)
	DeleteTopic(name string, clientIp string) error
	CreateQueue(name string, clientIP string)
	DeleteQueue(name string, clientIp string) error
	Subscribe(brokerName string, address string, messageType message.Type) error
	Unsubscribe(brokerName string, address string, messageType message.Type) error
	SendMessage(brokerName string, payload string, messageType message.Type) error
	EnableConsumer(brokerName string, consumerIP string, messageType message.Type) error
	GetBroker(brokerName string, messageType message.Type) (broker.Broker, error)
	GetConfig() *cluster.Config
}

type momService struct {
	topics      map[string]*broker.Topic
	queues      map[string]*broker.Queue
	connections []string
	Config      *cluster.Config
}

func NewMomService() MomService {
	m := &momService{
		queues:      map[string]*broker.Queue{},
		connections: []string{},
		Config:      cluster.NewConfig(),
	}
	if m.Config.IsLeader() {
		m.StartConsumption()
	}
	return m
}

func (s *momService) GetConfig() *cluster.Config {
	return s.Config
}

func (s *momService) StartConsumption() {
	for _, queue := range s.queues {
		queue.Consume()
	}
	for _, topic := range s.topics {
		topic.Consume()
	}
}

// Create a new connection and add it to the list of active connections
func (s *momService) CreateConnection(address string) {
	s.connections[address] = conn
}

// Delete a connection from the list of active connections
func (s *momService) DeleteConnection(address string) {
	delete(s.connections, address)
}

// Create a new queue and add it to the list of active queues
func (s *momService) CreateQueue(name string, clientIP string) {
	queue := broker.NewQueue(clientIP)
	s.queues[name] = queue
}

// Delete the queue with the given name
func (s *momService) DeleteQueue(name string, clientIp string) error {
	if clientIp != s.queues[name].Creator {
		return ErrBrokerNotFound
	}
	delete(s.queues, name)
	return nil
}

// Create a new queue and add it to the list of active queues
func (s *momService) CreateTopic(name string, clientIP string) {
	topic := broker.NewTopic(clientIP)
	s.queues[name] = topic
}

// Delete the queue with the given name
func (s *momService) DeleteTopic(name string, clientIp string) error {
	if clientIp != s.topics[name].Creator {
		return ErrBrokerNotFound
	}
	delete(s.topics, name)
	return nil
}

// Add a subscriber to a queue
func (s *momService) Subscribe(brokerName string, address string, messageType message.Type) error {
	broker, err := s.GetBroker(brokerName, messageType)
	if err != nil {
		return err
	}
	broker.AddConsumer(address)
}

// Remove a subscriber from a queue
func (s *momService) Unsubscribe(brokerName string, address string, messageType message.Type) error {
	broker, err := s.GetBroker(brokerName, messageType)
	if err != nil {
		return err
	}
	broker.RemoveConsumer(address)
	return nil
}

// Send a message to a queue or topic
func (s *momService) SendMessage(brokerName string, payload string, messageType message.Type) error {
	broker, err := s.GetBroker(brokerName, messageType)
	if err != nil {
		return err
	}
	// sincronizar con replicas
	broker.AddMessage(payload)
	return nil
}

func (s *momService) EnableConsumer(brokerName string, consumerIP string, messageType message.Type) error {
	broker, err := s.GetBroker(brokerName, messageType)
	if err != nil {
		return err
	}
	broker.EnableConsumer(consumerIP)
	return nil
}

func (s *momService) GetBroker(brokerName string, messageType message.Type) (broker.Broker, error) {
	if messageType == message.Type_QUEUE {
		for name, queue := range s.queues {
			if name == brokerName {
				return queue, nil
			}
		}
	} else {
		for name, topic := range s.topics {
			if name == brokerName {
				return topic, nil
			}
		}
	}
	return nil, ErrBrokerNotFound
}
