package handler

import (
	"context"
	"fmt"

	proto_cluster "github.com/jdramirezl/proyecto-1-topicos/mom/internal/proto/cluster"
	proto_message "github.com/jdramirezl/proyecto-1-topicos/mom/internal/proto/message"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/golang/protobuf/ptypes/empty"
)

func (c *ClusterService) Reset(ctx context.Context, emp *empty.Empty) (*empty.Empty, error) {
	c.momService.Reset()
	return &emptypb.Empty{}, nil
}

func (c *ClusterService) Update() {
	c.momService.Update()
}

func (c *ClusterService) AddMessagingSystem(ctx context.Context, req *proto_cluster.SystemRequest) (*empty.Empty, error) {
	Type := req.Type
	Name := req.Name
	Creator := req.Creator

	if Type == proto_cluster.Type_QUEUE {
		c.momService.CreateQueue(Name, Creator)
	} else {
		c.momService.CreateTopic(Name, Creator)
	}

	return &empty.Empty{}, nil
}

func (c *ClusterService) RemoveMessagingSystem(ctx context.Context, req *proto_cluster.SystemRequest) (*empty.Empty, error) {

	Type := req.Type
	Name := req.Name
	Creator := req.Creator

	if Type == proto_cluster.Type_QUEUE {
		fmt.Println("DELETING QUEUE: " + Name)
		c.momService.DeleteQueue(Name, Creator)
	} else {
		c.momService.DeleteTopic(Name, Creator)
	}

	return &empty.Empty{}, nil
}

func (c *ClusterService) AddSubscriber(ctx context.Context, req *proto_cluster.SubscriberRequest) (*empty.Empty, error) {
	Type := req.Type
	Name := req.Name
	Creator := req.Ip

	var final proto_cluster.Type
	if Type == 0 {
		final = proto_cluster.Type_QUEUE
	} else {
		final = proto_cluster.Type_TOPIC
	}

	c.momService.Subscribe(Name, Creator, final)

	return &empty.Empty{}, nil
}

func (c *ClusterService) RemoveSubscriber(ctx context.Context, req *proto_cluster.SubscriberRequest) (*empty.Empty, error) {
	Type := req.Type
	Name := req.Name
	Creator := req.Ip

	var final proto_cluster.Type
	if Type == 0 {
		final = proto_cluster.Type_QUEUE
	} else {
		final = proto_cluster.Type_TOPIC
	}

	c.momService.Unsubscribe(Name, Creator, final)

	return &empty.Empty{}, nil
}

func (c *ClusterService) AddConnection(ctx context.Context, req *proto_cluster.ConnectionRequest) (*empty.Empty, error) {
	connectionIp := req.Ip
	c.momService.CreateConnection(connectionIp)
	return &empty.Empty{}, nil
}

func (c *ClusterService) RemoveConnection(ctx context.Context, req *proto_cluster.ConnectionRequest) (*empty.Empty, error) {
	connectionIp := req.Ip
	c.momService.DeleteConnection(connectionIp)
	return &empty.Empty{}, nil
}

func (c *ClusterService) AddPeer(ctx context.Context, req *proto_cluster.PeerRequest) (*empty.Empty, error) {
	conf := c.momService.GetConfig()

	if conf.IsLeader() {
		c.Update()
	}

	// mom := c.momService
	// fmt.Println("why are you gae")
	// fmt.Println(req.Ip)
	conf.AddPeer(req.Ip)

	return &empty.Empty{}, nil
}

func (c *ClusterService) RemovePeer(ctx context.Context, req *proto_cluster.PeerRequest) (*empty.Empty, error) {
	c.momService.GetConfig().RemovePeer(req.Ip)
	return &empty.Empty{}, nil
}

func (c *ClusterService) Heartbeat(ctx context.Context, emp *empty.Empty) (*empty.Empty, error) {
	// fmt.Println("Got refresh request!")
	c.momService.GetConfig().RefreshTimeout()
	return &empty.Empty{}, nil
}

func (c *ClusterService) ElectLeader(ctx context.Context, emp *empty.Empty) (*proto_cluster.ElectLeaderResponse, error) {
	uptime := c.momService.GetConfig().GetUptime()
	res := &proto_cluster.ElectLeaderResponse{
		Uptime: uptime,
	}
	return res, nil
}

func (c *ClusterService) EnableConsumer(ctx context.Context, req *proto_cluster.EnableConsumerRequest) (*empty.Empty, error) {
	messageType := req.Type

	systemType := proto_message.MessageType_MESSAGEQUEUE
	if messageType == proto_cluster.Type_TOPIC {
		systemType = proto_message.MessageType_MESSAGETOPIC
	}

	err := c.momService.EnableConsumer(req.Ip, req.BrokerName, systemType)
	if err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}
