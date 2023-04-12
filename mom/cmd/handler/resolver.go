package handler

import (
	"context"

	proto_resolver "github.com/jdramirezl/proyecto-1-topicos/mom/internal/proto/resolver"

	"github.com/golang/protobuf/ptypes/empty"
)

func (c *ResolverService) NewMaster(ctx context.Context, req *proto_resolver.MasterMessage) (*empty.Empty, error) {
	newIP := req.Ip
	c.Master.SetMaster(newIP)
	// fmt.Println("NEW MASTER: " + string(newIP))
	return &empty.Empty{}, nil
}

func (c *ResolverService) GetMaster(ctx context.Context, emp *empty.Empty) (*proto_resolver.MasterMessage, error) {
	currentIp := c.Master.GetMasterIP()
	res := proto_resolver.MasterMessage{
		Ip: currentIp,
	}
	// fmt.Println("============")
	// fmt.Println(currentIp)
	// p, _ := peer.FromContext(ctx)
	// fmt.Println("Calling get master from: " + string(p.Addr.String()))
	return &res, nil
}
