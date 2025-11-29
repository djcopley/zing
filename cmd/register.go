/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"syscall"

	"github.com/djcopley/zing/internal/api"
	"github.com/djcopley/zing/internal/client"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register [server]",
	Short: "Create a new account on a server and store the login token",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if username == "" {
			if _, err := fmt.Fprintf(cmd.OutOrStdout(), "Username: "); err != nil {
				return fmt.Errorf("failed to prompt for username: %s", err)
			}
			if _, err := fmt.Fscanln(cmd.InOrStdin(), &username); err != nil {
				return fmt.Errorf("failed to read username: %s", err)
			}
		}

		if _, err := fmt.Fprintf(cmd.OutOrStdout(), "Password: "); err != nil {
			return fmt.Errorf("failed to prompt for password: %s", err)
		}
		passwordBytes, err := term.ReadPassword(syscall.Stdin)
		if err != nil {
			return fmt.Errorf("failed to read password: %s", err)
		}
		password := string(passwordBytes)

		// Confirm password
		if _, err := fmt.Fprintf(cmd.OutOrStdout(), "\nConfirm Password: "); err != nil {
			return fmt.Errorf("failed to prompt for password confirmation: %s", err)
		}
		confirmBytes, err := term.ReadPassword(syscall.Stdin)
		if err != nil {
			return fmt.Errorf("failed to read password confirmation: %s", err)
		}
		confirm := string(confirmBytes)
		if password != confirm {
			// Move to a new line after hidden input
			_, _ = fmt.Fprintln(cmd.OutOrStdout())
			return fmt.Errorf("passwords do not match")
		}

		// Move to a new line after hidden input before proceeding
		_, _ = fmt.Fprintln(cmd.OutOrStdout())

		addr := args[0]
		cl, err := client.NewClient(addr, insecureFlag, plaintextFlag)
		if err != nil {
			return fmt.Errorf("failed to create client: %s", err)
		}
		defer cl.Close()

		resp, err := cl.Register(cmd.Context(), &api.LoginRequest{Username: username, Password: password})
		if err != nil {
			return fmt.Errorf("failed to register: %s", err)
		}

		// Save connection settings and token
		tokenStr := resp.GetToken()
		if err := config.SetToken(tokenStr); err != nil {
			return fmt.Errorf("failed to save token: %s", err)
		}
		if err := config.SetServerAddr(addr); err != nil {
			return fmt.Errorf("failed to save server address: %s", err)
		}
		if err := config.SetPlaintext(plaintextFlag); err != nil {
			return fmt.Errorf("failed to save plaintext setting: %s", err)
		}
		if err := config.SetInsecure(insecureFlag); err != nil {
			return fmt.Errorf("failed to save insecure setting: %s", err)
		}

		_, err = fmt.Fprintln(cmd.OutOrStdout(), "Registration successful. Token stored.")
		return err
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	// Per-command flags
	registerCmd.Flags().StringVarP(&username, "username", "u", "", "Username to register with")
}
