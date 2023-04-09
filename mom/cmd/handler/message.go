package handler

import (
	"context"
	"mom/internal/proto/message"

	"github.com/golang/protobuf/ptypes/empty"
)

func (q *MessageService) AddMessage(ctx context.Context, messageRequest *message.MessageRequest) (*empty.Empty, error) {
	q.momService.SendMessage(messageRequest.Name, messageRequest.Payload, messageRequest.Type)
	return &empty.Empty{}, nil
}

func (q *MessageService) RemoveMessage(ctx context.Context, messageRequest *message.MessageRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (q *MessageService) ConsumeMessage(stream message.MessageService_ConsumeMessageServer) error {
	errorChan := make(chan error)
	go func() {
		for {
			request, err := stream.Recv()
			if err != nil {
				errorChan <- err
				return
			}
		}
	}()
	go func() {
		for {
			response := &message.ConsumeMessageResponse{}
			err := stream.Send(response)
			if err != nil {
				errorChan <- err
				return
			}
		}
	}()
	return <-errorChan
}
