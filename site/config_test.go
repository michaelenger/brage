package site

import (
	"os"
	"path"
	"testing"
)

var exampleConfig = `
title: Young Fatigue
description: New Single ‘Dislocation‘ Out Now!
image: icon.png
rootUrl: https://youngfatigue.com/

data:
  instagram: https://www.instagram.com/youngfatigue/
  quotes:
    - It’s actually really good... really good.
    - "[Dislocation] is super fun!"
`

func createExampleSite() (string, error) {
	// Site directory
	temporaryDirectory, err := os.MkdirTemp("", "examplesite")
	if err != nil {
		return "", err
	}

	// Config file
	configFile, err := os.Create(path.Join(temporaryDirectory, "config.yaml"))
	if err != nil {
		return "", err
	}
	defer configFile.Close()
	_, err = configFile.WriteString(exampleConfig)
	if err != nil {
		return "", err
	}

	// Pages
	pagesPath := path.Join(temporaryDirectory, "pages")
	err = os.Mkdir(pagesPath, 0755)
	if err != nil {
		return "", err
	}
	pageFile, err := os.Create(path.Join(pagesPath, "index.html"))
	if err != nil {
		return "", err
	}
	pageFile.Close()
	pageFile, err = os.Create(path.Join(pagesPath, "some-page.html"))
	if err != nil {
		return "", err
	}
	pageFile.Close()
	subPagesPath := path.Join(pagesPath, "sub")
	err = os.Mkdir(subPagesPath, 0755)
	if err != nil {
		return "", err
	}
	pageFile, err = os.Create(path.Join(subPagesPath, "index.html"))
	if err != nil {
		return "", err
	}
	pageFile.Close()
	pageFile, err = os.Create(path.Join(subPagesPath, "subsubsub.html"))
	if err != nil {
		return "", err
	}
	pageFile.Close()

	return temporaryDirectory, nil
}

func TestLoad(t *testing.T) {
	dirPath, err := createExampleSite()
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	defer os.RemoveAll(dirPath)

	config, err := Load(dirPath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if config.Config.Title != "Young Fatigue" {
		t.Fatalf("Incorrect config.Config.Title: %v", config.Config.Title)
	}
	if config.Config.Description != "New Single ‘Dislocation‘ Out Now!" {
		t.Fatalf("Incorrect config.Config.Description: %v", config.Config.Description)
	}
	if config.Config.Image != "icon.png" {
		t.Fatalf("Incorrect config.Config.Image: %v", config.Config.Image)
	}
	if config.Config.RootUrl != "https://youngfatigue.com/" {
		t.Fatalf("Incorrect config.Config.RootUrl: %v", config.Config.RootUrl)
	}
	if config.Path != dirPath {
		t.Fatalf("Incorrect config.Path: %v", config.Path)
	}
	if config.Config.Data["instagram"] != "https://www.instagram.com/youngfatigue/" {
		t.Fatalf("Incorrect config.Config.Data: %v", config.Config.Data)
	}
	if len(config.Pages) != 4 {
		t.Fatalf("Incorrect config.Pages: %v", config.Pages)
	}
}

func TestLoadMissingSite(t *testing.T) {
	dirPath, err := createExampleSite()
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	os.RemoveAll(dirPath) // remove the directory so it doesn't exist

	_, err = Load(dirPath)
	if err == nil {
		t.Fatalf("Expected error but got nil")
	}
}

func TestLoadMissingConfig(t *testing.T) {
	dirPath, err := createExampleSite()
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	os.Remove(path.Join(dirPath, "config.yaml")) // remove the config file so it doesn't exist

	_, err = Load(dirPath)
	if err == nil {
		t.Fatalf("Expected error but got nil")
	}
}

func TestLoadMissingPages(t *testing.T) {
	dirPath, err := createExampleSite()
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	os.RemoveAll(path.Join(dirPath, "pages")) // remove the directory so it doesn't exist

	_, err = Load(dirPath)
	if err == nil {
		t.Fatalf("Expected error but got nil")
	}
}
