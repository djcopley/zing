package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/djcopley/zing/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

var (
	connectToken string
)

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
		request := &api.GetMessagesRequest{
			Token: connectToken,
		}
		r, err := c.GetMessages(ctx, request)
		if err != nil {
			log.Fatalf("failed to connect to server: %s\n", err)
		}
		for {
			res, err := r.Recv()

			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				log.Fatalf("failed to receive message from server: %s\n", err)
			}

			log.Println(res)
		}
	},
}

func init() {
	rootCmd.AddCommand(connectCommand)
	connectCommand.Flags().StringVarP(&connectToken, "token", "T", "", "token")
}
