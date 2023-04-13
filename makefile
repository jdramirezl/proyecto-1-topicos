.PHONY: proto

proto:
	@protoc --go_out=./mom/internal --go-grpc_out=./mom/internal ./mom/proto/cluster.proto
	@protoc --go_out=./mom/internal --go-grpc_out=./mom/internal ./mom/proto/message.proto
	@protoc --go_out=./mom/internal --go-grpc_out=./mom/internal ./mom/proto/resolver.proto


gateway:
	@docker-compose up -f ./docker-compose.gateway.yml

consumer:
	@docker-compose up -f ./docker-compose.consumer.yml

resolver:
	@docker-compose up -f ./docker-compose.resolver.yml

node1:
	@docker-compose up -f ./docker-compose.node1.yml

node2:
	@docker-compose up -f ./docker-compose.node2.yml

node3:
	@docker-compose up -f ./docker-compose.node3.yml
