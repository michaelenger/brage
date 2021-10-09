package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [PATH]",
	Short: "Initialize a new blank site",
	Long:  "Initialize a new blank site at the specified path (defaults to current directory)",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var path string
		if len(args) > 0 {
			path = args[0]
		} else {
			path = "CURRENT DIR"
		}

		fmt.Printf("TODO: Init %v\n", path)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
