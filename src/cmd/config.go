package cmd

import (
	"app/src/config"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Display config information",
	Run:   func(_ *cobra.Command, _ []string) { config.Display() },
}

func init() {
	rootCmd.AddCommand(configCmd)
}
