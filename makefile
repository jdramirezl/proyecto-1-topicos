.PHONY: proto

proto:
	@protoc --go_out=./mom/internal --go-grpc_out=./mom/internal ./mom/proto/cluster.proto
	@protoc --go_out=./mom/1/internal --go-grpc_out=./mom/internal ./mom/proto/message.proto
