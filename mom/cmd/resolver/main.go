package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"jdramirezl/proyecto-1-topicos/mom/cmd/handler"
	proto_resolver "jdramirezl/proyecto-1-topicos/mom/internal/proto/resolver"
	resolver "jdramirezl/proyecto-1-topicos/mom/internal/resolver"

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
	h := handler.NewHandler(nil)
	proto_resolver.RegisterResolverServiceServer(server, h.ResolverService)
	reflection.Register(server)
	Resolver := resolver.NewMaster()
	Resolver.GetMasterIP()

	log.Printf("Listening on port: %v\n", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
