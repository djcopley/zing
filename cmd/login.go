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

var loginCommand = &cobra.Command{
	Use:   "login [username] [password]",
	Short: "Login to the server",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			log.Fatal("username and password required")
		}
		username := args[0]
		password := args[1]

		addr := fmt.Sprintf("%s:%d", host, port)
		conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("failed to connect to server: %s\n", err)
		}
		defer conn.Close()
		c := api.NewZingClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		token, err := c.Login(ctx, &api.LoginRequest{Username: username, Password: password})
		if err != nil {
			log.Fatalf("failed to login to server: %s\n", err)
		}
		log.Println("token: ", token)
	},
}

func init() {
	rootCmd.AddCommand(loginCommand)
}
