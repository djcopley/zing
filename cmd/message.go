package cmd

import (
	"github.com/spf13/cobra"
)

// messageCmd represents the message command
var messageCmd = &cobra.Command{
	Use:   "message",
	Short: "A brief description of your command",
}

func init() {
	rootCmd.AddCommand(messageCmd)
}
