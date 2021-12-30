package cmd

import (
	"fmt"

	"github.com/michaelenger/brage/site"
	"github.com/michaelenger/brage/utils"
	"github.com/spf13/cobra"
)

// Path to send build files to
var outputPath string

var buildCmd = &cobra.Command{
	Use:   "build [PATH]",
	Short: "Build the site",
	Long:  "Build the site",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var path string
		if len(args) > 0 {
			path = args[0]
		} else {
			path = "."
		}
		if outputPath == "" {
			outputPath = "./build"
		}

		fmt.Printf("TODO: Build %v to %v\n", utils.AbsolutePath(path), utils.AbsolutePath(outputPath))
		conf, err := site.Load(utils.AbsolutePath(path))
		fmt.Printf("%+v or %v", conf, err)
	},
}

func init() {
	buildCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Directory to output files to")

	rootCmd.AddCommand(buildCmd)
}
