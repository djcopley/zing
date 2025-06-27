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

var (
	to      string
	token   string
	message string
)

func init() {
	rootCmd.AddCommand(messageCommand)
	messageCommand.Flags().StringVarP(&to, "to", "t", "", "User to send the message to")
	messageCommand.Flags().StringVarP(&token, "token", "T", "", "Credentials for user")
	messageCommand.Flags().StringVarP(&message, "message", "m", "", "Message to send")
}

var messageCommand = &cobra.Command{
	Use:   "message",
	Short: "Message a user",
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
