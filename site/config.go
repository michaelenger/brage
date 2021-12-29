package site

import (
	"fmt"
	"os"
	"path"

	"github.com/michaelenger/brage/utils"
	"gopkg.in/yaml.v2"
	"regexp"
)

type DataMap map[interface{}]interface{}

type SiteConfig struct {
	Title       string
	Description string
	Image       string
	RootUrl     string `yaml:"rootUrl"`

	Path  string
	Pages []string

	Data DataMap
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

// Load the site config based on a specified path.
func Load(sitePath string) (SiteConfig, error) {
	var config SiteConfig

	if _, err := os.Stat(sitePath); os.IsNotExist(err) {
		return config, fmt.Errorf("No site found at specified path: %v", sitePath)
	}

	configPath := path.Join(sitePath, "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return config, fmt.Errorf("No site config found at specified path: %v", configPath)
	}

	pagesPath := path.Join(sitePath, "pages")
	if _, err := os.Stat(pagesPath); os.IsNotExist(err) {
		return config, fmt.Errorf("No pages found at specified path: %v", pagesPath)
	}

	pages, err := utils.ListPages(pagesPath, "/")
	if err != nil {
		return config, err
	}

	config.Path = sitePath
	for k := range pages {
		config.Pages = append(config.Pages, k)
	}

	contents, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal([]byte(contents), &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
