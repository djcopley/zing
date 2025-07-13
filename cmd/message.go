package cmd

import (
	"context"
	"github.com/djcopley/zing/api"
	"github.com/djcopley/zing/client"
	"github.com/djcopley/zing/config"
	"github.com/djcopley/zing/editor"
	"github.com/spf13/cobra"
	"log"
)

var messageCommand = &cobra.Command{
	Use:   "message [user]",
	Short: "Message a user",
	Args:  cobra.ExactArgs(1),
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

		message, err := editor.ComposeMessage()
		if err != nil {
			log.Fatalf("failed to compose message: %s\n", err)
		}

		user := args[0]
		_, err = client.SendMessage(ctx, &api.SendMessageRequest{
			Token:   token,
			To:      &api.User{Username: user},
			Message: &api.Message{Content: message},
		})
		if err != nil {
			log.Fatalf("failed to send message: %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(messageCommand)
}
