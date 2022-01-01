package cmd

import (
	"log"
	"os"
	"path"

	"github.com/michaelenger/brage/utils"
	"github.com/spf13/cobra"
)

var configContent = `title: My Site
description: This is my Brage site.
image: dog.png
rootUrl: https://example.org

data:
  words:
    - banana
    - happy
    - explosion
`

var layoutContent = `<!DOCTYPE html>
<html>
<head>
<link rel="stylesheet" type="text/css" href="/assets/style.css">
<title>{{ .Site.Title }}</title>
</head>

<body>
{{ .Content }}
</body>
</html>
`

var styleContent = `body {
	background: #eee;
	color: #222;
	font-family: Helvetica, sans-serif;
}`

var indexContent = `<p>This is the main page.</p>
{{ template "extra" . }}
<a href="/about">About</a>
`

var aboutContent = `<p>This is the about page.</p>
{{ template "extra" . }}
<a href="/">Home</a>
`

var extraContent = `<ul>
{{ range .Data.words }}
<li>{{ . }}</li>
{{ end }}
</ul>
`

var initCmd = &cobra.Command{
	Use:   "init [PATH]",
	Short: "Initialize a new blank site",
	Long:  "Initialize a new blank site at the specified path (defaults to current directory)",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logger := log.Default()

		var targetPath string
		if len(args) > 0 {
			targetPath = args[0]
		} else {
			targetPath = "."
		}
		targetPath = utils.AbsolutePath(targetPath)

		if _, err := os.Stat(targetPath); !os.IsNotExist(err) {
			logger.Fatalf("ERROR! Directory already exists: %v", targetPath)
		}

		logger.Printf("Creating site in: %v", targetPath)

		if err := os.MkdirAll(targetPath, 0755); err != nil {
			logger.Fatalf("ERROR! Unable to create target directory: %v", err)
		}

		// Config

		configFile := path.Join(targetPath, "config.yaml")
		if err := utils.WriteFile(configFile, configContent); err != nil {
			logger.Fatalf("ERROR! Unable to create config file: %v", err)
		}

		logger.Print("Created: /config.yaml")

		// Layout

		layoutFile := path.Join(targetPath, "layout.html")
		if err := utils.WriteFile(layoutFile, layoutContent); err != nil {
			logger.Fatalf("ERROR! Unable to create layout file: %v", err)
		}

		logger.Print("Created: /layout.html")

		// Assets

		assetsDirectory := path.Join(targetPath, "assets")
		if err := os.Mkdir(assetsDirectory, 0755); err != nil {
			logger.Fatalf("ERROR! Unable to create assets directory: %v", err)
		}

		logger.Print("Created: /assets")

		styleFile := path.Join(assetsDirectory, "style.css")
		if err := utils.WriteFile(styleFile, styleContent); err != nil {
			logger.Fatalf("ERROR! Unable to create style file: %v", err)
		}

		logger.Print("Created: /assets/style.css")

		// Pages

		pagesDirectory := path.Join(targetPath, "pages")
		if err := os.Mkdir(pagesDirectory, 0755); err != nil {
			logger.Fatalf("ERROR! Unable to create pages directory: %v", err)
		}

		logger.Print("Created: /pages")

		indexFile := path.Join(pagesDirectory, "index.html")
		if err := utils.WriteFile(indexFile, indexContent); err != nil {
			logger.Fatalf("ERROR! Unable to create index file: %v", err)
		}

		logger.Print("Created: /pages/index.html")

		aboutFile := path.Join(pagesDirectory, "about.html")
		if err := utils.WriteFile(aboutFile, aboutContent); err != nil {
			logger.Fatalf("ERROR! Unable to create about file: %v", err)
		}

		logger.Print("Created: /pages/about.html")

		// Templates

		templatesDirectory := path.Join(targetPath, "templates")
		if err := os.Mkdir(templatesDirectory, 0755); err != nil {
			logger.Fatalf("ERROR! Unable to create templates directory: %v", err)
		}

		logger.Print("Created: /templates")

		extraFile := path.Join(templatesDirectory, "extra.html")
		if err := utils.WriteFile(extraFile, extraContent); err != nil {
			logger.Fatalf("ERROR! Unable to create extra file: %v", err)
		}

		logger.Print("Created: /templates/extra.html")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
