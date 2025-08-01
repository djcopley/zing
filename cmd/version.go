package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of Zing!",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Zing! client v0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
