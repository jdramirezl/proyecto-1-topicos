package main

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)


func List(conn *grpc.ClientConn) (*m1.FileResponse, error) {
	client := m1.NewFileServiceClient(conn)
	listrequest := m1.ListRequest{}
	res, err := client.List(context.Background(), &listrequest)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreateConnection(url string) (*grpc.ClientConn, error) {
	return grpc.Dial(url, grpc.WithInsecure())
}

func runGRPC(listenAddr string) *grpc.ClientConn {
	// GRPC
	fmt.Println("Connecting to GRPC server at", listenAddr)
	connGRPC, err := CreateConnection(listenAddr)
	failOnError(err, "Failed to connect to GRPC")
	
	fmt.Println("Connected successfully")
	return connGRPC
}