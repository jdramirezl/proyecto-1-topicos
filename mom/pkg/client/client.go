package client

import (
	"context"
	"fmt"
	"net"

	"github.com/jdramirezl/proyecto-1-topicos/mom/internal/proto/cluster"
	"github.com/jdramirezl/proyecto-1-topicos/mom/internal/proto/message"
	"github.com/jdramirezl/proyecto-1-topicos/mom/internal/proto/resolver"
	"github.com/jdramirezl/proyecto-1-topicos/mom/pkg/internal/connection"

	"github.com/golang/protobuf/ptypes/empty"
)

type Client struct {
	messageClient          message.MessageServiceClient
	clusterClient          cluster.ClusterServiceClient
	resolverClient         resolver.ResolverServiceClient
	isConsuming            bool
	brokerName             string
	queueMessageClient     message.MessageService_ConsumeMessageClient
	masterIP               string
	resetDuringConsumption bool
	selfIp                 string
}

func NewClient(host, port, selfIp string) Client {
	grpcConnResolver, err := connection.NewGrpcClient(host, port)
	if err != nil {
		panic(err)
	}
	resolverClient := resolver.NewResolverServiceClient(grpcConnResolver)

	res, err := resolverClient.GetMaster(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}

	ip := res.Ip
	// clusterPort := os.Getenv("CLUSTER_PORT")

	host, port, _ = net.SplitHostPort(ip)
	grpcConn, err := connection.NewGrpcClient(host, port)
	if err != nil {
		panic(err)
	}
	messageClient := message.NewMessageServiceClient(grpcConn)
	clusterClient := cluster.NewClusterServiceClient(grpcConn)

	return Client{
		masterIP:       ip,
		resolverClient: resolverClient,
		messageClient:  messageClient,
		clusterClient:  clusterClient,
		selfIp:         selfIp,
	}
}

func (c *Client) checkIP() {

	res, err := c.resolverClient.GetMaster(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}
	if res.Ip != c.masterIP {

		ip := res.Ip
		// clusterPort := os.Getenv("CLUSTER_PORT")
		fmt.Println("am here")
		fmt.Println(ip)
		host, port, _ := net.SplitHostPort(ip)
		grpcConn, err := connection.NewGrpcClient(host, port)
		c.masterIP = ip
		if err != nil {
			panic(err)
		}
		messageClient := message.NewMessageServiceClient(grpcConn)
		clusterClient := cluster.NewClusterServiceClient(grpcConn)
		c.messageClient = messageClient
		c.clusterClient = clusterClient

		if c.isConsuming {
			c.resetDuringConsumption = true
		}
	}
}

func (c *Client) AddConnection(ip string) error {
	c.checkIP()
	request := &cluster.ConnectionRequest{Ip: ip}
	_, err := c.clusterClient.AddConnection(context.Background(), request)
	return err
}

func (c *Client) RemoveConnection(payload string) error {
	c.checkIP()
	request := &cluster.ConnectionRequest{Ip: payload}
	_, err := c.clusterClient.RemoveConnection(context.Background(), request)
	return err
}

func (c *Client) PublishQueue(queue string, payload string) error {
	c.checkIP()
	request := message.MessageRequest{Name: queue, Payload: payload, Type: message.MessageType_MESSAGEQUEUE}
	_, err := c.messageClient.AddMessage(context.Background(), &request)
	return err
}

func (c *Client) PublishTopic(queue string, payload string) error {
	c.checkIP()
	request := message.MessageRequest{Name: queue, Payload: payload, Type: message.MessageType_MESSAGETOPIC}
	_, err := c.messageClient.AddMessage(context.Background(), &request)
	return err
}

func (c *Client) SubscribeQueue(name string, ip string) error {
	c.checkIP()
	request := cluster.SubscriberRequest{Name: name, Type: cluster.Type_QUEUE, Ip: ip}
	_, err := c.clusterClient.AddSubscriber(context.Background(), &request)
	return err
}

func (c *Client) SubscribeTopic(name string, ip string) error {
	c.checkIP()
	request := cluster.SubscriberRequest{Name: name, Type: cluster.Type_TOPIC, Ip: ip}
	_, err := c.clusterClient.AddSubscriber(context.Background(), &request)
	return err

}

func (c *Client) UnSubscribeQueue(name string, ip string) error {
	c.checkIP()
	request := cluster.SubscriberRequest{Name: name, Type: cluster.Type_QUEUE, Ip: ip}
	_, err := c.clusterClient.RemoveSubscriber(context.Background(), &request)
	return err
}

func (c *Client) UnSubscribeTopic(name string, ip string) error {
	c.checkIP()
	request := cluster.SubscriberRequest{Name: name, Type: cluster.Type_TOPIC, Ip: ip}
	_, err := c.clusterClient.RemoveSubscriber(context.Background(), &request)
	return err

}

func (c *Client) CreateQueue(name string, creator string) error {
	c.checkIP()
	request := cluster.SystemRequest{Name: name, Type: cluster.Type_QUEUE, Creator: creator}
	_, err := c.clusterClient.AddMessagingSystem(context.Background(), &request)
	return err
}

func (c *Client) CreateTopic(name string, creator string) error {
	c.checkIP()
	request := cluster.SystemRequest{Name: name, Type: cluster.Type_TOPIC, Creator: creator}
	_, err := c.clusterClient.AddMessagingSystem(context.Background(), &request)
	return err
}

func (c *Client) DeleteQueue(name string, creator string) error {
	c.checkIP()
	request := cluster.SystemRequest{Name: name, Type: cluster.Type_QUEUE, Creator: creator}
	_, err := c.clusterClient.RemoveMessagingSystem(context.Background(), &request)
	return err
}

func (c *Client) DeleteTopic(name string, creator string) error {
	c.checkIP()
	request := cluster.SystemRequest{Name: name, Type: cluster.Type_TOPIC, Creator: creator}
	_, err := c.clusterClient.RemoveMessagingSystem(context.Background(), &request)
	return err
}

func (c *Client) Connect(name string) error {
	c.checkIP()
	client, err := c.messageClient.ConsumeMessage(context.Background())
	c.brokerName = name
	c.queueMessageClient = client
	c.isConsuming = true
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ReceiveQueueMessage() (string, error) {
	c.checkIP()

	if !c.isConsuming {
		return "", fmt.Errorf("client is not consuming")
	}
	if c.resetDuringConsumption {
		c.Connect(c.brokerName)
		c.resetDuringConsumption = false
	}

	request := message.ConsumeMessageRequest{Name: c.brokerName, Type: message.MessageType_MESSAGEQUEUE, Ip: c.selfIp}
	err := c.queueMessageClient.Send(&request)
	if err != nil {
		return "", nil
	}
	msg, err := c.queueMessageClient.Recv()
	if err != nil {
		return "", nil
	}
	return msg.Payload, nil
}

func (c *Client) ReceiveTopicMessage() (string, error) {
	c.checkIP()

	if !c.isConsuming {
		return "", fmt.Errorf("client is not consuming")
	}
	if c.resetDuringConsumption {
		c.Connect(c.brokerName)
		c.resetDuringConsumption = false
	}

	request := message.ConsumeMessageRequest{Name: c.brokerName, Type: message.MessageType_MESSAGETOPIC, Ip: c.selfIp}
	err := c.queueMessageClient.Send(&request)
	if err != nil {
		return "", nil
	}
	msg, err := c.queueMessageClient.Recv()
	if err != nil {
		return "", nil
	}
	return msg.Payload, nil
}
