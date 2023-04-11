package handler

import (
	"jdramirezl/proyecto-1-topicos/mom/internal/mom"
	"jdramirezl/proyecto-1-topicos/mom/internal/proto/cluster"
	"jdramirezl/proyecto-1-topicos/mom/internal/proto/message"
	proto_resolver "jdramirezl/proyecto-1-topicos/mom/internal/proto/resolver"
	resolver "jdramirezl/proyecto-1-topicos/mom/internal/resolver"
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
	proto_resolver.UnimplementedResolverServiceServer
	Master *resolver.Master
}

type Handler struct {
	QueueService    *MessageService
	ClusterService  *ClusterService
	ResolverService *ResolverService
}

func NewHandler(momService mom.MomService, master *resolver.Master) *Handler {
	return &Handler{
		QueueService:    &MessageService{momService: momService},
		ClusterService:  &ClusterService{momService: momService},
		ResolverService: &ResolverService{Master: master},
	}
}
