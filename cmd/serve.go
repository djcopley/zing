package cmd

import (
	"fmt"
	"github.com/djcopley/zing/api"
	"github.com/djcopley/zing/server"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	rootCmd.AddCommand(serveCommand)
}

var serveCommand = &cobra.Command{
	Use:   "serve",
	Short: "Serve the zing server",
	Run: func(cmd *cobra.Command, args []string) {
		addr := fmt.Sprintf("%s:%d", host, port)
		log.Printf("starting zing server @ %s\n", addr)
		lis, err := net.Listen("tcp", fmt.Sprintf("%s", addr))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		api.RegisterZingServer(s, &server.Server{})
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}
