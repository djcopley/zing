package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/djcopley/zing/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
)

type Client struct {
	conn *grpc.ClientConn
	api.ZingClient
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func NewInsecureClient(address string) (*Client, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := api.NewZingClient(conn)
	return &Client{
		conn:       conn,
		ZingClient: client,
	}, nil
}

func AddAuthMetadata(ctx context.Context, token string) context.Context {
	md := metadata.New(map[string]string{"authorization": "Bearer " + token})
	return metadata.NewOutgoingContext(ctx, md)
}

func NewSecureClient(address string) (*Client, error) {
	certPool, err := x509.SystemCertPool()
	if err != nil {
		log.Fatalf("Failed to load system root CAs: %v", err)
	}

	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}
	creds := credentials.NewTLS(tlsConfig)

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}
	client := api.NewZingClient(conn)

	return &Client{
		conn:       conn,
		ZingClient: client,
	}, nil
}
