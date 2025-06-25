package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	host string
	port int
)

func init() {
	rootCmd.AddCommand(connectCommand)
	connectCommand.Flags().StringVarP(&host, "host", "H", "localhost", "Host to connect to")
	connectCommand.Flags().IntVarP(&port, "port", "P", 8080, "Port to connect to")
}

var connectCommand = &cobra.Command{
	Use:   "connect",
	Short: "Connect to the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Connecting to %s:%d\n", host, port)
	},
}
