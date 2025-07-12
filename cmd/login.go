package cmd

import (
	"context"
	"fmt"
	"github.com/djcopley/zing/api"
	"github.com/djcopley/zing/config"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"syscall"
)

var username string

var loginCommand = &cobra.Command{
	Use:   "login [server]",
	Short: "Login to the server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addr := args[0]

		if username == "" {
			fmt.Print("Username: ")
			_, err := fmt.Scanln(&username)
			if err != nil {
				log.Fatalf("failed to read username: %s\n", err)
			}
		}

		fmt.Print("Password: ")
		passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatalf("failed to read password: %s\n", err)
		}
		password := string(passwordBytes)

		conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("failed to connect to server: %s\n", err)
		}
		defer conn.Close()
		c := api.NewZingClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		response, err := c.Login(ctx, &api.LoginRequest{Username: username, Password: password})
		if err != nil {
			log.Fatalf("failed to login to server: %s\n", err)
		}

		// Store token in config
		tokenStr := response.GetToken()
		if err := config.SetToken(tokenStr); err != nil {
			log.Printf("Warning: Failed to save token: %s\n", err)
		}

		log.Println("Login successful. Token stored.")
	},
}

func init() {
	rootCmd.AddCommand(loginCommand)
	loginCommand.Flags().StringVarP(&username, "username", "u", "", "Username to login with")
}
