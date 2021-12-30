package site

import (
	"fmt"
	"os"
	"path"
	"regexp"

	"github.com/michaelenger/brage/utils"
	"gopkg.in/yaml.v2"
)

type DataMap map[interface{}]interface{}

type SiteConfig struct {
	Title       string
	Description string
	Image       string
	RootUrl     string `yaml:"rootUrl"`
	Data        DataMap
}

type SiteDescription struct {
	Config          SiteConfig
	SourceDirectory string
	Pages           []Page
}

// Recursively build a list of pages.
func listPages(dirPath string, prefixPath string) ([]string, error) {
	pages := []string{}
	htmlFilePattern := regexp.MustCompile(`^(.+?)\.html$`)

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return pages, err
	}

	for _, file := range files {
		filename := file.Name()
		if file.IsDir() {
			_, err := listPages(path.Join(dirPath, filename), path.Join(prefixPath, filename))
			if err != nil {
				return pages, err
			}
		}

		filenameMatch := htmlFilePattern.FindStringSubmatch(filename)
		if len(filenameMatch) == 0 {
			continue
		}

		page := filenameMatch[1]
		if page == "index" {
			page = ""
		}

		fmt.Printf("%v => %v\n", filename, path.Join(prefixPath, page))
	}

	return pages, nil
}

// Load the site config based on a specified path and build the site description.
func LoadSite(siteDirectory string) (SiteDescription, error) {
	var description SiteDescription

	if _, err := os.Stat(siteDirectory); os.IsNotExist(err) {
		return description, fmt.Errorf("No site found at specified path: %v", siteDirectory)
	}

	configPath := path.Join(siteDirectory, "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return description, fmt.Errorf("No site config found at specified path: %v", configPath)
	}

	pagesPath := path.Join(siteDirectory, "pages")
	if _, err := os.Stat(pagesPath); os.IsNotExist(err) {
		return description, fmt.Errorf("No pages found at specified path: %v", pagesPath)
	}

	pages, err := utils.ListPages(pagesPath, "/")
	if err != nil {
		return description, err
	}

	description.SourceDirectory = siteDirectory
	for k, v := range pages {
		description.Pages = append(description.Pages, Page{k, v})
	}

	contents, err := os.ReadFile(configPath)
	if err != nil {
		return description, err
	}

	err = yaml.Unmarshal([]byte(contents), &description.Config)
	if err != nil {
		return description, err
	}

	return description, nil
}