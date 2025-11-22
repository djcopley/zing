package cmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	api2 "github.com/djcopley/zing/internal/api"
	"github.com/djcopley/zing/internal/client"
	"github.com/djcopley/zing/internal/config"
	"github.com/djcopley/zing/internal/editor"
	"github.com/spf13/cobra"
)

var messageFlag string

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
			return fmt.Errorf("unable to create client: %w", err)
		}
		defer client.Close()

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		// Determine message content priority: flag -> piped stdin -> editor
		message := messageFlag
		if message == "" {
			// Check if there is piped stdin
			if fi, err := os.Stdin.Stat(); err == nil && (fi.Mode()&os.ModeCharDevice) == 0 {
				// Data is being piped in; read it all
				stdinBytes, err := io.ReadAll(os.Stdin)
				if err != nil {
					return fmt.Errorf("failed to read stdin: %s", err)
				}
				message = string(stdinBytes)
			}
		}
		if message == "" {
			var err error
			message, err = editor.ComposeMessage()
			if err != nil {
				return fmt.Errorf("failed to compose message: %s", err)
			}
		}

		if message == "" {
			return fmt.Errorf("message is empty")
		}

		if !strings.HasSuffix(message, "\n") {
			message += "\n"
		}

		user := args[0]
		_, err = client.SendMessage(ctx, &api2.SendMessageRequest{
			To:      &api2.User{Username: user},
			Message: &api2.Message{Content: message},
		})
		if err != nil {
			return fmt.Errorf("failed to send message: %s", err)
		}
		return nil
	},
}

func init() {
	messageSendCmd.Flags().StringVarP(&messageFlag, "message", "m", "", "Message content to send. If omitted, reads from piped stdin or opens the editor.")
	messageCmd.AddCommand(messageSendCmd)
}
