package cmd

import (
	"log"
	"os"
	"path"

	"brage/utils"
	"github.com/spf13/cobra"
)

var overwriteFiles bool

func runInitCommand(cmd *cobra.Command, args []string) {
	logger := log.Default()

	var targetPath string
	if len(args) > 0 {
		targetPath = args[0]
	} else {
		targetPath = "."
	}
	targetPath = utils.AbsolutePath(targetPath)

	logger.Printf("Creating site in: %v", targetPath)

	files := map[string]string{
		"config.yaml": `title: My Site
description: This is my Brage site.
image: dog.png
rootUrl: https://example.org

data:
  words:
    - banana
    - happy
    - explosion
`,
		"layout.html": `<!DOCTYPE html>
<html>
<head>
<link rel="stylesheet" type="text/css" href="/assets/style.css">
<title>{{ .Site.Title }}</title>
</head>

<body>
<h1>{{ .Site.Title }}</h1>
{{ .Content }}
</body>
</html>
`,
		"assets/style.css": `body {
background: #eee;
color: #222;
font-family: Helvetica, sans-serif;
}`,
		"pages/index.html": `<p>This is the main page.</p>
{{ template "extra" . }}
<a href="/about">About</a>
`,
		"pages/about.markdown": `This is the about page.

[Home](/)
`,
		"templates/extra.html": `<ul>
{{ range .Data.words }}
<li>{{ . }}</li>
{{ end }}
</ul>
`,
	}

	for filePath, contents := range files {
		fullPath := path.Join(targetPath, filePath)

		if err := os.MkdirAll(path.Dir(fullPath), 0755); err != nil {
			logger.Fatalf("ERROR! Unable to create directory: %v", err)
		}

		if !overwriteFiles {
			if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
				logger.Fatalf("ERROR! File already exists: %v", fullPath)
			}
		}

		if err := utils.WriteFile(fullPath, contents); err != nil {
			logger.Fatalf("ERROR! Unable to create file: %v", err)
		}

		logger.Printf("Created: %v", filePath)
	}
}

var initCommand = &cobra.Command{
	Use:   "init [PATH]",
	Short: "Initialize a new blank site",
	Long:  "Initialize a new blank site at the specified path (defaults to current directory)",
	Args:  cobra.MaximumNArgs(1),
	Run:   runInitCommand,
}

func init() {
	initCommand.Flags().BoolVarP(&overwriteFiles, "force", "f", false, "Force creating of site - overwriting existing files")

	rootCmd.AddCommand(initCommand)
}
