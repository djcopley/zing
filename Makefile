.PHONY: all server client test

all: test server client

server:
	go build -o build/server cmd/server/main.go

client:
	go build -o ./build/client cmd/client/main.go

protos:
	protoc internal/api/* --go_out=./ --go-grpc_out=./ --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative

test:
	go test ./...
