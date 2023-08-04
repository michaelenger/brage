package site

import (
	"fmt"
	"os"
	"path"

	"brage/files"
	"gopkg.in/yaml.v2"
)

type DataMap map[interface{}]interface{}

type SiteConfig struct {
	Title       string
	Description string
	Image       string
	RootUrl     string `yaml:"root_url"`
	Redirects   map[string]string
	Data        DataMap
}

type Site struct {
	Config          SiteConfig
	SourceDirectory string
	Layout          string
	Pages           []Page
	Partials        map[string]string
	Posts           []Post
}

// Load partials from the given directory.
func loadPartials(dirPath string) (map[string]string, error) {
	partials := map[string]string{}

	partialsFileInfo, err := os.Stat(dirPath)
	if err != nil || !partialsFileInfo.IsDir() {
		return partials, nil
	}

	partialFiles, err := files.ReadFiles(dirPath, "")
	if err != nil {
		return partials, err
	}
	for name, file := range partialFiles {
		if path.Base(name) == "index" {
			name = path.Clean(name[:len(name)-5])
		}

		partials[name] = file.Render()
	}

	return partials, nil
}

// Load posts from the given directory.
func loadPosts(dirPath string) []Post {
	posts := []Post{}

	postsFileInfo, err := os.Stat(dirPath)
	if err != nil || !postsFileInfo.IsDir() {
		return posts
	}

	postFiles, err := files.ReadFiles(dirPath, "")
	for _, file := range postFiles {
		posts = append(posts, MakePost(file, "/blog"))
	}

	return posts
}

// Load the site config based on a specified path and build the site description.
func Load(siteDirectory string) (Site, error) {
	var site Site

	if _, err := os.Stat(siteDirectory); os.IsNotExist(err) {
		return site, fmt.Errorf("No site directory found at specified path: %v", siteDirectory)
	}

	// Config

	configPath := path.Join(siteDirectory, "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return site, fmt.Errorf("No site config found at specified path: %v", configPath)
	}

	contents, err := os.ReadFile(configPath)
	if err != nil {
		return site, err
	}

	err = yaml.Unmarshal([]byte(contents), &site.Config)
	if err != nil {
		return site, err
	}

	// Layout

	layoutPath := path.Join(siteDirectory, "layout.html")
	if _, err := os.Stat(layoutPath); os.IsNotExist(err) {
		site.Layout = "{{{ content }}}"
	} else {
		contents, err = os.ReadFile(layoutPath)
		if err != nil {
			return site, fmt.Errorf("Unable to load layout template at path: %v", layoutPath)
		}
		site.Layout = string(contents)
	}

	// Pages

	pagesPath := path.Join(siteDirectory, "pages")
	if _, err := os.Stat(pagesPath); os.IsNotExist(err) {
		return site, fmt.Errorf("No pages found at specified path: %v", pagesPath)
	}

	pageFiles, err := files.ReadFiles(pagesPath, "/")
	if err != nil {
		return site, err
	}

	site.SourceDirectory = siteDirectory
	for name, file := range pageFiles {
		if path.Base(name) == "index" {
			name = path.Clean(name[:len(name)-5])
		}

		template := file.Render()

		site.Pages = append(site.Pages, Page{name, template})
	}

	// Partials

	site.Partials, err = loadPartials(path.Join(siteDirectory, "partials"))
	if err != nil {
		return site, err
	}

	// Posts

	site.Posts = loadPosts(path.Join(siteDirectory, "posts"))

	return site, nil
}

// Make the site context used when rendering pages and posts.
func (site Site) MakeContext() map[string]interface{} {
	return map[string]interface{}{
		"title":       site.Config.Title,
		"description": site.Config.Description,
		"image":       site.Config.Image,
		"root_url":    site.Config.RootUrl,
		"redirects":   site.Config.Redirects,
	}
}
