package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "elasticsearch-tool",
	Short: "A CLI to manage elasticsearch",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error %s", err)
		os.Exit(1)
	}
}

