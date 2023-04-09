package handler

import (
	"context"
	"mom/cmd/handler"
)

// func (c *ClusterService) AddMessagingSystem(ctx context.Context, _ *cluster.ConsumeMessageRequest) (*files.ConsumeMessageResponse, error) {
// }
// func (c *ClusterService) RemoveMessagingSystem(ctx context.Context, _ *proto_queue.ConsumeMessageRequest) (*files.ConsumeMessageResponse, error) {
// }
// func (c *ClusterService) AddSubscriber(ctx context.Context, _ *proto_queue.ConsumeMessageRequest) (*files.ConsumeMessageResponse, error) {
// }
// func (c *ClusterService) RemoveSubscriber(ctx context.Context, _ *proto_queue.ConsumeMessageRequest) (*files.ConsumeMessageResponse, error) {
// }
// func (c *ClusterService) AddConnection(ctx context.Context, _ *proto_queue.ConsumeMessageRequest) (*files.ConsumeMessageResponse, error) {
// }
// func (c *ClusterService) RemoveConnection(ctx context.Context, _ *proto_queue.ConsumeMessageRequest) (*files.ConsumeMessageResponse, error) {
// }

func (c *ClusterService) AddPeer(ctx context.Context, req *cluster.PeerRequest) {
	// Add peer to config
	c.momService.Config.addPeer(req.Ip)

}

func (c *ClusterService) RemovePeer(ctx context.Context, req *cluster.PeerRequest) {
	// Remove peer from config
	c.momService.Config.removePeer(req.Ip)
}

// func (c *ClusterService) NewMaster(ctx context.Context, _ *proto_queue.ConsumeMessageRequest) (*files.ConsumeMessageResponse, error) {
// }
func (c *ClusterService) Heartbeat(ctx context.Context) {
	// Reset TTL
	main.Config.timeout = time.Now().UnixNano() // TODO
}

// func (c *ClusterService) ElectLeader(ctx context.Context, _ *proto_queue.ConsumeMessageRequest) (*files.ConsumeMessageResponse, error) {
// }
