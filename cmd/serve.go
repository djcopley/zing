package cmd

import (
	"fmt"
	"github.com/djcopley/zing/api"
	"github.com/djcopley/zing/server"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"net"
	"os"
)

func init() {
	rootCmd.AddCommand(serveCommand)
}

var serveCommand = &cobra.Command{
	Use:   "serve",
	Short: "Serve the zing server",
	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen("tcp", ":8080")
		if err != nil {
			fmt.Printf("failed to listen: %v", err)
			os.Exit(1)
		}
		s := grpc.NewServer()
		api.RegisterZingServer(s, &server.Server{})
		if err := s.Serve(lis); err != nil {
			fmt.Printf("failed to serve: %v", err)
			os.Exit(1)
		}
	},
}
