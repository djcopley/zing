package cmd

import (
	"fmt"
	"github.com/djcopley/zing/api"
	"github.com/djcopley/zing/client"
	"github.com/djcopley/zing/config"
	"github.com/spf13/cobra"
)

var messageReadCmd = &cobra.Command{
	Use:   "read",
	Short: "Read your messages",
	RunE: func(cmd *cobra.Command, args []string) error {
		token := config.GetToken()
		if token == "" {
			return fmt.Errorf("authentication token is required; please login first")
		}
		ctx := client.AddAuthMetadata(cmd.Context(), token)

		addr := config.GetServerAddr()
		client, err := client.NewInsecureClient(addr)
		if err != nil {
			return fmt.Errorf("failed to connect to server: %s", err)
		}
		defer client.Close()

		request := &api.ListMessagesRequest{}
		r, err := client.ListMessages(ctx, request)
		if err != nil {
			return fmt.Errorf("failed to connect to server: %s", err)
		}
		for _, message := range r.Messages {
			formattedMsg := fmt.Sprintf("%s: %s", message.Metadata.From.Username, message.Content)
			_, _ = fmt.Fprint(cmd.OutOrStdout(), formattedMsg)
		}
		return nil
	},
}

func init() {
	messageCmd.AddCommand(messageReadCmd)
}
