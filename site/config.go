package site

import (
	"errors"
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type SiteConfig struct {
	Title       string
	Description string
	Image       string
	RootUrl     string `yaml:"rootUrl"`

	Path  string
	Pages []string

	Data map[interface{}]interface{}
}

// Load the site config based on a specified path.
func Load(sitePath string) (SiteConfig, error) {
	var config SiteConfig

	if _, err := os.Stat(sitePath); os.IsNotExist(err) {
		return config, errors.New(fmt.Sprintf("No site found at specified path: %v", sitePath))
	}

	configPath := path.Join(sitePath, "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return config, errors.New(fmt.Sprintf("No site config found at specified path: %v", configPath))
	}

	config.Path = sitePath
	// TODO pages

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
