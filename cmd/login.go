package cmd

import (
	"context"
	"fmt"
	"github.com/djcopley/zing/api"
	"github.com/djcopley/zing/client"
	"github.com/djcopley/zing/config"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"log"
	"syscall"
)

var username string

var loginCommand = &cobra.Command{
	Use:   "login [server]",
	Short: "Login to the server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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

		addr := args[0]
		client, err := client.NewInsecureClient(addr)
		if err != nil {
			log.Fatalln(err)
		}
		defer client.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		response, err := client.Login(ctx, &api.LoginRequest{Username: username, Password: password})
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
