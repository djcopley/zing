package cmd

import (
	"log"

	"github.com/djcopley/zing/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zing",
	Short: "Zing is a command line messenger",
}

func init() {
	config.InitConfig()
	if err := config.EnsureConfigFile(); err != nil {
		log.Printf("Warning: Could not create config file: %s\n", err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
