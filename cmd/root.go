package cmd

import (
	"log"

	configpkg "github.com/djcopley/zing/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zing",
	Short: "Zing is a command line messenger",
}

// Global connection flags
var (
	config *configpkg.Config
)

func init() {
	config = configpkg.NewConfig()

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
