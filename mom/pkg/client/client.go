package client

import (
	"context"
	"mom/internal/proto/cluster"
	"mom/internal/proto/message"
	"mom/pkg/internal/connection"
)

type Client struct {
	messageClient message.MessageServiceClient
	clusterClient cluster.ClusterServiceClient
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

func (c *Client) SubscribeQueue(queue string) chan string {
	// request := message.
	client, err := c.messageClient.ConsumeMessage(context.Background(), &request)
	go func() {
		for {
			msg, err := client.Recv()
		}
	}()
	return err
}
