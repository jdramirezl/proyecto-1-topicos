
package consumer



import (
    "github.com/google/uuid"
)


type Consumer struct {
	ID string	
	Available bool
}


func NewConsumer() Consumer {
	id := uuid.New().String()
	return Consumer{ID: id, Available: true}		
}