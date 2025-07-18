package cmd

import (
	"fmt"
	"github.com/djcopley/zing/api"
	"github.com/djcopley/zing/client"
	"github.com/djcopley/zing/config"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"syscall"
)

var username string

var loginCmd = &cobra.Command{
	Use:   "login [server]",
	Short: "Login to the server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if username == "" {
			_, err := fmt.Fprintf(cmd.OutOrStdout(), "Username: ")
			if err != nil {
				return fmt.Errorf("failed to read username: %s", err)
			}
			_, err = fmt.Fscanln(cmd.InOrStdin(), &username)
			if err != nil {
				return fmt.Errorf("failed to read username: %s", err)
			}
		}

		fmt.Print("Password: ")
		passwordBytes, err := term.ReadPassword(syscall.Stdin)
		if err != nil {
			return fmt.Errorf("failed to read password: %s", err)
		}
		password := string(passwordBytes)

		addr := args[0]
		client, err := client.NewInsecureClient(addr)
		if err != nil {
			return fmt.Errorf("failed to create client: %s", err)
		}
		defer client.Close()

		response, err := client.Login(cmd.Context(), &api.LoginRequest{Username: username, Password: password})
		if err != nil {
			return fmt.Errorf("failed to login to server: %s", err)
		}

		// Store token in config
		tokenStr := response.GetToken()
		if err := config.SetToken(tokenStr); err != nil {
			return fmt.Errorf("failed to save token: %s", err)
		}

		_, err = fmt.Fprintf(cmd.OutOrStdout(), "Login successful. Token stored.")
		return err
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&username, "username", "u", "", "Username to login with")
}
