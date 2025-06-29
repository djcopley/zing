package cmd

import (
	"fmt"
	"github.com/djcopley/zing/config"
	"github.com/spf13/cobra"
	"os"
)

var (
	host string
	port int
)

var rootCmd = &cobra.Command{
	Use:   "zing",
	Short: "Zing is a command line messenger",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&host, "host", "H", "localhost", "host")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "P", 5129, "port")
}

func Execute(conf *config.Config) {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
