grpc-server:
	go run ./cmd/grpc/server/main.go

grpc-client:
	go run ./cmd/grpc/client/main.go

grpc-generate-protos:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    cmd/grpc/proto/greeter.proto

redis-server-up:
	docker-compose -f redis.yml up -d

redis-server-down:
	docker-compose -f redis.yml down

redis-client:
	go run ./cmd/redis/client/main.go

