# include .env
# export

run: build
	./bin/main

build: protoc
	go build -o ./bin cmd/main.go

protoc:
	protoc --go_out=./pkg/api/test \
		--go_opt=paths=source_relative \
		--go-grpc_out=./pkg/api/test --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./pkg/api/test --grpc-gateway_opt paths=source_relative \
		./api/*proto