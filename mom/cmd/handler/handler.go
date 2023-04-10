package handler

import (
	"jdramirezl/proyecto-1-topicos/mom/internal/mom"
	"jdramirezl/proyecto-1-topicos/mom/internal/proto/cluster"
	"jdramirezl/proyecto-1-topicos/mom/internal/proto/message"
	"jdramirezl/proyecto-1-topicos/mom/internal/proto/resolver"
)

type MessageService struct {
	message.UnimplementedMessageServiceServer
	momService mom.MomService
}

type ClusterService struct {
	cluster.UnimplementedClusterServiceServer
	momService mom.MomService
}

type ResolverService struct {
	resolver.UnimplementedResolverServiceServer
	momService mom.MomService
}

type Handler struct {
	QueueService    *MessageService
	ClusterService  *ClusterService
	ResolverService *ResolverService
}

func NewHandler(momService mom.MomService) *Handler {
	return &Handler{
		QueueService:   &MessageService{momService: momService},
		ClusterService: &ClusterService{momService: momService},
	}
}
