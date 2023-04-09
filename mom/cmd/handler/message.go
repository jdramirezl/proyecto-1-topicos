package handler

import (
	"context"
	"fmt"
	"mom/internal/proto/message"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/peer"
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
	peer, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("failed to extract peer from context")
	}
	addr := peer.Addr.String()
	consumerIP, _, err := net.SplitHostPort(addr)
	if err != nil {
		return fmt.Errorf("failed to split host and port: %v", err)
	}

	for {
		request, err := stream.Recv()
		if err != nil {
			return err
		}
		err = q.momService.EnableConsumer(request.Name, consumerIP, request.Type)
		if err != nil {
			return err
		}
		broker, err := q.momService.GetBroker(request.Name, request.Type)
		if err != nil {
			return err
		}
		payload := <-broker.GetConsumerChannel(consumerIP)
		response := &message.ConsumeMessageResponse{Payload: payload}
		// sincronizar con replicas
		err = stream.Send(response)
		if err != nil {
			return err
		}
	}
}
