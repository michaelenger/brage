package site

import (
	"bytes"
	"path"
	"strings"
	"text/template"
)

type Page struct {
	Path         string
	TemplateFile string
}

// Get the title of a page.
func (page Page) Title() string {
	return strings.Title(
		strings.ReplaceAll(
			strings.ReplaceAll(
				path.Base(page.Path), "_", " "),
			"-", " "))
}

// Render a page using a specific site config and layout file.
func (page Page) Render(site SiteDescription) (string, error) {
	layoutFilePath := path.Join(site.SourceDirectory, "layout.html")
	layoutTemplate, err := template.ParseFiles(layoutFilePath)
	if err != nil {
		return "", err
	}
	pageTemplate, err := template.ParseFiles(page.TemplateFile)
	if err != nil {
		return "", err
	}

	pageData := struct {
		Site SiteConfig
		Page Page
		Data DataMap
	}{
		site.Config,
		page,
		site.Config.Data,
	}

	var buffer bytes.Buffer
	err = pageTemplate.Execute(&buffer, pageData)
	if err != nil {
		return "", err
	}

	layoutData := struct {
		Site    SiteConfig
		Page    Page
		Data    DataMap
		Content string
	}{
		site.Config,
		page,
		site.Config.Data,
		buffer.String(),
	}

	buffer.Reset()
	err = layoutTemplate.Execute(&buffer, layoutData)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
