package main;

import (

)

type connection struct {
    id      int64
    address string
    // add other relevant metadata here
}

type queue struct {
	name    string
    channel chan *mom.Message
}


type momServer struct {
	queues      map[string]*momQueue
    connections map[int64]*momConnection
    nextConnID  int64
    lock        sync.Mutex
}

// Create a new connection and add it to the list of active connections
func (s *momServer) createConnection(address string) int64 {
    conn := &connection{
        id:      1, // TODO: babas + UUID len(s.connections) + 1,
        address: address,
    }
    s.connections = append(s.connections, conn)
    return conn.id
}

// Delete a connection from the list of active connections
func (s *momServer) deleteConnection(connID int) {
    for _, conn := range s.connections {
        if conn.id == connID {
            s.connections = append(s.connections[:i], s.connections[i+1:]...)
            break
        }
    }
}

// Get a list of all active connections
func (s *momServer) getConnections() []*connection {
    return s.connections
}

// Create a new queue and add it to the list of active queues
func (s *momServer) createQueue(name string) {
    s.queues[name] = &momQueue{
        name:    name,
        channel: make(chan *mom.Message),
    }
}

// Delete the queue with the given name
func (s *momServer) deleteQueue(name string) {
    delete(s.queues, name)
}

// Get a list of all active queues
func (s *momServer) getQueues() []*queue {
    return s.queues
}

// Add a subscriber to a queue
func (s *momServer) subscribe(queueName string, connID int) {
    for _, queue := range s.queues {
        if queue.name == queueName {
            // TODO: llamar metodo
        }
    }
}

// Remove a subscriber from a queue
func (s *momServer) unsubscribe(queueName string, connID int) {
    for _, queue := range s.queues {
        if queue.name == queueName {
            // TODO: llamar metodo
        }
    }
}

// Send a message to a queue
func (s *momServer) receiveMessage(queueName, message string) {
    for _, queue := range s.queues {
        if queue.name == queueName {
            // TODO: llamar metodo
        }
    }
}

// Handle requests from clients by calling the appropriate server-side methods
func (s *momServer) HandleRequest() {
    
}
