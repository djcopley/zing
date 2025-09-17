package cmd

import (
	"fmt"
	"strings"

	"github.com/djcopley/zing/api"
	"github.com/djcopley/zing/client"
	"github.com/djcopley/zing/config"
	"github.com/djcopley/zing/pager"
	"github.com/spf13/cobra"
)

var pageSize int32
var pageToken string

var messageReadCmd = &cobra.Command{
	Use:   "read",
	Short: "Read the messages sent to you",
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

		request := &api.ListMessagesRequest{PageSize: pageSize, PageToken: pageToken}
		r, err := client.ListMessages(ctx, request)
		if err != nil {
			return fmt.Errorf("failed to connect to server: %s", err)
		}

		var b strings.Builder
		if len(r.Messages) == 0 {
			b.WriteString("No messages.\n")
		} else {
			for i, message := range r.Messages {
				b.WriteString(formatMessage(message))
				if i < len(r.Messages)-1 {
					b.WriteString(strings.Repeat("-", 60))
					b.WriteString("\n")
				}
			}
			if r.NextPageToken != "" {
				b.WriteString(strings.Repeat("-", 60))
				b.WriteString("\n")
				b.WriteString(fmt.Sprintf("More messages available. Fetch next page with: zing message read --page-token %s --page-size %d\n", r.NextPageToken, pageSize))
			}
		}

		p := pager.NewPager(cmd.OutOrStdout(), cmd.ErrOrStderr())
		if err := p.Page(b.String()); err != nil {
			// Fall back to plain stdout if pager is unavailable
			_, _ = fmt.Fprint(cmd.OutOrStdout(), b.String())
		}
		return nil
	},
}

func init() {
	messageCmd.AddCommand(messageReadCmd)
	messageReadCmd.Flags().Int32Var(&pageSize, "page-size", 0, "Number of messages per page (default 50, max 1000)")
	messageReadCmd.Flags().StringVar(&pageToken, "page-token", "", "Page token from a previous response to fetch the next page")
}

func formatMessage(message *api.Message) string {
	// Timestamp formatting (local time)
	var ts string
	if message.Metadata != nil && message.Metadata.Timestamp != nil {
		ts = message.Metadata.Timestamp.AsTime().Local().Format("2006-01-02 15:04:05")
	} else {
		ts = "(no timestamp)"
	}

	from := "unknown"
	if message.Metadata != nil && message.Metadata.From != nil && message.Metadata.From.Username != "" {
		from = message.Metadata.From.Username
	}

	header := fmt.Sprintf("[%s] From: %s", ts, from)
	content := message.Content
	if content == "" {
		content = "(no content)"
	}

	return fmt.Sprintf("%s\n%s", header, content)
}
