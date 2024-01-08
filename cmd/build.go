package cmd

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path"

	"github.com/gorilla/feeds"
	"github.com/michaelenger/brage/files"
	"github.com/michaelenger/brage/site"
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

	sourcePath = files.AbsolutePath(sourcePath)
	destinationPath = files.AbsolutePath(destinationPath)

	logger.Printf("Loading site from: %v", sourcePath)

	siteData, err := site.Load(sourcePath)
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

		assets, err := files.CopyDirectory(assetsDirectory, destinationPath)
		if err != nil {
			logger.Fatalf("ERROR! Unable to copy assets: %v", err)
		}
		logger.Printf("Copied %v assets", assets)
	}

	for uri, targetUrl := range siteData.Config.Redirects {
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

		err = files.WriteFile(filePath, content)
		if err != nil {
			logger.Fatalf("ERROR! Unable to create redirect file: %v", err)
		}

		logger.Printf("Added redirect: %v => %v", uri, targetUrl)
	}

	logger.Printf("Found %d pages...", len(siteData.Pages))
	for _, page := range siteData.Pages {
		filePath := path.Join(destinationPath, page.Path, "index.html")
		content, err := page.Render(siteData)
		if err != nil {
			logger.Fatalf("ERROR! Unable to render page file: %v", err)
		}
		files.WriteFile(filePath, content)
		logger.Printf("Wrote file for: %v", page.Path)
	}

	logger.Printf("Found %d posts...", len(siteData.Posts))
	for _, post := range siteData.Posts {
		filePath := path.Join(destinationPath, post.Path, "index.html")
		content, err := post.Render(siteData)
		if err != nil {
			logger.Fatalf("ERROR! Unable to render post file: %v", err)
		}
		files.WriteFile(filePath, content)
		logger.Printf("Wrote file for: %v", post.Path)
	}

	if len(siteData.Posts) > 0 {
		author := &feeds.Author{Name: siteData.Config.Author}

		feed := &feeds.Feed{
			Title:       siteData.Config.Title,
			Link:        &feeds.Link{Href: siteData.Config.RootUrl},
			Description: siteData.Config.Description,
			Author:      author,
			Created:     siteData.Posts[len(siteData.Posts)-1].Date,
			Updated:     siteData.Posts[0].Date,
		}
		if siteData.Config.Image != "" {
			imageUrl, err := url.JoinPath(siteData.Config.RootUrl, "assets", siteData.Config.Image)
			if err != nil {
				logger.Fatalf("ERROR! Unable to create post URL: %v", err)
			}
			feed.Image = &feeds.Image{
				Url:   imageUrl,
				Title: siteData.Config.Title,
				Link:  siteData.Config.RootUrl,
			}
		}

		var feedItems []*feeds.Item
		for _, post := range siteData.Posts {
			postUrl, err := url.JoinPath(siteData.Config.RootUrl, post.Path)
			if err != nil {
				logger.Fatalf("ERROR! Unable to create post URL: %v", err)
			}
			feedItems = append(feedItems, &feeds.Item{
				Title:       post.Title,
				Link:        &feeds.Link{Href: postUrl},
				Description: post.Description,
				Author:      author,
				Created:     post.Date,
			})
		}
		feed.Items = feedItems

		filePath := path.Join(destinationPath, "feed.rss")
		content, err := feed.ToRss()
		if err != nil {
			logger.Fatalf("ERROR! Unable to generate RSS feed: %v", err)
		}
		files.WriteFile(filePath, content)
		logger.Print("Wrote RSS file: feed.rss")
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
