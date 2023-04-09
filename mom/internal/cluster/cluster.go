package cluster

import (
	"mom/internal/broker"
	cluster "mom/internal/proto/cluster"
	message "mom/internal/proto/message"
	"strings"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type config struct {
	isVoting        bool
	uptime          int64
	peerIPs         []string
	peerConnections []*grpc.ClientConn
	leaderIP        string
	timeout         int64
	interval        time.Duration
	selfIP          string
	solverConn      *grpc.ClientConn
}

func NewConfig(solver_conn *grpc.ClientConn) config {
	// Create config
	// Decides new master!
	return config{
		isVoting:        false,
		uptime:          time.Now().UnixNano(),
		peerIPs:         []string{},
		peerConnections: []*grpc.ClientConn{},
		leaderIP:        "",                    // TODO
		timeout:         time.Now().UnixNano(), // Revisar
		interval:        2 * time.Second,
		selfIP:          "",  // TODO
		solverConn:      nil, // TODO
	}
}

// ---------- Check Leader ----------
func (c *config) isLeader() bool {
	return c.selfIP == c.leaderIP
}

// ---------- Add a peer ----------
// Receiver
func (c *config) addPeer(peerIP string) {
	c.peerIPs = append(c.peerIPs, peerIP)
	peerConn := grpc.Dial(peerIP)
	c.peerConnections = append(c.peerConnections, peerConn) // TODO: Cambiar el dial

	if c.isLeader() {
		// Add peer to others
		c.join(peerIP)

	}

}

// Sender
func (c *config) join(ip string) {
	for _, conn := range c.peerConnections {
		client := cluster.NewClusterServiceClient(conn)

		req := cluster.PeerRequest{Ip: ip}

		client.AddPeer(context.Background(), &req)
	}
}

// ---------- Remove a peer ----------

// Receiver
func (c *config) removePeer(peerIP string) {
	var newPeerIPs []string
	for _, val := range c.peerIPs {
		if val == peerIP {
			continue
		}
		c.peerIPs = append(c.peerIPs, peerIP)
	}
	c.peerIPs = newPeerIPs

	var newPeerConnections []*grpc.ClientConn
	for _, val := range c.peerConnections {
		if val == peerIP {
			continue
		}
		c.peerConnections = append(c.peerConnections, peerIP)
	}
	c.peerConnections = newPeerConnections
}

// Sender
func (c *config) orderRemove(peerIP string) {
	for _, conn := range c.peerConnections {
		client := cluster.NewClusterServiceClient(conn)

		req := cluster.PeerRequest{Ip: peerIP}

		client.RemovePeer(context.Background(), &req)
	}
}

// ---------- Watchers ----------
// Followers
func (c *config) watchLeader() {
	go func() {
		for {
			if time.Now().UnixNano() > c.timeout {
				c.startElection()
				time.Sleep(c.interval)
			}
			time.Sleep(c.interval)
		}
	}()
}

// Master
func (c *config) watchPeers() {
	go func() {
		for {
			c.heartbeat()
			time.Sleep(c.interval)
		}
	}()
}

// Receiver
func (c *config) refreshTimeout() {
	c.timeout = time.Now().UnixNano() // TODO
}

// Sender
func (c *config) heartbeat() { // TODO: cambiar el tiempo!
	for _, conn := range c.peerConnections {
		client := cluster.NewClusterServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		_, err := client.Heartbeat(ctx)

		// Check if slave is still alive
		if err != nil {
			if status.Code(err) == codes.DeadlineExceeded {
				// slave did not respond to heartbeat, remove it from the list
				c.orderRemove(getHost(conn.Target()))
			}
		}
	}
}

func getHost(target string) string {
	parts := strings.Split(target, ":")
	return parts[0]
}

// ---------- Election ----------
func (c *config) startElection() {
	newLeader := c.selfIP
	bestTime := c.uptime
	c.isVoting = true

	for _, conn := range c.peerConnections {
		client := cluster.NewClusterServiceClient(conn)
		res := client.ElectLeader(context.Background())

		if res > int(bestTime) {
			newLeader = getHost(conn.Target())
		}
	}

	c.leaderIP = newLeader

	if c.isLeader() {
		client := cluster.NewClusterServiceClient(c.solverConn)
		req := cluster.MasterRequest{Ip: c.selfIP}
		client.NewMaster(context.Background(), &req)
	}

	c.isVoting = false
}

func (c *config) catchYouUp(follower_conn *grpc.ClientConn, Mom *mom.momService) {
	cluster_client := cluster.NewClusterServiceClient(follower_conn)
	message_client := message.NewQueueServiceClient(follower_conn)

	// Add Peers
	for ip, _ := range c.peerIPs {
		req := cluster.ConnectionRequest{
			Ip: ip,
		}
		cluster_client.AddPeer(context.Background(), &req)
	}

	// Add connections
	for _, ip := range Mom.Connections {
		req := cluster.ConnectionRequest{
			Ip: ip,
		}
		cluster_client.AddConnection(context.Background(), &req)
	}

	// Add queues
	for queue_name, queue := range Mom.Queues {
		c.messageSystemCatchUp(
			queue_name,
			message_client,
			cluster_client,
			queue)
	}

	// Add topic
	for topic_name, topic := range Mom.Topics {
		c.messageSystemCatchUp(
			topic_name,
			message_client,
			cluster_client,
			topic)
	}
}

func (c *config) messageSystemCatchUp(
	name string,
	_type int,
	m_client *grpc.ClientConn,
	c_client *grpc.ClientConn,
	system *broker.Broker) {
	req := cluster.SystemRequest{
		Name:    name,
		Type:    _type,
		Creator: system.Creator,
	}
	c_client.AddMessagingSystem(context.Background(), &req)

	// Add messages
	n := system.Messages.Head
	for n != nil {
		req := message.MessageRequest{
			Name:    name,
			Type:    _type,
			Payload: n.Val,
		}
		m_client.AddMessage(context.Background(), &req)

		n = n.Next
	}

	// Add consumers
	for _, consumer := range system.Consumers {
		req := cluster.SubscriberRequest{
			Name: name,
			Type: _type,
			Ip:   consumer.IP,
		}
		c_client.AddSubscriber(context.Background(), &req)
	}
}
