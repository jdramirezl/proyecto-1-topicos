package cluster

import (
	"fmt"
	broker "jdramirezl/proyecto-1-topicos/mom/internal/broker"
	proto_cluster "jdramirezl/proyecto-1-topicos/mom/internal/proto/cluster"
	proto_message "jdramirezl/proyecto-1-topicos/mom/internal/proto/message"
	proto_resolver "jdramirezl/proyecto-1-topicos/mom/internal/proto/resolver"
	"os"

	"strings"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Config struct {
	isVoting        bool
	uptime          int64
	peerIPs         []string
	peerConnections []*grpc.ClientConn
	leaderIP        string
	timeout         int64
	interval        time.Duration
	selfIP          string
	resolverConn    *grpc.ClientConn
}

func NewConfig() *Config {
	// Create config
	// Decides new master!

	_selfIP := os.Getenv("SELF_IP")
	_selfPort := os.Getenv("PORT")
	_selfAddress := _selfIP + ":" + _selfPort

	_resolverIP := os.Getenv("RESOLVER_IP")
	_resolverPort := os.Getenv("RESOLVER_PORT")
	_resolverAddress := _resolverIP + ":" + _resolverPort

	_resolverConn, err := grpc.Dial(_resolverAddress, grpc.WithInsecure())
	fmt.Print(err)
	_leaderIP := GetLeader(_selfAddress, _resolverConn)

	conf := Config{
		isVoting:        false,
		uptime:          time.Now().UnixNano(),
		peerIPs:         []string{},
		peerConnections: []*grpc.ClientConn{},
		leaderIP:        _leaderIP,
		timeout:         time.Now().UnixNano(), // Revisar
		interval:        2 * time.Second,
		selfIP:          _selfAddress,
		resolverConn:    _resolverConn, // TODO
	}
	fmt.Println(conf)

	if !conf.IsLeader() {
		conf.AddPeer(_leaderIP)
		conf.watchLeader()
	} else {
		conf.watchPeers()
	}

	return &conf
}

// ------ GETS -----------
func (c *Config) GetPeers() []*grpc.ClientConn {
	return c.peerConnections
}

// ---------- Check Leader ----------
func (c *Config) IsLeader() bool {
	return c.selfIP == c.leaderIP
}

func GetLeader(selfIP string, resolverConn *grpc.ClientConn) string {
	client := proto_resolver.NewResolverServiceClient(resolverConn)

	res, err := client.GetMaster(context.Background(), &emptypb.Empty{})
	fmt.Println(err)
	fmt.Println("8===============D")
	_leaderIP := res.Ip

	if _leaderIP == "" {
		_leaderIP = selfIP
		req := proto_resolver.MasterMessage{Ip: selfIP}
		client.NewMaster(context.Background(), &req)
	}

	return _leaderIP
}

// ---------- Add a peer ----------
// Receiver
func (c *Config) AddPeer(peerIP string) {
	c.peerIPs = append(c.peerIPs, peerIP)
	peerConn, _ := grpc.Dial(peerIP, grpc.WithInsecure())
	c.peerConnections = append(c.peerConnections, peerConn) // TODO: Cambiar el dial

	if c.IsLeader() {
		// Add peer to others
		c.join(peerIP)
	}

}

// Sender
func (c *Config) join(ip string) {
	for _, conn := range c.peerConnections {
		client := proto_cluster.NewClusterServiceClient(conn)

		req := proto_cluster.PeerRequest{Ip: ip}

		client.AddPeer(context.Background(), &req)
	}
}

// ---------- Remove a peer ----------

// Receiver
func (c *Config) RemovePeer(peerIP string) {
	var newPeerIPs []string
	for _, val := range c.peerIPs {
		if val == peerIP {
			continue
		}
		newPeerIPs = append(newPeerIPs, peerIP)
	}
	c.peerIPs = newPeerIPs

	var newPeerConnections []*grpc.ClientConn
	for _, peerConn := range c.peerConnections {
		IP := getHost(peerConn.Target())
		if IP == peerIP {
			continue
		}
		newPeerConnections = append(newPeerConnections, peerConn)
	}
	c.peerConnections = newPeerConnections
}

// Sender
func (c *Config) orderRemove(peerIP string) {

	c.RemovePeer(peerIP)

	for _, conn := range c.peerConnections {
		client := proto_cluster.NewClusterServiceClient(conn)

		req := proto_cluster.PeerRequest{Ip: peerIP}

		client.RemovePeer(context.Background(), &req)
	}
}

// ---------- Watchers ----------
// Followers
func (c *Config) watchLeader() {
	go func() {
		for {
			if c.isVoting {
				break
			}

			if time.Now().UnixNano() > c.timeout {
				c.startElection()
				time.Sleep(c.interval)
			}
			time.Sleep(c.interval)
		}
	}()
}

// Master
func (c *Config) watchPeers() {
	go func() {

		for {
			if c.isVoting {
				break
			}

			c.heartbeat()
			time.Sleep(c.interval)
		}
	}()
}

// Receiver
func (c *Config) RefreshTimeout() {
	c.timeout = time.Now().UnixNano() // TODO
}

// Sender
func (c *Config) heartbeat() { // TODO: cambiar el tiempo!
	for _, conn := range c.peerConnections {
		client := proto_cluster.NewClusterServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		_, err := client.Heartbeat(ctx, &emptypb.Empty{})

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
func (c *Config) startElection() {
	newLeader := c.selfIP
	bestTime := c.uptime
	c.isVoting = true

	for _, conn := range c.peerConnections {
		client := proto_cluster.NewClusterServiceClient(conn)
		res, _ := client.ElectLeader(context.Background(), &emptypb.Empty{})

		uptime := res.Uptime

		if uptime > bestTime {
			newLeader = getHost(conn.Target())
		}
	}

	c.leaderIP = newLeader

	c.isVoting = false

	if c.IsLeader() {
		client := proto_resolver.NewResolverServiceClient(c.resolverConn)
		req := proto_resolver.MasterMessage{Ip: c.selfIP}
		client.NewMaster(context.Background(), &req)

		c.watchPeers()

	} else {
		c.watchLeader()
	}

}

func (c *Config) GetUptime() int64 {
	return c.uptime
}

func (c *Config) CatchYouUp(
	connections []string,
	queues map[string]*broker.Queue,
	topics map[string]*broker.Topic,
) {
	follower_conn := c.peerConnections[len(c.peerConnections)-1]
	cluster_client := proto_cluster.NewClusterServiceClient(follower_conn)
	message_client := proto_message.NewMessageServiceClient(follower_conn)

	// Add Peers
	for _, ip := range c.peerIPs {
		req := proto_cluster.PeerRequest{
			Ip: ip,
		}
		cluster_client.AddPeer(context.Background(), &req)
	}

	// Add connections
	for _, ip := range connections {
		req := proto_cluster.ConnectionRequest{
			Ip: ip,
		}
		cluster_client.AddConnection(context.Background(), &req)
	}

	// Add queues
	for queue_name, queue := range queues {
		c.messageSystemCatchUp(
			queue_name,
			proto_message.MessageType_MESSAGEQUEUE,
			message_client,
			cluster_client,
			queue,
		)
	}

	// Add topic
	for topic_name, topic := range topics {
		c.messageSystemCatchUp(
			topic_name,
			proto_message.MessageType_MESSAGEQUEUE,
			message_client,
			cluster_client,
			topic,
		)
	}
}

func (c *Config) messageSystemCatchUp(
	name string,
	system_type proto_message.MessageType,
	m_client proto_message.MessageServiceClient,
	c_client proto_cluster.ClusterServiceClient,
	system broker.Broker,
) {
	system_req_type := proto_cluster.Type_QUEUE
	message_req_type := proto_message.MessageType_MESSAGEQUEUE
	if system_type == proto_message.MessageType_MESSAGETOPIC {
		system_req_type = proto_cluster.Type_TOPIC
		message_req_type = proto_message.MessageType_MESSAGETOPIC
	}

	req := proto_cluster.SystemRequest{
		Name:    name,
		Type:    system_req_type,
		Creator: system.GetCreator(),
	}
	c_client.AddMessagingSystem(context.Background(), &req)

	// Add messages
	linkedlist := system.GetMessages()
	n := linkedlist.Head
	for n != nil {
		req := proto_message.MessageRequest{
			Name:    name,
			Type:    message_req_type,
			Payload: n.Val,
		}
		m_client.AddMessage(context.Background(), &req)

		n = n.Next
	}

	// Add consumers
	consumers := system.GetConsumers()
	for _, consumer := range *consumers {
		req := proto_cluster.SubscriberRequest{
			Name: name,
			Type: system_req_type,
			Ip:   consumer.IP,
		}
		c_client.AddSubscriber(context.Background(), &req)
	}
}
