package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"jdramirezl/proyecto-1-topicos/mom/cmd/handler"
	"jdramirezl/proyecto-1-topicos/mom/internal/mom"
	"jdramirezl/proyecto-1-topicos/mom/internal/proto/cluster"
	"jdramirezl/proyecto-1-topicos/mom/internal/proto/message"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := os.Getenv("PORT")
	ip := "0.0.0.0"

	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", ip, port))
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	momService := mom.NewMomService()
	h := handler.NewHandler(momService, nil)
	message.RegisterMessageServiceServer(server, h.QueueService)
	cluster.RegisterClusterServiceServer(server, h.ClusterService)

	reflection.Register(server)
	log.Printf("Listening on port: %v\n", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
