package connection

import (
	"fmt"

	"google.golang.org/grpc"
)

func NewGrpcClient(host, port string) (*grpc.ClientConn, error) {
	address := fmt.Sprintf("%v:%v", host, port)
	return grpc.Dial(address, grpc.WithInsecure())
}
