package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Port to serve the site on
var port string

var serveCmd = &cobra.Command{
	Use:   "serve [PATH]",
	Short: "Run a local server preview of the site",
	Long:  "Run a local server preview of the site",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var path string
		if len(args) > 0 {
			path = args[0]
		} else {
			path = "CURRENT DIR"
		}

		fmt.Printf("TODO: Serve %v on localhost:%v\n", path, port)
	},
}

func init() {
	serveCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to serve the site on")

	rootCmd.AddCommand(serveCmd)
}
