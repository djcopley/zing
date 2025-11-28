package cmd

import (
	"log"

	configpkg "github.com/djcopley/zing/internal/config"
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
	config        *configpkg.Config
)

func init() {
	config = configpkg.NewConfig()

	rootCmd.PersistentFlags().BoolVarP(&insecureFlag, "insecure", "k", false, "Allow invalid TLS certificates (skip verification)")
	rootCmd.PersistentFlags().BoolVarP(&plaintextFlag, "plaintext", "p", false, "Use plaintext connection (no TLS)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
