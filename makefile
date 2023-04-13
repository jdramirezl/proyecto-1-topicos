.PHONY: proto gateway consumer resolver node1 node2 node3

proto:
	@protoc --go_out=./mom/internal --go-grpc_out=./mom/internal ./mom/proto/cluster.proto
	@protoc --go_out=./mom/internal --go-grpc_out=./mom/internal ./mom/proto/message.proto
	@protoc --go_out=./mom/internal --go-grpc_out=./mom/internal ./mom/proto/resolver.proto


gateway:
	@docker-compose -f ./docker-compose.gateway.yml up

consumer:
	@docker-compose -f ./docker-compose.consumer.yml up

resolver:
	@docker-compose -f ./docker-compose.resolver.yml up

node1:
	@docker-compose -f ./docker-compose.node1.yml up

node2:
	@docker-compose -f ./docker-compose.node2.yml up

node3:
	@docker-compose -f ./docker-compose.node3.yml up
