package handler

import (
	"context"
	"mom/internal/proto/message"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
)

func (q *MessageService) AddMessage(ctx context.Context, messageRequest *message.MessageRequest) (*empty.Empty, error) {
	err := q.momService.SendMessage(messageRequest.Name, messageRequest.Payload, messageRequest.Type)
	if err != nil {
		return &empty.Empty{}, err
	}
	return &empty.Empty{}, nil
}

func (q *MessageService) RemoveMessage(ctx context.Context, messageRequest *message.MessageRequest) (*empty.Empty, error) {
	broker, err := q.momService.GetBroker(messageRequest.Name, messageRequest.Type)
	if err != nil {
		return &empty.Empty{}, err
	}
	broker.PopMessage()
	return &empty.Empty{}, nil
}

func (q *MessageService) ConsumeMessage(stream message.MessageService_ConsumeMessageServer) error {
	errorChan := make(chan error)
	consumerIP, _, _ := net.SplitHostPort(stream.Context().Peer())
	// TODO! brokerName := pending
	// TODO! brokerType := pending
	go func() {
		for {
			request, err := stream.Recv()
			if err != nil {
				errorChan <- err
				return
			}
			err = q.momService.EnableConsumer(consumerIP)
			if err != nil {
				errorChan <- err
				return
			}
		}
	}()

	go func() {
		for {
			response := &message.ConsumeMessageResponse{Payload: payload}
			broker, err := q.momService.GetBroker(brokerName, messageRequest.Type)
			if err != nil {
				errorChan <- err
				return
			}
			payload <- broker.ConsumerMap[consumerIP]
			// sincronizar con replicas
			err := stream.Send(response)
			if err != nil {
				errorChan <- err
				return
			}
		}
	}()
	return <-errorChan
}
