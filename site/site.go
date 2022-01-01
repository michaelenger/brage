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
	Layout          string
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

	// Layout

	layoutPath := path.Join(siteDirectory, "layout.html")
	if _, err := os.Stat(layoutPath); os.IsNotExist(err) {
		return site, fmt.Errorf("No layout template found at specified path: %v", layoutPath)
	}
	contents, err = os.ReadFile(layoutPath)
	if err != nil {
		return site, fmt.Errorf("Unable to load layout template at path: %v", layoutPath)
	}
	site.Layout = string(contents)

	// Pages

	pagesPath := path.Join(siteDirectory, "pages")
	if _, err := os.Stat(pagesPath); os.IsNotExist(err) {
		return site, fmt.Errorf("No pages found at specified path: %v", pagesPath)
	}

	pages, err := utils.LoadTemplateFiles(pagesPath, "/")
	if err != nil {
		return site, err
	}

	site.SourceDirectory = siteDirectory
	for name, template := range pages {
		if path.Base(name) == "index" {
			name = path.Clean(name[:len(name)-5])
		}
		site.Pages = append(site.Pages, Page{name, template})
	}

	// Templates

	templatesPath := path.Join(siteDirectory, "templates")
	templatesFileInfo, _ := os.Stat(templatesPath)
	if !templatesFileInfo.IsDir() {
		return site, nil
	}

	templates, err := utils.LoadTemplateFiles(templatesPath, "")
	if err != nil {
		return site, err
	}
	site.Templates = templates

	return site, nil
}
