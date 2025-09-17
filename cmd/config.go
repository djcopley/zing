package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage zing configuration",
}

func init() {
	rootCmd.AddCommand(configCmd)
}
