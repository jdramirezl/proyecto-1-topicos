package mom

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jdramirezl/proyecto-1-topicos/mom/internal/broker"

	"github.com/jdramirezl/proyecto-1-topicos/mom/internal/cluster"
	proto_cluster "github.com/jdramirezl/proyecto-1-topicos/mom/internal/proto/cluster"
	"github.com/jdramirezl/proyecto-1-topicos/mom/internal/proto/message"
)

var (
	ErrBrokerNotFound = errors.New("broker not found")
	ErrNotConnected   = errors.New("user is not connected, action forbidden")
	ErrNotSubscribed  = errors.New("user is not subscribed, action forbidden")
	ErrSystemExists   = errors.New("system with the same name already exists")
)

type MomService interface {
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
	SendMessage(brokerName string, payload string, messageType message.MessageType) error
	EnableConsumer(brokerName string, consumerIP string, messageType message.MessageType) error
	GetBroker(brokerName string, systemType proto_cluster.Type) (broker.Broker, error)
	GetConfig() *cluster.Config
	Reset()
	GetImGaye() func()
	Update()
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

	m.Config.SetFunc(m.Update)

	m.LeaderActions()
	return m
}

func (s *momService) LeaderActions() {
	go func() {
		for {
			if s.Config.IsLeader() {
				for _, queue := range s.queues {
					queue.Consume()
				}

				for _, topic := range s.topics {
					topic.Consume()
				}
				return
			}
		}
	}()
}

func (s *momService) Reset() {
	s.topics = map[string]*broker.Topic{}
	s.queues = map[string]*broker.Queue{}
	s.connections = []string{}
	s.Config.Reset()
}

func (s *momService) Update() {
	go func() {
		conf := s.GetConfig()
		for {

			time.Sleep(time.Second * 60)

			fmt.Println("Started catching up")
			for _, conn := range conf.GetPeers() {

				conf.CatchYouUp(
					conn,
					s.GetConnections(),
					s.GetQueues(),
					s.GetTopics(),
				)
			}

		}
	}()
}

func (s *momService) GetImGaye() func() {
	return func() {

	}
}

func (s *momService) IsConnected(userIP string) bool {
	for _, conn := range s.connections {
		if conn == userIP {
			return true
		}
	}
	return false
}

func (s *momService) IsSubscribed(name string, messageType message.MessageType, userIP string) bool {
	var system broker.Broker
	if messageType == message.MessageType_MESSAGEQUEUE {
		system = s.queues[name]
	} else {
		system = s.topics[name]
	}

	for _, cons := range system.GetConsumers() {
		if cons.IP == userIP {
			return true
		}
	}

	return false
}

func (s *momService) SystemExists(name string, messageType message.MessageType) bool {
	var ok bool
	if messageType == message.MessageType_MESSAGEQUEUE {
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

// Create a new connection and add it to the list of active connections
func (s *momService) CreateConnection(address string) {
	fmt.Println("New connection received: " + address)

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
	fmt.Println("Connection delete received: " + deleteAddress)
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
		fmt.Printf("Client %s is not connected", clientIP)
		return ErrNotConnected
	}

	if s.SystemExists(name, message.MessageType_MESSAGEQUEUE) {
		fmt.Printf("Queue %s already exists", name)
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

	if s.Config.IsLeader() {
		s.queues[name].Consume()
	}

	fmt.Println("Creating queue: "+name, "with creator: "+clientIP)

	return nil
}

// Delete the queue with the given name
func (s *momService) DeleteQueue(name string, clientIp string) error {
	if clientIp != s.queues[name].Creator {
		fmt.Printf("Client %s is not the creator of queue %s", clientIp, name)
		return ErrBrokerNotFound
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
	fmt.Println("Deleting queue: "+name, "with creator: "+clientIp)
	return nil
}

// Create a new queue and add it to the list of active queues
func (s *momService) CreateTopic(name string, clientIP string) error {

	if !s.IsConnected(clientIP) {
		fmt.Printf("Client %s is not connected", clientIP)
		return ErrNotConnected
	}

	if s.SystemExists(name, message.MessageType_MESSAGETOPIC) {
		fmt.Printf("Topic %s already exists", name)
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
	fmt.Println("Creating topic: "+name, "with creator: "+clientIP)

	return nil
}

// Delete the queue with the given name
func (s *momService) DeleteTopic(name string, clientIp string) error {
	if clientIp != s.topics[name].Creator {
		fmt.Printf("Client %s is not the creator of topic %s", clientIp, name)
		return ErrBrokerNotFound
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
	fmt.Println("Deleting topic: "+name, "with creator: "+clientIp)
	return nil
}

// Add a subscriber to a queue
func (s *momService) Subscribe(brokerName string, address string, systemType proto_cluster.Type) error {
	if !s.IsConnected(address) {
		fmt.Printf("Client %s is not connected", address)
		return ErrNotConnected
	}

	broker, err := s.GetBroker(brokerName, systemType)

	if err != nil {
		fmt.Printf("Broker %s not found", brokerName)
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
	fmt.Println("Adding subscriber: "+address, "to broker: "+brokerName)

	return nil
}

// Remove a subscriber from a queue
func (s *momService) Unsubscribe(brokerName string, address string, systemType proto_cluster.Type) error {
	if !s.IsConnected(address) {
		fmt.Printf("Client %s is not connected", address)
		return ErrNotConnected
	}

	broker, err := s.GetBroker(brokerName, systemType)
	if err != nil {
		fmt.Printf("Broker %s not found", brokerName)
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
	fmt.Println("Removing subscriber: "+address, "from broker: "+brokerName)
	return nil
}

// Send a message to a queue or topic
func (s *momService) SendMessage(brokerName string, payload string, messageType message.MessageType) error {

	systemType := proto_cluster.Type_QUEUE
	if messageType == message.MessageType_MESSAGETOPIC {
		systemType = proto_cluster.Type_TOPIC
	}



	broker, err := s.GetBroker(brokerName, systemType)
	if err != nil {
		fmt.Printf("Broker %s not found", brokerName)
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
	fmt.Println("Sending message: "+payload, "to broker: "+brokerName)
	return nil
}

func (s *momService) EnableConsumer(brokerName string, consumerIP string, messageType message.MessageType) error {
	systemType := proto_cluster.Type_QUEUE
	if messageType == message.MessageType_MESSAGETOPIC {
		systemType = proto_cluster.Type_TOPIC
	}
	
	broker, err := s.GetBroker(brokerName, systemType)
	if err != nil {
		fmt.Printf("Broker %s not found", brokerName)
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
	fmt.Println("Enabling consumer: "+consumerIP, "to broker: "+brokerName)
	return nil
}

func (s *momService) GetBroker(brokerName string, systemType proto_cluster.Type) (broker.Broker, error) {
	if systemType == proto_cluster.Type_QUEUE {
		
		for name, queue := range s.queues {
			if name == brokerName {
				fmt.Println("Found queue: "+name)
				return queue, nil
			}
		}
	} else {
		for name, topic := range s.topics {
			if name == brokerName {
				fmt.Println("Found topic: "+name)
				return topic, nil
			}
		}
	}

	
	return nil, ErrBrokerNotFound
}
