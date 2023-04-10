package client

import (
	"context"
	"fmt"
	"mom/internal/proto/cluster"
	"mom/internal/proto/message"
	"mom/pkg/internal/connection"
)

type Client struct {
	messageClient      message.MessageServiceClient
	clusterClient      cluster.ClusterServiceClient
	isConsuming        bool
	brokerName         string
	queueMessageClient message.MessageService_ConsumeMessageClient
}

func NewClient(host, port string) Client {
	grpcConn, err := connection.NewGrpcClient(host, port)
	if err != nil {
		panic(err)
	}
	messageClient := message.NewMessageServiceClient(grpcConn)
	clusterClient := cluster.NewClusterServiceClient(grpcConn)

	return Client{messageClient: messageClient, clusterClient: clusterClient}
}

func (c *Client) AddConnection(payload string) error {
	request := cluster.ConnectionRequest{Ip: payload}
	_, err := c.clusterClient.addConnection(context.Background(), request)
	return err
}

func (c *Client) RemoveConnection(payload string) error {
	request := cluster.ConnectionRequest{Ip: payload}
	_, err := c.clusterClient.removeConnection(context.Background(), request)
	return err
}

func (c *Client) PublishQueue(payload string, queue string) error {
	request := message.MessageRequest{Name: queue, Payload: payload, Type: message.Type_QUEUE}
	_, err := c.messageClient.AddMessage(context.Background(), &request)
	return err
}

func (c *Client) PublishTopic(payload string, queue string) error {
	request := message.MessageRequest{Name: queue, Payload: payload, Type: message.Type_TOPIC}
	_, err := c.messageClient.AddMessage(context.Background(), &request)
	return err
}

func (c *Client) SubscribeQueue(queueName string) error {
	client, err := c.messageClient.ConsumeMessage(context.Background())
	c.brokerName = queueName
	c.queueMessageClient = client
	c.isConsuming = true
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ReceiveQueueMessage() (string, error) {
	if !c.isConsuming {
		return "", fmt.Errorf("client is not consuming")
	}

	request := message.ConsumeMessageRequest{Name: c.brokerName, Type: message.Type_QUEUE}
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
