package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jdramirezl/proyecto-1-topicos/mom/cmd/handler"
	proto_resolver "github.com/jdramirezl/proyecto-1-topicos/mom/internal/proto/resolver"
	resolver "github.com/jdramirezl/proyecto-1-topicos/mom/internal/resolver"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	Resolver := resolver.NewMaster()
	// Resolver.GetMasterIP()

	port := os.Getenv("PORT")
	ip := "0.0.0.0"

	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", ip, port))
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	h := handler.NewHandler(nil, Resolver)
	proto_resolver.RegisterResolverServiceServer(server, h.ResolverService)
	reflection.Register(server)

	log.Printf("Listening on port: %v\n", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
