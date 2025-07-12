package client

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/djcopley/zing/api"
	"github.com/djcopley/zing/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func NewSecureClient() (api.ZingClient, error) {
	// Load system root CAs
	certPool, err := x509.SystemCertPool()
	if err != nil {
		log.Fatalf("Failed to load system root CAs: %v", err)
	}

	// Create TLS config with system roots
	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}

	creds := credentials.NewTLS(tlsConfig)

	// Get server address from config
	addr := config.GetServerAddr()
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("failed to connect to server: %s\n", err)
	}
	defer conn.Close()
	c := api.NewZingClient(conn)

	return c, nil
}
