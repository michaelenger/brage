package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "brage",
	Short: "Brage is a simple Static Site Generator",
}

func Execute() error {
	return rootCmd.Execute()
}
