package handler

import (
	"context"

	proto_cluster "mom/internal/proto/cluster"
	"mom/internal/proto/message"

	"github.com/golang/protobuf/ptypes/empty"
)

func (c *ClusterService) AddMessagingSystem(ctx context.Context, req *proto_cluster.SystemRequest) (*empty.Empty, error) {
	Type := req.Type
	Name := req.Name
	Creator := req.Creator

	if Type == 0 {
		c.momService.CreateQueue(Name, Creator)
	} else {
		c.momService.CreateTopic(Name, Creator)
	}

	return &empty.Empty{}, nil
}

func (c *ClusterService) RemoveMessagingSystem(ctx context.Context, req *proto_cluster.ConsumeMessageRequest) (*files.ConsumeMessageResponse, error) {

	Type := req.Type
	Name := req.Name
	Creator := req.Creator

	if Type == 0 {
		c.momService.DeleteQueue(Name, Creator)
	} else {
		c.momService.DeleteTopic(Name, Creator)
	}

	return &empty.Empty{}, nil
}

func (c *ClusterService) AddSubscriber(ctx context.Context, req *proto_cluster.SubscriberRequest) (*files.ConsumeMessageResponse, error) {
	Type := req.Type
	Name := req.Name
	Creator := req.Creator

	if Type == 0 {
		Type = message.Type_QUEUE
	} else {
		Type = message.Type_TOPIC
	}

	c.momService.Subscribe(Name, Creator, Type)

	return &empty.Empty{}, nil
}

func (c *ClusterService) RemoveSubscriber(ctx context.Context, req *proto_cluster.SubscriberRequest) (*files.ConsumeMessageResponse, error) {
	Type := req.Type
	Name := req.Name
	Creator := req.Creator

	if Type == 0 {
		Type = message.Type_QUEUE
	} else {
		Type = message.Type_TOPIC
	}

	c.momService.Unsubscribe(Name, Creator, Type)

	return &empty.Empty{}, nil
}

func (c *ClusterService) AddConnection(ctx context.Context, _ *proto_cluster.ConnectionRequest) (*files.ConsumeMessageResponse, error) {
	connectionIp := req.Ip
	c.momService.CreateConnection(connectionIp)
	return &empty.Empty{}, nil
}

func (c *ClusterService) RemoveConnection(ctx context.Context, _ *proto_cluster.ConnectionRequest) (*files.ConsumeMessageResponse, error) {
	connectionIp := req.Ip
	c.momService.DeleteConnection(connectionIp)
	return &empty.Empty{}, nil
}

func (c *ClusterService) AddPeer(ctx context.Context, req *proto_cluster.PeerRequest) (*empty.Empty, error) {
	c.momService.GetConfig().AddPeer(req.Ip)
	return &empty.Empty{}, nil
}

func (c *ClusterService) RemovePeer(ctx context.Context, req *proto_cluster.PeerRequest) (*empty.Empty, error) {
	c.momService.GetConfig().RemovePeer(req.Ip)
	return &empty.Empty{}, nil
}

// func (c *ClusterService) NewMaster(ctx context.Context, _ *proto_cluster.ConsumeMessageRequest) (*files.ConsumeMessageResponse, error) {
// }

func (c *ClusterService) Heartbeat(ctx context.Context, emp *empty.Empty) (*empty.Empty, error) {
	c.momService.GetConfig().RefreshTimeout()
	return &empty.Empty{}, nil
}

func (c *ClusterService) ElectLeader(ctx context.Context, emp *empty.Empty) (*proto_cluster.ElectLeaderResponse, error) {
	uptime := c.momService.GetConfig().GetUptime()
	res := p
	return
}
