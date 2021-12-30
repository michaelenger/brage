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

func TestLoadSite(t *testing.T) {
	dirPath, err := createExampleSite()
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	defer os.RemoveAll(dirPath)

	site, err := LoadSite(dirPath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if site.Config.Title != "Young Fatigue" {
		t.Fatalf("Incorrect site.Config.Title: %v", site.Config.Title)
	}
	if site.Config.Description != "New Single ‘Dislocation‘ Out Now!" {
		t.Fatalf("Incorrect site.Config.Description: %v", site.Config.Description)
	}
	if site.Config.Image != "icon.png" {
		t.Fatalf("Incorrect site.Config.Image: %v", site.Config.Image)
	}
	if site.Config.RootUrl != "https://youngfatigue.com/" {
		t.Fatalf("Incorrect site.Config.RootUrl: %v", site.Config.RootUrl)
	}
	if site.SourceDirectory != dirPath {
		t.Fatalf("Incorrect site.SourceDirectory: %v", site.SourceDirectory)
	}
	if site.Config.Data["instagram"] != "https://www.instagram.com/youngfatigue/" {
		t.Fatalf("Incorrect site.Config.Data: %v", site.Config.Data)
	}
	if len(site.Pages) != 4 {
		t.Fatalf("Incorrect site.Pages: %v", site.Pages)
	}
}

func TestLoadSiteMissingSite(t *testing.T) {
	dirPath, err := createExampleSite()
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	os.RemoveAll(dirPath) // remove the directory so it doesn't exist

	_, err = LoadSite(dirPath)
	if err == nil {
		t.Fatalf("Expected error but got nil")
	}
}

func TestLoadSiteMissingConfig(t *testing.T) {
	dirPath, err := createExampleSite()
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	os.Remove(path.Join(dirPath, "config.yaml")) // remove the config file so it doesn't exist

	_, err = LoadSite(dirPath)
	if err == nil {
		t.Fatalf("Expected error but got nil")
	}
}

func TestLoadSiteMissingPages(t *testing.T) {
	dirPath, err := createExampleSite()
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	os.RemoveAll(path.Join(dirPath, "pages")) // remove the directory so it doesn't exist

	_, err = LoadSite(dirPath)
	if err == nil {
		t.Fatalf("Expected error but got nil")
	}
}
