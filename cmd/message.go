package cmd

import (
	"context"
	"github.com/djcopley/zing/api"
	"github.com/djcopley/zing/config"
	"github.com/djcopley/zing/editor"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	to string
)

var messageCommand = &cobra.Command{
	Use:   "message",
	Short: "Message a user",
	Run: func(cmd *cobra.Command, args []string) {
		token := config.GetToken()
		if token == "" {
			log.Fatalf("No authentication token found. Please login first.")
		}

		addr := config.GetServerAddr()
		conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("failed to connect to server: %s\n", err)
		}
		defer conn.Close()
		c := api.NewZingClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		message, err := editor.ComposeMessage()
		if err != nil {
			log.Fatalf("failed to compose message: %s\n", err)
		}

		_, err = c.SendMessage(ctx, &api.SendMessageRequest{
			Token:   token,
			To:      &api.User{Username: to},
			Message: &api.Message{Content: message},
		})
		if err != nil {
			log.Fatalf("failed to send message: %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(messageCommand)
	messageCommand.Flags().StringVarP(&to, "to", "t", "", "User to send the message to")
}
