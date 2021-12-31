package site

import (
	"bytes"
	"os"
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

// Load the sub-templates.
func loadSubTemplates(mainTemplate *template.Template, templateFiles map[string]string) error {
	for name, filePath := range templateFiles {
		subTemplate := mainTemplate.New(name)
		contents, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}
		_, err = subTemplate.Parse(string(contents))
		if err != nil {
			return err
		}
	}

	return nil
}

// Render a page using a specific site config and layout file.
func (page Page) Render(site Site) (string, error) {
	layoutFilePath := path.Join(site.SourceDirectory, "layout.html")
	layoutTemplate, err := template.ParseFiles(layoutFilePath)
	if err != nil {
		return "", err
	}

	err = loadSubTemplates(layoutTemplate, site.Templates)
	if err != nil {
		return "", err
	}

	pageTemplate, err := template.ParseFiles(page.TemplateFile)
	if err != nil {
		return "", err
	}

	err = loadSubTemplates(pageTemplate, site.Templates)
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
