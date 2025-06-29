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
	username string
	password string
)

var loginCommand = &cobra.Command{
	Use:   "login",
	Short: "Login to the server",
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
		token, err := c.Login(ctx, &api.LoginRequest{Username: username, Password: password})
		if err != nil {
			log.Fatalf("failed to login to server: %s\n", err)
		}
		log.Println("token: ", token)
	},
}

func init() {
	rootCmd.AddCommand(loginCommand)
	loginCommand.Flags().StringVarP(&username, "username", "u", "", "username")
	loginCommand.Flags().StringVarP(&password, "password", "p", "", "password")
	loginCommand.MarkFlagRequired("username")
	loginCommand.MarkFlagRequired("password")
}
