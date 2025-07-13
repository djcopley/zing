.PHONY: all zing protos test

all: test protos zing

zing:
	go build -o build/zing main.go

protos:
	protoc api/*.proto --go_out=./ --go-grpc_out=./ --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative

test:
	go test ./...
