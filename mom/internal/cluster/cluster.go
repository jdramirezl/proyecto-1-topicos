package cluster

import (
	"time"
)

type config struct {
	isVoting bool
	uptime   int64
	peerIPs  []string
	leaderIP string
	timeout  int64
	interval time.Duration
	selfIP string
}

func NewConfig() config {
	// Create config
	// Decides new master!
	return config{isVoting: false, uptime: time.Now().UnixNano(), interval: 2*time.Second}
}


func (c *config) addPeer(peerIP string) {
	c.peerIPs = append(c.peerIPs, peerIP)
}

func (c *config) removePeer(peerIP string) {
	var newPeerIPs []string
	for _, val := range c.peerIPs {
		if val == peerIP{
			continue
		}
		c.peerIPs = append(c.peerIPs, peerIP)
	}
	c.peerIPs = newPeerIPs
}


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

func (c *config) startElection() {
	newLeader := c.selfIP
	bestTime := c.uptime

	for _, ip := range c.peerIPs {
		res := 0 // RPC Methodcall
		if res > int(bestTime) {
			newLeader = ip
		}
	}
	c.leaderIP = newLeader	

	if newLeader == c.selfIP {
		// TODO: DECIRLE AL GATEWAY QUYE IM THE CAPTAIN
	}
}

// Comunicacion in-cluster
func join() {
	// Nodo nuevo envia IP a Nodo maestro
	// tambien a Coordinator
	// If voting ? not join!!
}

func heartbeat() {
	// Nodo en control envia heartbeat a los esclavos
	// Si esclavos no tienen heartbeat del control, hacen consenso
}

func caregiver() {
	// Check TTL of slaves, remove ones with TTL from config
	// TODO: Read TTL
}

func catchmeup() {
	// Llamado en JOIN
	// SOlicitar la info
}

func catchyouup() {
	// Llamado desde el main si nos solicitan info
	// Goroutine????
}
