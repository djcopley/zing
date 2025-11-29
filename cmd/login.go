package cmd

import (
	"fmt"
	"syscall"

	"github.com/djcopley/zing/internal/api"
	"github.com/djcopley/zing/internal/client"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var (
	username      string
	insecureFlag  bool
	plaintextFlag bool
)

var loginCmd = &cobra.Command{
	Use:   "login [server]",
	Short: "Login to a server",
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

		_, err := fmt.Fprintf(cmd.OutOrStdout(), "Password: ")
		if err != nil {
			return fmt.Errorf("failed to read password: %s", err)
		}
		passwordBytes, err := term.ReadPassword(syscall.Stdin)
		if err != nil {
			return fmt.Errorf("failed to read password: %s", err)
		}
		password := string(passwordBytes)

		// Move to a new line after hidden input before proceeding
		_, _ = fmt.Fprintln(cmd.OutOrStdout())

		addr := args[0]
		client, err := client.NewClient(addr, insecureFlag, plaintextFlag)
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
		if err := config.SetServerAddr(addr); err != nil {
			return fmt.Errorf("failed to save server address: %s", err)
		}
		if err := config.SetPlaintext(plaintextFlag); err != nil {
			return fmt.Errorf("failed to save plaintext setting: %s", err)
		}
		if err := config.SetInsecure(insecureFlag); err != nil {
			return fmt.Errorf("failed to save insecure setting: %s", err)
		}

		_, err = fmt.Fprintln(cmd.OutOrStdout(), "Login successful. Token stored.")
		return err
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&username, "username", "u", "", "Username to login with")
	rootCmd.PersistentFlags().BoolVarP(&plaintextFlag, "plaintext", "p", false, "Use plaintext connection (no TLS)")
	rootCmd.PersistentFlags().BoolVarP(&insecureFlag, "insecure", "k", false, "Allow invalid TLS certificates (skip verification)")
}
