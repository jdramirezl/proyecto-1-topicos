package mom

import (
	"fmt"
	"mom/internal/proto/message"
	"mom/internal/queue"
	"mom/internal/cluster"
)

var (
	n_client int64 = 0
)

type connection struct {
	id      int64
	address string
	// add other relevant metadata here
}

type MomService interface {
	SendMessage(brokerName string, payload string, messageType message.Type) error
}

type momService struct {
	topics      map[string]*topic.Topic
	queues      map[string]*queue.Queue
	connections map[string]*connection
	config 	cluster.Config
}

func NewMomService() MomService {
	return &momService{
		queues:      map[string]*queue.Queue{},
		connections: map[string]*connection{},
		config: cluster.NewConfig(),
	}
}

// Create a new connection and add it to the list of active connections
func (s *momService) CreateConnection(address string) {
	conn := &connection{
		id:      n_client,
		address: address,
	}

	n_client += 1

	s.connections[address] = conn
}

// Delete a connection from the list of active connections
func (s *momService) DeleteConnection(address string) {
	delete(s.connections, address)
}

// Get a list of all active connections
func (s *momService) GetConnections() [][]string {
	conns := make([][]string, len(s.connections))

	for key, value := range s.connections {
		conns = append(conns, []string{fmt.Sprint(value.id), key})
	}

	return conns
}

// Create a new queue and add it to the list of active queues
func (s *momService) CreateQueue(name string, creator_ip string) {
	queue := queue.NewQueue(creator_ip)
	s.queues[name] = queue
}

// Delete the queue with the given name
func (s *momService) DeleteQueue(name string,  user_ip string) {
	if user_ip != s.queues[name].Creator {
		return
	}
	delete(s.queues, name)
}

// Add a subscriber to a queue
func (s *momService) Subscribe(queueName string, address string) {
	for name, queue := range s.queues {
		if name == queueName {
			queue.AddConsumer(address)
			break
		}
	}
}

// Remove a subscriber from a queue
func (s *momService) Unsubscribe(queueName string, address string) {
	for name, queue := range s.queues {
		if name == queueName {
			queue.RemoveConsumer(address)
			break
		}
	}
}

// Send a message to a queue or topic
func (s *momService) SendMessage(brokerName string, payload string, messageType message.Type) error {
	if messageType == message.Type_QUEUE {
		for name, queue := range s.queues {
			if name == brokerName {
				// TODO sync with replicas
				queue.AddMessage(payload)
				break
			}
		}
	} else {
		for name, topic := range s.topics {
			if name == brokerName {
				// TODO sync with replicas
				topic.AddMessage(payload)
				break
			}
		}
	}
	return nil
}
