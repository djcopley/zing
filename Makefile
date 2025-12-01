.PHONY: all zing protos test

all: test protos zing

zing:
	go build -o build/zing main.go

protos:
	protoc -I internal/api --go_out=internal/api --go-grpc_out=internal/api --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative internal/api/*.proto

test:
	go test ./...
