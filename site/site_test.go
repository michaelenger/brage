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
root_url: https://youngfatigue.com/
redirects:
  /example: https://example.org/

data:
  instagram: https://www.instagram.com/youngfatigue/
  cool_quotes:
    - It’s actually really good... really good.
    - "[Dislocation] is super fun!"
`

func createExampleSite(t *testing.T) string {
	// Site directory
	temporaryDirectory, err := os.MkdirTemp("", "examplesite")
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}

	// Config file
	configFile, err := os.Create(path.Join(temporaryDirectory, "config.yaml"))
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	_, err = configFile.WriteString(exampleConfig)
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	configFile.Close()

	// Layout file
	layoutFile, err := os.Create(path.Join(temporaryDirectory, "layout.html"))
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	_, err = layoutFile.WriteString("This is a layout")
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	layoutFile.Close()

	// Pages
	pagesPath := path.Join(temporaryDirectory, "pages")
	err = os.Mkdir(pagesPath, 0755)
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	pageFile, err := os.Create(path.Join(pagesPath, "index.html"))
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	pageFile.Close()
	pageFile, err = os.Create(path.Join(pagesPath, "some-page.html"))
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	pageFile.Close()
	subPagesPath := path.Join(pagesPath, "sub")
	err = os.Mkdir(subPagesPath, 0755)
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	pageFile, err = os.Create(path.Join(subPagesPath, "index.html"))
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	pageFile.Close()
	pageFile, err = os.Create(path.Join(subPagesPath, "subsubsub.html"))
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	pageFile.Close()

	// Partials
	partialsPath := path.Join(temporaryDirectory, "partials")
	err = os.Mkdir(partialsPath, 0755)
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	templateFile, err := os.Create(path.Join(partialsPath, "one.html"))
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	templateFile.Close()
	templateFile, err = os.Create(path.Join(partialsPath, "something.markdown"))
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	_, err = templateFile.WriteString("**bold text**")
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	templateFile.Close()
	subTemplatesPath := path.Join(partialsPath, "sub")
	err = os.Mkdir(subTemplatesPath, 0755)
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	templateFile, err = os.Create(path.Join(subTemplatesPath, "two.html"))
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	templateFile.Close()

	// Posts
	postsPath := path.Join(temporaryDirectory, "posts")
	err = os.Mkdir(postsPath, 0755)
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	templateFile, err = os.Create(path.Join(postsPath, "test.markdown"))
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	templateFile.Close()
	templateFile, err = os.Create(path.Join(postsPath, "something-else.markdown"))
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	templateFile.Close()

	return temporaryDirectory
}

func TestLoad(t *testing.T) {
	dirPath := createExampleSite(t)
	defer os.RemoveAll(dirPath)

	site, err := Load(dirPath)
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
	if len(site.Config.Redirects) != 1 {
		t.Fatalf("Incorrect site.Config.Redirects: %v", site.Config.Redirects)
	}
	if site.Config.Redirects["/example"] != "https://example.org/" {
		t.Fatalf("Incorrect site.Config.Redirects[\"/example\"]: %v", site.Config.Redirects["/example"])
	}
	if site.SourceDirectory != dirPath {
		t.Fatalf("Incorrect site.SourceDirectory: %v", site.SourceDirectory)
	}
	if site.Layout != "This is a layout" {
		t.Fatalf("Incorrect site.Layout: %v", site.Layout)
	}
	if site.Config.Data["instagram"] != "https://www.instagram.com/youngfatigue/" {
		t.Fatalf("Incorrect site.Config.Data[\"instagram\"]: %v", site.Config.Data["instagram"])
	}
	if site.Config.Data["cool_quotes"].([]interface{})[1] != "[Dislocation] is super fun!" {
		t.Fatalf("Incorrect site.Config.Data[\"cool_quotes\"]: %v", site.Config.Data["cool_quotes"])
	}
	if len(site.Pages) != 4 {
		t.Fatalf("Incorrect site.Pages: %v", site.Pages)
	}
	if len(site.Partials) != 3 {
		t.Fatalf("Incorrect site.Partials: %v", site.Partials)
	}
	if len(site.Posts) != 2 {
		t.Fatalf("Incorrect site.Posts: %v", site.Posts)
	}
}

func TestLoadWithoutLayout(t *testing.T) {
	dirPath := createExampleSite(t)
	defer os.RemoveAll(dirPath)
	os.Remove(path.Join(dirPath, "layout.html"))

	site, err := Load(dirPath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if site.Layout != "{{{ content }}}" {
		t.Fatalf("Incorrect site.Layout: %v", site.Layout)
	}
}

func TestLoadWithoutPartials(t *testing.T) {
	dirPath := createExampleSite(t)
	defer os.RemoveAll(dirPath)
	os.RemoveAll(path.Join(dirPath, "partials"))

	_, err := Load(dirPath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestLoadMissingSite(t *testing.T) {
	dirPath := createExampleSite(t)
	os.RemoveAll(dirPath)

	_, err := Load(dirPath)
	if err == nil {
		t.Fatalf("Expected error but got nil")
	}
}

func TestLoadMissingConfig(t *testing.T) {
	dirPath := createExampleSite(t)
	defer os.RemoveAll(dirPath)
	os.Remove(path.Join(dirPath, "config.yaml"))

	_, err := Load(dirPath)
	if err == nil {
		t.Fatalf("Expected error but got nil")
	}
}

func TestLoadMissingPages(t *testing.T) {
	dirPath := createExampleSite(t)
	defer os.RemoveAll(dirPath)
	os.RemoveAll(path.Join(dirPath, "pages"))

	_, err := Load(dirPath)
	if err == nil {
		t.Fatalf("Expected error but got nil")
	}
}
