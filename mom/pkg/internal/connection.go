package connection

func NewGrpcClient(host, port string) (*grpc.ClientConn, error) {
	address := fmt.Sprintf("%v:%v", ip, port)
	return grpc.Dial(address, grpc.WithInsecure())
}

