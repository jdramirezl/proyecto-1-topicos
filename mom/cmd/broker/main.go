package main

import (
	"mom/internal/linked_list"
	"mom/internal/queue"
	"sync"

	"github.com/google/uuid"
    "mom/internal/consumer"
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
	queues      map[string] *queue.Queue
    connections map[string] *connection
    nextConnID  int64
    lock        sync.Mutex
}


func main(){
    
}

// Create a new connection and add it to the list of active connections
func (s *momServer) createConnection(address string) {
    conn := &connection{
        id:      n_client, // TODO: babas + UUID len(s.connections) + 1,
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
        conns = append(conns, []string{string(value.id), key})
    }

    return conns
}

// Create a new queue and add it to the list of active queues
func (s *momServer) createQueue(name string) {
    

    queue := queue.NewQueue()

    s.queues[name] = &

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

// Handle requests from clients by calling the appropriate server-side methods
func (s *momServer) HandleRequest() {
    
}
