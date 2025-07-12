package cmd

import (
	"fmt"
	"github.com/djcopley/zing/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "zing",
	Short: "Zing is a command line messenger",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// If flag is provided, override config value
		if cmd.Flags().Changed("server") {
			serverAddr, _ := cmd.Flags().GetString("server")
			viper.Set("server_addr", serverAddr)
		}
	},
}

func init() {
	// Initialize config
	config.InitConfig()

	// Ensure config file exists
	if err := config.EnsureConfigFile(); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not create config file: %s\n", err)
	}
}

func Execute(conf *config.Config) {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
