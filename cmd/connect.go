package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(connectCommand)
}

var connectCommand = &cobra.Command{
	Use:   "connect",
	Short: "Connect to the server",
}
