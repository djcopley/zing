package cmd

import (
	"fmt"
	"github.com/djcopley/zing/config"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "zing",
	Short: "Zing is a command line messenger",
}

func Execute(conf *config.Config) {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
