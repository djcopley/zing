package cmd

import (
	"log"

	"github.com/djcopley/zing/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zing",
	Short: "Zing is a command line messenger",
}

// Global connection flags
var (
	insecureFlag  bool
	plaintextFlag bool
)

func init() {
	config.InitConfig()
	if err := config.EnsureConfigFile(); err != nil {
		log.Printf("Warning: Could not create config file: %s\n", err)
	}

	rootCmd.PersistentFlags().BoolVarP(&insecureFlag, "insecure", "k", false, "Allow invalid TLS certificates (skip verification)")
	rootCmd.PersistentFlags().BoolVarP(&plaintextFlag, "plaintext", "p", false, "Use plaintext connection (no TLS)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
