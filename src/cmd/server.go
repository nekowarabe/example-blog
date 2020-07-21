package cmd

import (
	"app/src/server"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:           "server",
	Short:         "Start server",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE:          func(_ *cobra.Command, _ []string) error { return server.Run() },
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
