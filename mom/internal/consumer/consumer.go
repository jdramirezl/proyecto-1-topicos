
package consumer



import (
    "github.com/google/uuid"
)


type Consumer struct {
	ID string	
	Available bool
	IP string
}


func NewConsumer(address string) Consumer {
	id := uuid.New().String()
	return Consumer{ID: id, Available: true, IP: address}		
}