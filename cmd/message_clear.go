package cmd

import (
	"fmt"

	"github.com/djcopley/zing/internal/api"
	"github.com/djcopley/zing/internal/client"
	"github.com/djcopley/zing/internal/config"
	"github.com/spf13/cobra"
)

var messageClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all messages sent to you",
	RunE: func(cmd *cobra.Command, args []string) error {
		token := config.GetToken()
		if token == "" {
			return fmt.Errorf("authentication token is required; please login first")
		}
		ctx := client.AddAuthMetadata(cmd.Context(), token)

		addr := config.GetServerAddr()
		c, err := client.NewClient(addr, insecureFlag, plaintextFlag)
		if err != nil {
			return fmt.Errorf("failed to connect to server: %s", err)
		}
		defer c.Close()

		_, err = c.ClearMessages(ctx, &api.ClearMessagesRequest{})
		if err != nil {
			return fmt.Errorf("failed to clear messages: %s", err)
		}
		_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Messages cleared.")
		return nil
	},
}

func init() {
	messageCmd.AddCommand(messageClearCmd)
}
