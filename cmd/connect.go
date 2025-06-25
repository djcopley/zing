package cmd

import (
	"context"
	"fmt"
	"github.com/djcopley/zing/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func init() {
	rootCmd.AddCommand(connectCommand)
}

var connectCommand = &cobra.Command{
	Use:   "connect",
	Short: "Connect to the server",
	Run: func(cmd *cobra.Command, args []string) {
		addr := fmt.Sprintf("%s:%d", host, port)
		conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("failed to connect to server: %s\n", err)
		}
		defer conn.Close()
		c := api.NewZingClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		r, err := c.GetMessages(ctx, &api.GetMessagesRequest{})
		if err != nil {
			log.Fatalf("failed to connect to server: %s\n", err)
		}
		res, err := r.Recv()
		if err != nil {
			log.Fatalf("failed to receive message from server: %s\n", err)
		}
		log.Println(res)
	},
}
