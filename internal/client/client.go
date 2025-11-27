package client

import (
    "context"
    "crypto/tls"
    "crypto/x509"
    "log"

	"github.com/djcopley/zing/internal/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
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

// NewSecureClient creates a TLS client that verifies the server certificate using
// system root CAs. This is the default secure mode.
func NewSecureClient(address string) (*Client, error) {
    certPool, err := x509.SystemCertPool()
    if err != nil {
        log.Fatalf("Failed to load system root CAs: %v", err)
    }

    tlsConfig := &tls.Config{
        RootCAs:            certPool,
        InsecureSkipVerify: false,
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

// NewClient creates a client using the provided security options.
// If plaintext is true, a non-TLS (insecure) connection is used.
// Else, TLS is used. If insecure is true, TLS verification is skipped (InsecureSkipVerify).
func NewClient(address string, insecureFlag, plaintext bool) (*Client, error) {
    if plaintext {
        // Plaintext/no TLS
        conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
        if err != nil {
            return nil, err
        }
        c := api.NewZingClient(conn)
        return &Client{conn: conn, ZingClient: c}, nil
    }

    // TLS
    certPool, err := x509.SystemCertPool()
    if err != nil {
        log.Fatalf("Failed to load system root CAs: %v", err)
    }
    tlsConfig := &tls.Config{
        RootCAs:            certPool,
        InsecureSkipVerify: insecureFlag,
    }
    creds := credentials.NewTLS(tlsConfig)
    conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(creds))
    if err != nil {
        return nil, err
    }
    c := api.NewZingClient(conn)
    return &Client{conn: conn, ZingClient: c}, nil
}
