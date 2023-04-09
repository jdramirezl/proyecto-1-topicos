package client

import (
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

func (c *Client) Publish(payload string) {

}

func (c *Client) Subscribe(queue string) chan string {

}
