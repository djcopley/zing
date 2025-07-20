package cmd

import (
	"context"
	"fmt"
	"github.com/djcopley/zing/api"
	"github.com/djcopley/zing/client"
	"github.com/djcopley/zing/config"
	"github.com/djcopley/zing/editor"
	"github.com/spf13/cobra"
	"log"
)

var messageSendCmd = &cobra.Command{
	Use:   "send [user]",
	Short: "Send a message to a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		token := config.GetToken()
		if token == "" {
			return fmt.Errorf("authentication token is required; please login first")
		}
		ctx := client.AddAuthMetadata(cmd.Context(), token)

		addr := config.GetServerAddr()
		client, err := client.NewInsecureClient(addr)
		if err != nil {
			log.Fatalln(err)
		}
		defer client.Close()

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		message, err := editor.ComposeMessage()
		if err != nil {
			return fmt.Errorf("failed to compose message: %s", err)
		}

		user := args[0]
		_, err = client.SendMessage(ctx, &api.SendMessageRequest{
			To:      &api.User{Username: user},
			Message: &api.Message{Content: message},
		})
		if err != nil {
			return fmt.Errorf("failed to send message: %s", err)
		}
		return nil
	},
}

func init() {
	messageCmd.AddCommand(messageSendCmd)
}
