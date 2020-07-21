package cmd

import (
	"app/src/configs"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Display config information",
	Run:   func(_ *cobra.Command, _ []string) { configs.Display() },
}

func init() {
	rootCmd.AddCommand(configCmd)
}
