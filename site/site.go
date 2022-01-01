package site

import (
	"fmt"
	"os"
	"path"

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

type Site struct {
	Config          SiteConfig
	SourceDirectory string
	Pages           []Page
	Templates       map[string]string
}

// Load the site config based on a specified path and build the site description.
func Load(siteDirectory string) (Site, error) {
	var site Site

	if _, err := os.Stat(siteDirectory); os.IsNotExist(err) {
		return site, fmt.Errorf("No site found at specified path: %v", siteDirectory)
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

	// Pages

	pagesPath := path.Join(siteDirectory, "pages")
	if _, err := os.Stat(pagesPath); os.IsNotExist(err) {
		return site, fmt.Errorf("No pages found at specified path: %v", pagesPath)
	}

	pages, err := utils.ListTemplateFiles(pagesPath, "/")
	if err != nil {
		return site, err
	}

	site.SourceDirectory = siteDirectory
	for relativePath, filePath := range pages {
		if path.Base(relativePath) == "index" {
			relativePath = path.Clean(relativePath[:len(relativePath)-5])
		}
		site.Pages = append(site.Pages, Page{relativePath, filePath})
	}

	// Templates

	templatesPath := path.Join(siteDirectory, "templates")
	templatesFileInfo, _ := os.Stat(templatesPath)
	if !templatesFileInfo.IsDir() {
		return site, nil
	}

	templates, err := utils.ListTemplateFiles(templatesPath, "")
	if err != nil {
		return site, err
	}
	site.Templates = templates
	for relativePath, filePath := range templates {
		contents, err := os.ReadFile(filePath)
		if err != nil {
			return site, err
		}
		site.Templates[relativePath] = string(contents)
	}

	return site, nil
}
