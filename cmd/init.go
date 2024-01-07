package cmd

import (
	"log"
	"os"
	"path"

	"github.com/michaelenger/brage/files"
	"github.com/spf13/cobra"
)

const ABOUT_TEMPLATE = `This is the about page.

[Home](/)
`

const CONFIG_TEMPLATE = `title: My Site
description: This is my Brage site.
image: dog.png
root_url: https://example.org
redirects:
  /example: https://example.org/

data:
  words:
    - banana
    - happy
    - explosion
`

const EXTRA_TEMPLATE = `<ul>
{{ #data.words }}
<li>{{ . }}</li>
{{ /data.words }}
</ul>
`

const FIRST_TEMPLATE = `---
title: First Post
date: 2023-07-21
---
This is my **first** post!
`

const INDEX_TEMPLATE = `<p>This is the main page.</p>
<h3>Words I like</h3>
{{ > extra }}

<h3>Posts</h3>
<ul>
	{{# site.posts }}
		<li>{{ date }} <a href="{{ path }}">{{ title }}</a></li>
	{{/ site.posts }}
</ul>

<a href="/about">About</a>
`

const LAYOUT_TEMPLATE = `<!DOCTYPE html>
<html>
<head>
<link rel="stylesheet" type="text/css" href="/assets/style.css">
<title>{{ site.title }}</title>
</head>

<body>
<h1>{{ site.title }}</h1>
{{{ content }}}
</body>
</html>
`

const STYLE_TEMPLATE = `body {
background: #eee;
color: #222;
font-family: Helvetica, sans-serif;
}`

var overwriteFiles bool

func runInitCommand(cmd *cobra.Command, args []string) {
	logger := log.Default()

	var targetPath string
	if len(args) > 0 {
		targetPath = args[0]
	} else {
		targetPath = "."
	}
	targetPath = files.AbsolutePath(targetPath)

	logger.Printf("Creating site in: %v", targetPath)

	siteFiles := map[string]string{
		"config.yaml":          CONFIG_TEMPLATE,
		"layout.html":          LAYOUT_TEMPLATE,
		"assets/style.css":     STYLE_TEMPLATE,
		"pages/index.html":     INDEX_TEMPLATE,
		"pages/about.markdown": ABOUT_TEMPLATE,
		"posts/first.markdown": FIRST_TEMPLATE,
		"partials/extra.html":  EXTRA_TEMPLATE,
	}

	for filePath, contents := range siteFiles {
		fullPath := path.Join(targetPath, filePath)

		if err := os.MkdirAll(path.Dir(fullPath), 0755); err != nil {
			logger.Fatalf("ERROR! Unable to create directory: %v", err)
		}

		if !overwriteFiles {
			if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
				logger.Fatalf("ERROR! File already exists: %v", fullPath)
			}
		}

		if err := files.WriteFile(fullPath, contents); err != nil {
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
