.PHONY: proto

proto:
	@protoc --go_out=./mom/internal --go-grpc_out=./mom/internal ./proto/cluster.proto
	@protoc --go_out=./gateway/internal --go-grpc_out=./gateway/internal ./proto/cluster.proto
	@protoc --go_out=./mom/internal --go-grpc_out=./mom/internal ./proto/message.proto
