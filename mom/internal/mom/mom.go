package mom

import (
	"context"
	"errors"
	"jdramirezl/proyecto-1-topicos/mom/internal/broker"

	"jdramirezl/proyecto-1-topicos/mom/internal/cluster"
	proto_cluster "jdramirezl/proyecto-1-topicos/mom/internal/proto/cluster"
	"jdramirezl/proyecto-1-topicos/mom/internal/proto/message"
)

var (
	ErrBrokerNotFound = errors.New("broker not found")
	ErrNotConnected   = errors.New("user is not connected, action forbidden")
	ErrNotSubscribed  = errors.New("user is not subscribed, action forbidden")
	ErrSystemExists   = errors.New("system with the same name already exists")
)

type MomService interface {
	StartConsumption()
	CreateConnection(address string)
	DeleteConnection(address string)
	GetConnections() []string
	CreateTopic(name string, clientIP string) error
	DeleteTopic(name string, clientIp string) error
	GetTopics() map[string]*broker.Topic
	CreateQueue(name string, clientIP string) error
	DeleteQueue(name string, clientIp string) error
	GetQueues() map[string]*broker.Queue
	Subscribe(brokerName string, address string, systemType proto_cluster.Type) error
	Unsubscribe(brokerName string, address string, systemType proto_cluster.Type) error
	SendMessage(brokerName string, payload string, messageType message.Type) error
	EnableConsumer(brokerName string, consumerIP string, messageType message.Type) error
	GetBroker(brokerName string, systemType proto_cluster.Type) (broker.Broker, error)
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

func (s *momService) IsConnected(userIP string) bool {
	for _, conn := range s.connections {
		if conn == userIP {
			return true
		}
	}
	return false
}

func (s *momService) IsSubscribed(name string, messageType message.Type, userIP string) bool {
	var system broker.Broker
	if messageType == message.Type_QUEUE {
		system = s.queues[name]
	} else {
		system = s.topics[name]
	}

	for _, cons := range *system.GetConsumers() {
		if cons.IP == userIP {
			return true
		}
	}

	return false
}

func (s *momService) SystemExists(name string, messageType message.Type) bool {
	var ok bool
	if messageType == message.Type_QUEUE {
		_, ok = s.queues[name]
	} else {
		_, ok = s.topics[name]
	}

	return ok
}

func (s *momService) GetQueues() map[string]*broker.Queue {
	return s.queues
}

func (s *momService) GetTopics() map[string]*broker.Topic {
	return s.topics
}

func (s *momService) GetConnections() []string {
	return s.connections
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
	if s.Config.IsLeader() {
		for _, conn := range s.GetConfig().GetPeers() {
			client := proto_cluster.NewClusterServiceClient(conn)
			req := proto_cluster.ConnectionRequest{Ip: address}
			client.AddConnection(context.Background(), &req)
		}
	}

	s.connections = append(s.connections, address)

}

// Delete a connection from the list of active connections
func (s *momService) DeleteConnection(deleteAddress string) {

	if s.Config.IsLeader() {
		for _, conn := range s.GetConfig().GetPeers() {
			client := proto_cluster.NewClusterServiceClient(conn)
			req := proto_cluster.ConnectionRequest{Ip: deleteAddress}
			client.RemoveConnection(context.Background(), &req)
		}
	}

	newConnections := []string{}
	for _, conn := range s.connections {
		if conn == deleteAddress {
			continue
		}
		newConnections = append(newConnections, conn)
	}

	s.connections = newConnections
}

// Create a new queue and add it to the list of active queues
func (s *momService) CreateQueue(name string, clientIP string) error {
	if !s.IsConnected(clientIP) {
		return ErrNotConnected
	}

	if s.SystemExists(name, message.Type_QUEUE) {
		return ErrSystemExists
	}

	if s.Config.IsLeader() {
		for _, conn := range s.GetConfig().GetPeers() {
			client := proto_cluster.NewClusterServiceClient(conn)
			req := proto_cluster.SystemRequest{
				Name:    name,
				Type:    proto_cluster.Type_QUEUE,
				Creator: clientIP,
			}
			client.AddMessagingSystem(context.Background(), &req)
		}
	}

	queue := broker.NewQueue(clientIP)
	s.queues[name] = queue

	return nil
}

// Delete the queue with the given name
func (s *momService) DeleteQueue(name string, clientIp string) error {
	if clientIp != s.queues[name].Creator {
		return ErrBrokerNotFound
	}

	if !s.IsSubscribed(name, message.Type_QUEUE, clientIp) {
		return ErrNotSubscribed
	}

	if s.Config.IsLeader() {
		for _, conn := range s.GetConfig().GetPeers() {
			client := proto_cluster.NewClusterServiceClient(conn)
			req := proto_cluster.SystemRequest{
				Name:    name,
				Type:    proto_cluster.Type_QUEUE,
				Creator: clientIp,
			}
			client.RemoveMessagingSystem(context.Background(), &req)
		}
	}

	delete(s.queues, name)
	return nil
}

// Create a new queue and add it to the list of active queues
func (s *momService) CreateTopic(name string, clientIP string) error {

	if !s.IsConnected(clientIP) {
		return ErrNotConnected
	}

	if s.SystemExists(name, message.Type_TOPIC) {
		return ErrSystemExists
	}

	if s.Config.IsLeader() {
		for _, conn := range s.GetConfig().GetPeers() {
			client := proto_cluster.NewClusterServiceClient(conn)
			req := proto_cluster.SystemRequest{
				Name:    name,
				Type:    proto_cluster.Type_TOPIC,
				Creator: clientIP,
			}
			client.AddMessagingSystem(context.Background(), &req)
		}
	}

	topic := broker.NewTopic(clientIP)
	s.topics[name] = topic

	return nil
}

// Delete the queue with the given name
func (s *momService) DeleteTopic(name string, clientIp string) error {
	if clientIp != s.topics[name].Creator {
		return ErrBrokerNotFound
	}

	if !s.IsSubscribed(name, message.Type_TOPIC, clientIp) {
		return ErrNotSubscribed
	}

	if s.Config.IsLeader() {
		for _, conn := range s.GetConfig().GetPeers() {
			client := proto_cluster.NewClusterServiceClient(conn)
			req := proto_cluster.SystemRequest{
				Name:    name,
				Type:    proto_cluster.Type_TOPIC,
				Creator: clientIp,
			}
			client.RemoveMessagingSystem(context.Background(), &req)
		}
	}

	delete(s.topics, name)
	return nil
}

// Add a subscriber to a queue
func (s *momService) Subscribe(brokerName string, address string, systemType proto_cluster.Type) error {
	if !s.IsConnected(address) {
		return ErrNotConnected
	}

	broker, err := s.GetBroker(brokerName, systemType)
	if err != nil {
		return err
	}

	if s.Config.IsLeader() {
		for _, conn := range s.GetConfig().GetPeers() {
			client := proto_cluster.NewClusterServiceClient(conn)
			req := proto_cluster.SubscriberRequest{
				Name: brokerName,
				Type: systemType,
				Ip:   address,
			}
			client.AddSubscriber(context.Background(), &req)
		}
	}

	broker.AddConsumer(address)

	return nil
}

// Remove a subscriber from a queue
func (s *momService) Unsubscribe(brokerName string, address string, systemType proto_cluster.Type) error {
	if !s.IsConnected(address) {
		return ErrNotConnected
	}

	broker, err := s.GetBroker(brokerName, systemType)
	if err != nil {
		return err
	}

	if s.Config.IsLeader() {
		for _, conn := range s.GetConfig().GetPeers() {
			client := proto_cluster.NewClusterServiceClient(conn)
			req := proto_cluster.SubscriberRequest{
				Name: brokerName,
				Type: systemType,
				Ip:   address,
			}
			client.RemoveSubscriber(context.Background(), &req)
		}
	}

	broker.RemoveConsumer(address)
	return nil
}

// Send a message to a queue or topic
func (s *momService) SendMessage(brokerName string, payload string, messageType message.Type) error {

	systemType := proto_cluster.Type_QUEUE
	if messageType == message.Type_TOPIC {
		systemType = proto_cluster.Type_TOPIC
	}

	broker, err := s.GetBroker(brokerName, systemType)
	if err != nil {
		return err
	}

	if s.Config.IsLeader() {
		for _, conn := range s.GetConfig().GetPeers() {
			client := message.NewMessageServiceClient(conn)
			req := message.MessageRequest{
				Name:    brokerName,
				Type:    messageType,
				Payload: payload,
			}
			client.AddMessage(context.Background(), &req)
		}
	}

	broker.AddMessage(payload)
	return nil
}

func (s *momService) EnableConsumer(brokerName string, consumerIP string, messageType message.Type) error {
	systemType := proto_cluster.Type_QUEUE
	if messageType == message.Type_TOPIC {
		systemType = proto_cluster.Type_TOPIC
	}

	broker, err := s.GetBroker(brokerName, systemType)
	if err != nil {
		return err
	}
	broker.EnableConsumer(consumerIP)

	if s.Config.IsLeader() {

		for _, conn := range s.GetConfig().GetPeers() {
			client := proto_cluster.NewClusterServiceClient(conn)
			req := proto_cluster.EnableConsumerRequest{
				Ip:         consumerIP,
				BrokerName: brokerName,
				Type:       systemType,
			}
			client.EnableConsumer(context.Background(), &req)
		}
	}

	return nil
}

func (s *momService) GetBroker(brokerName string, systemType proto_cluster.Type) (broker.Broker, error) {
	if systemType == proto_cluster.Type_QUEUE {
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
