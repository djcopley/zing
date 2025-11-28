package cmd

import (
	"fmt"

	"github.com/djcopley/zing/internal/editor"
	"github.com/spf13/cobra"
)

var configEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open the configuration file in your editor",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfgPath := config.ConfigFileUsed()
		if cfgPath == "" {
			return fmt.Errorf("config file path is not set")
		}
		return editor.Open(cfgPath)
	},
}

func init() {
	configCmd.AddCommand(configEditCmd)
}
