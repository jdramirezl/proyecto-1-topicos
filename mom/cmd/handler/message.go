package handler

import (
	"context"
	"fmt"

	proto_cluster "github.com/jdramirezl/proyecto-1-topicos/mom/internal/proto/cluster"
	"github.com/jdramirezl/proyecto-1-topicos/mom/internal/proto/message"
	proto_message "github.com/jdramirezl/proyecto-1-topicos/mom/internal/proto/message"

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
	systemType := proto_cluster.Type_QUEUE
	messageType := messageRequest.Type
	if messageType == proto_message.MessageType_MESSAGETOPIC {
		systemType = proto_cluster.Type_TOPIC
	}

	broker, err := q.momService.GetBroker(messageRequest.Name, systemType)
	if err != nil {
		return &empty.Empty{}, err
	}
	broker.PopMessage()
	return &empty.Empty{}, nil
}

func (q *MessageService) ConsumeMessage(stream message.MessageService_ConsumeMessageServer) error {

	for {
		request, err := stream.Recv()
		if err != nil {
			return err
		}
		fmt.Println("1")
		consumerIp := request.Ip
		err = q.momService.EnableConsumer(request.Name, consumerIp, request.Type)
		if err != nil {
			return err
		}

		fmt.Println("2")
		systemType := proto_cluster.Type_QUEUE
		messageType := request.Type
		if messageType == proto_message.MessageType_MESSAGETOPIC {
			systemType = proto_cluster.Type_TOPIC
		}

		fmt.Println("3")
		broker, err := q.momService.GetBroker(request.Name, systemType)
		if err != nil {
			return err
		}

		fmt.Println("4")
		fmt.Println(consumerIp)
		payload := <-broker.GetConsumerChannel(consumerIp)
		response := &message.ConsumeMessageResponse{Payload: payload}
		// sincronizar con replicas
		fmt.Println("5")
		err = stream.Send(response)
		if err != nil {
			return err
		}
	}
}
