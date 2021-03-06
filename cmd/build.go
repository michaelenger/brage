package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"brage/site"
	"brage/utils"
	"github.com/spf13/cobra"
)

// Path to send build files to
var destinationPath string

// Whether to clean the assets dir
var cleanAssetDir bool

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
		if cleanAssetDir {
			err := os.RemoveAll(path.Join(destinationPath, "assets"))
			if err != nil {
				logger.Fatalf("ERROR! Unable to delete existing assets directory: %v", err)
			}
		}

		assets, err := utils.CopyDirectory(assetsDirectory, destinationPath)
		if err != nil {
			logger.Fatalf("ERROR! Unable to copy assets: %v", err)
		}
		logger.Printf("Copied %v assets", assets)
	}

	for uri, targetUrl := range site.Config.Redirects {
		filePath := path.Join(destinationPath, uri, "index.html")

		content := fmt.Sprintf(`<!DOCTYPE html>
<html>
	<head>
		<meta http-equiv="refresh" content="0; url='%s'" />
	</head>
	<body>
		<p>Sending you <a href="%s">here</a>.</p>
	</body>
</html>`, targetUrl, targetUrl)

		err = utils.WriteFile(filePath, content)
		if err != nil {
			logger.Fatalf("ERROR! Unable to create redirect file: %v", err)
		}

		logger.Printf("Added redirect: %v => %v", uri, targetUrl)
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
	buildCommand.Flags().BoolVarP(&cleanAssetDir, "clean", "c", false, "Clean the destination assets directory before building")

	rootCmd.AddCommand(buildCommand)
}
