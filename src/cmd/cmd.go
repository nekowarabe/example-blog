package cmd

import (
	"log"
	"os"

	"app/src/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "",
	Run: func(cmd *cobra.Command, _ []string) { cmd.Help() },

	Short: "A simple web server.",
}

func init() {
	cobra.OnInitialize(config.Init)
}

// Execute 執行應用程式
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
