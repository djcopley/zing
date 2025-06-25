package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(connectCommand)
}

var connectCommand = &cobra.Command{
	Use:   "connect",
	Short: "Connect to the server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("connecting to %s:%d\n", host, port)
	},
}
