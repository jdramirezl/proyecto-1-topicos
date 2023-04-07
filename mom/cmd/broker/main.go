package main

import (
	"context"
	"fmt"
	"log"
	"mom/internal/proto/cluster"
	"mom/internal/queue"
	"net"
	"os"

	"github.com/envoyproxy/go-control-plane/envoy/api/v2/cluster"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/text/message"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	n_client int64 = 0
)

type connection struct {
	id      int64
	address string
	// add other relevant metadata here
}

type momServer struct {
	queues      map[string]*queue.Queue
	connections map[string]*connection
}

var (
	mom *momServer
)

// func main() {
// 	mom := &momServer{}
// 	mom.init()
// }

func (s *momServer) init() {
	s.queues = map[string]*queue.Queue{}
	s.connections = map[string]*connection{}
}

// Create a new connection and add it to the list of active connections
func (s *momServer) createConnection(address string) {
	conn := &connection{
		id:      n_client,
		address: address,
	}

	n_client += 1

	s.connections[address] = conn
}

// Delete a connection from the list of active connections
func (s *momServer) deleteConnection(address string) {
	delete(s.connections, address)
}

// Get a list of all active connections
func (s *momServer) getConnections() [][]string {
	conns := make([][]string, len(s.connections))

	for key, value := range s.connections {
		conns = append(conns, []string{fmt.Sprint(value.id), key})
	}

	return conns
}

// Create a new queue and add it to the list of active queues
func (s *momServer) createQueue(name string) {
	queue := queue.NewQueue()
	s.queues[name] = queue
}

// Delete the queue with the given name
func (s *momServer) deleteQueue(name string) {
	delete(s.queues, name)
}

// Add a subscriber to a queue
func (s *momServer) subscribe(queueName string, address string) {
	for name, queue := range s.queues {
		if name == queueName {
			queue.AddConsumer(address)
			break
		}
	}
}

// Remove a subscriber from a queue
func (s *momServer) unsubscribe(queueName string, address string) {
	for name, queue := range s.queues {
		if name == queueName {
			queue.RemoveConsumer(address)
			break
		}
	}
}

// Send a message to a queue
func (s *momServer) sendMessage(queueName string, message string) {
	for name, queue := range s.queues {
		if name == queueName {
			queue.AddMessage(message)
			break
		}
	}
}

type MessageServer struct {
	message.UnimplementedMessageServer
}

func (m *MessageServer) AddMessage(ctx context.Context, _ *message.MessageRequest) (*google_protobuf.Empty, error) {
}

// func (f *MessageServer) RemoveMessage(ctx context.Context, _ *files.FileListRequest) (*files.FileListResponse, error) {
// }

// func (f *MessageServer) ConsumeMessage(ctx context.Context, _ *files.FileListRequest) (*files.FileListResponse, error) {
// }

type ClusterService struct {
	cluster.UnimplementedClusterServiceServer
}

func main() {
	port := os.Getenv("PORT")
	ip := "0.0.0.0"

	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", ip, port))
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	messageServiceServer := &MessageServer{}
	clusterServiceServer := &ClusterService{}
	message.RegisterQueueServiceServer(server, messageServiceServer)
	cluster.RegisterClusterServiceServer(server, clusterServiceServer)

	reflection.Register(server)
	log.Printf("Listening on port: %v\n", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
