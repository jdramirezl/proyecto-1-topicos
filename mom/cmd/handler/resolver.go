package handler

import (
	"context"
	"fmt"

	proto_resolver "github.com/jdramirezl/proyecto-1-topicos/mom/internal/proto/resolver"
	"google.golang.org/grpc/peer"

	"github.com/golang/protobuf/ptypes/empty"
)

func (c *ResolverService) NewMaster(ctx context.Context, req *proto_resolver.MasterMessage) (*empty.Empty, error) {
	newIP := req.Ip
	c.Master.SetMaster(newIP)
	
	fmt.Println("New master Request\n New master: " + string(newIP))

	return &empty.Empty{}, nil
}

func (c *ResolverService) GetMaster(ctx context.Context, emp *empty.Empty) (*proto_resolver.MasterMessage, error) {
	currentIp := c.Master.GetMasterIP()
	res := proto_resolver.MasterMessage{
		Ip: currentIp,
	}
	
	p, _ := peer.FromContext(ctx)
	fmt.Println("Get master Request\n From: " + string(p.Addr.String()))
	return &res, nil
}
