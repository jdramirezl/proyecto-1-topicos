package cluster

import (
	"fmt"
	"os"

	broker "github.com/jdramirezl/proyecto-1-topicos/mom/internal/broker"
	proto_cluster "github.com/jdramirezl/proyecto-1-topicos/mom/internal/proto/cluster"
	proto_message "github.com/jdramirezl/proyecto-1-topicos/mom/internal/proto/message"
	proto_resolver "github.com/jdramirezl/proyecto-1-topicos/mom/internal/proto/resolver"

	"strings"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
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
	ImGaye          func()
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

	_resolverConn, _ := grpc.Dial(_resolverAddress, grpc.WithInsecure())

	_leaderIP := GetLeader(_selfAddress, _resolverConn)

	conf := Config{
		isVoting:        false,
		uptime:          time.Now().UnixNano(),
		peerIPs:         []string{},
		peerConnections: []*grpc.ClientConn{},
		leaderIP:        _leaderIP,
		timeout:         time.Now().Add(30 * time.Second).Unix(), // Revisar
		interval:        10 * time.Second,
		selfIP:          _selfAddress,
		resolverConn:    _resolverConn, // TODO
	}

	if !conf.IsLeader() {
		conf.AddPeer(_leaderIP)
		conf.join(_selfAddress)
		conf.watchLeader()
	} else {

		conf.watchPeers()
	}

	return &conf
}

func (c *Config) SetFunc(f func()) {
	c.ImGaye = f
}

// ------ GETS -----------
func (c *Config) GetPeers() []*grpc.ClientConn {
	return c.peerConnections
}
func (c *Config) Reset() {
	c.peerIPs = []string{}
	c.peerConnections = []*grpc.ClientConn{}
}

// ---------- Check Leader ----------
func (c *Config) IsLeader() bool {
	return c.selfIP == c.leaderIP
}

func GetLeader(selfIP string, resolverConn *grpc.ClientConn) string {
	client := proto_resolver.NewResolverServiceClient(resolverConn)

	res, _ := client.GetMaster(context.Background(), &emptypb.Empty{})
	_leaderIP := res.Ip

	if _leaderIP == "" {
		fmt.Println("No leader currently, im: " + selfIP)

		_leaderIP = selfIP
		req := proto_resolver.MasterMessage{Ip: selfIP}
		client.NewMaster(context.Background(), &req)
	} else {
		fmt.Println("There is a leader: " + _leaderIP + " and my IP is: " + selfIP)
	}

	return _leaderIP
}

func printConnections(conns []*grpc.ClientConn) {
	fmt.Println("Printing connections------")
	fmt.Print("[")
	for _, conn := range conns {
		fmt.Print(getHost(conn.Target()) + ", ")
	}
	fmt.Print("]")
	fmt.Println()
}

// ---------- Add a peer ----------
// Receiver
func (c *Config) AddPeer(peerIP string) {
	if c.IsLeader() {
		c.join(peerIP)
	}

	fmt.Println("New peer: " + peerIP)
	c.peerIPs = append(c.peerIPs, peerIP)
	peerConn, _ := grpc.Dial(peerIP, grpc.WithInsecure())
	c.peerConnections = append(c.peerConnections, peerConn)
}

// Sender
func (c *Config) join(ip string) {
	fmt.Println("Sending new peer " + ip + " to current peers")
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
	fmt.Println("Removing peer: " + peerIP)
}

// Sender
func (c *Config) orderRemove(peerIP string) {
	fmt.Println("Sending remove peer " + peerIP + " to current peers")
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
				fmt.Println("Im voting so stop!: " + c.selfIP)
				break
			}

			if time.Now().Unix() > c.timeout {
				c.startElection()
				break
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
				fmt.Println("Im leader and voting so stop!: " + c.selfIP)
				break
			}

			c.heartbeat()
			time.Sleep(c.interval)
		}
	}()
}

// Receiver
func (c *Config) RefreshTimeout() {
	//fmt.Println("Refreshing timeout...")
	c.timeout = time.Now().Add(30 * time.Second).Unix() // TODO
	return
}

func (c *Config) sendBeat(conn *grpc.ClientConn) {
	client := proto_cluster.NewClusterServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := client.Heartbeat(ctx, &emptypb.Empty{})

	if err != nil {

		fmt.Println("Error found in beat!")
		c.orderRemove(getHost(conn.Target()))

		return
	}

}

// Sender
func (c *Config) heartbeat() { // TODO: cambiar el tiempo!
	for _, conn := range c.peerConnections {
		fmt.Println("Sending heartbeat to: " + getHost(conn.Target()))
		c.sendBeat(conn)
	}
}

func getHost(target string) string {
	parts := strings.Split(target, ":")
	return parts[0]
}

// ---------- Election ----------
func (c *Config) startElection() {
	fmt.Println("Starting Election")
	printConnections(c.peerConnections)

	// Delete
	c.RemovePeer(c.leaderIP)

	newLeader := c.selfIP
	bestTime := c.uptime
	c.isVoting = true

	for _, conn := range c.peerConnections {

		if getHost(conn.Target()) == getHost(c.leaderIP) {
			continue
		}

		client := proto_cluster.NewClusterServiceClient(conn)
		res, _ := client.ElectLeader(context.Background(), &emptypb.Empty{})

		uptime := res.Uptime

		if uptime > bestTime {
			newLeader = getHost(conn.Target())
		}
	}

	c.leaderIP = newLeader

	c.isVoting = false

	c.RefreshTimeout()

	if c.IsLeader() {
		fmt.Println("Im the new leader: " + c.selfIP + ", long live the king!")
		client := proto_resolver.NewResolverServiceClient(c.resolverConn)
		req := proto_resolver.MasterMessage{Ip: c.selfIP}
		client.NewMaster(context.Background(), &req)

		c.ImGaye()

		c.watchPeers()

	} else {
		fmt.Println("Im was not elected leader, im a follower " + c.selfIP)
		c.watchLeader()
	}

}

func (c *Config) GetUptime() int64 {
	return c.uptime
}

func (c *Config) CatchYouUp(
	follower_conn *grpc.ClientConn,
	connections []string,
	queues map[string]*broker.Queue,
	topics map[string]*broker.Topic,
) {
	// fmt.Println("Starting to catchUp! Catching up: " + getHost(follower_conn.Target()))
	cluster_client := proto_cluster.NewClusterServiceClient(follower_conn)
	message_client := proto_message.NewMessageServiceClient(follower_conn)

	cluster_client.Reset(context.Background(), &emptypb.Empty{})

	req := proto_cluster.PeerRequest{
		Ip: c.selfIP,
	}
	cluster_client.AddPeer(context.Background(), &req)

	// Add Peers
	for _, ip := range c.peerIPs {
		// fmt.Println("Sending ip in catch up: " + ip)
		if getHost(ip) == getHost(follower_conn.Target()) {
			continue
		}
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

	fmt.Println("Finished catching up: " + getHost(follower_conn.Target()))
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
	for _, consumer := range consumers {
		req := proto_cluster.SubscriberRequest{
			Name: name,
			Type: system_req_type,
			Ip:   consumer.IP,
		}
		c_client.AddSubscriber(context.Background(), &req)
	}

}
