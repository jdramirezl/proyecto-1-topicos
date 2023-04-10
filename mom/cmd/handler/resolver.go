package handler

import (
	"context"
	proto_resolver "jdramirezl/proyecto-1-topicos/mom/internal/proto/resolver"
	resolver "jdramirezl/proyecto-1-topicos/mom/internal/resolver"

	"github.com/golang/protobuf/ptypes/empty"
)

func (c *ResolverService) NewMaster(ctx context.Context, req *proto_resolver.MasterMessage) (*empty.Empty, error) {
	master := resolver.GetMaster()
	newIP := req.Ip
	master.SetMaster(newIP)
	return &empty.Empty{}, nil
}

func (c *ResolverService) GetMaster(ctx context.Context, emp *empty.Empty) (*proto_resolver.MasterMessage, error) {
	master := resolver.GetMaster()
	currentIp := master.GetMasterIP()
	res := proto_resolver.MasterMessage{
		Ip: currentIp,
	}
	return &res, nil
}
