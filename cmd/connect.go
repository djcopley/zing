package cmd

import (
	"context"
	"errors"
	"github.com/djcopley/zing/api"
	"github.com/djcopley/zing/client"
	"github.com/djcopley/zing/config"
	"github.com/spf13/cobra"
	"io"
	"log"
)

var connectCommand = &cobra.Command{
	Use:   "connect",
	Short: "Connect to the server",
	Run: func(cmd *cobra.Command, args []string) {
		token := config.GetToken()
		if token == "" {
			log.Fatalf("No authentication token found. Please login first.")
		}

		addr := config.GetServerAddr()
		client, err := client.NewInsecureClient(addr)
		if err != nil {
			log.Fatalln(err)
		}
		defer client.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		request := &api.GetMessagesRequest{
			Token: token,
		}
		r, err := client.GetMessages(ctx, request)
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
}
