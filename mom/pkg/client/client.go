package client

type Client struct {
	grpcConnection *grpc.ClientConn
}

func NewClient(host, port string) {
	grpcConn, err := client.NewGrpcClient(host, port)
	if err != nil {
		panic(err)
	}
	return Client{}
}

func (c *Client) Publish() {

}

func (c *Client) Subscribe() {

}