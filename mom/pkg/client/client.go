package client

import (
	"context"
	"fmt"
	"mom/internal/proto/cluster"
	"mom/internal/proto/message"
	"mom/pkg/internal/connection"
)

type Client struct {
	messageClient message.MessageServiceClient
	clusterClient cluster.ClusterServiceClient
	ready         chan bool
	isConsuming   bool
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

func (c *Client) Subscribe(queue string) (chan string, error) {
	request := message.ConsumeMessageRequest{Name: queue, Type: message.Type_QUEUE}
	client, err := c.messageClient.ConsumeMessage(context.Background(), &request)
	c.isConsuming = true
	if err != nil {
		return nil, err
	}
	ret := make(chan string)
	go func() {
		for {
			request := message.ConsumeMessageRequest{Name: queue, Type: message.Type_QUEUE}
			msg, err := client.Recv()
		}
	}()
	return ret, nil
}

func (c *Client) Ready() error {
	if !c.isConsuming {
		return fmt.Errorf("client is not consuming")
	}
	c.ready <- true
	return nil
}
