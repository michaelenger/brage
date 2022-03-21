package cmd

import (
	"log"
	"os"
	"path"

	"brage/site"
	"brage/utils"
	"github.com/spf13/cobra"
)

// Path to send build files to
var destinationPath string

func runBuildCommand(cmd *cobra.Command, args []string) {
	logger := log.Default()

	var sourcePath string
	if len(args) > 0 {
		sourcePath = args[0]
	} else {
		sourcePath = "."
	}
	if destinationPath == "" {
		destinationPath = path.Join(sourcePath, "build")
	}

	sourcePath = utils.AbsolutePath(sourcePath)
	destinationPath = utils.AbsolutePath(destinationPath)

	logger.Printf("Loading site from: %v", sourcePath)

	site, err := site.Load(sourcePath)
	if err != nil {
		logger.Fatalf("ERROR! Unable to load site: %v", err)
	}

	logger.Printf("Building site in: %v", destinationPath)

	assetsDirectory := path.Join(sourcePath, "assets")
	if fileInfo, err := os.Stat(assetsDirectory); !os.IsNotExist(err) && fileInfo.IsDir() {
		assets, err := utils.CopyDirectory(assetsDirectory, destinationPath)
		if err != nil {
			logger.Fatalf("ERROR! Unable to copy assets: %v", err)
		}
		logger.Printf("Copied %v assets", assets)
	}

	for _, page := range site.Pages {
		filePath := path.Join(destinationPath, page.Path, "index.html")
		content, err := page.Render(site)
		if err != nil {
			logger.Fatalf("ERROR! Unable to render page file: %v", err)
		}
		utils.WriteFile(filePath, content)
		logger.Printf("Wrote file for: %v", page.Path)
	}
}

var buildCommand = &cobra.Command{
	Use:   "build [PATH]",
	Short: "Build the site",
	Long:  "Build the site",
	Args:  cobra.MaximumNArgs(1),
	Run:   runBuildCommand,
}

func init() {
	buildCommand.Flags().StringVarP(&destinationPath, "output", "o", "", "Directory to output files to")

	rootCmd.AddCommand(buildCommand)
}
