package site

import (
	"bytes"
	"path"
	"strings"
	"text/template"
)

type Page struct {
	Path     string
	Template string
}

// Get the title of a page.
func (page Page) Title() string {
	return strings.Title(
		strings.ReplaceAll(
			strings.ReplaceAll(
				path.Base(page.Path), "_", " "),
			"-", " "))
}

// Add the extra templates to the main template.
func addTemplates(mainTemplate *template.Template, templateFiles map[string]string) error {
	for name, content := range templateFiles {
		subTemplate := mainTemplate.New(name)
		_, err := subTemplate.Parse(content)
		if err != nil {
			return err
		}
	}

	return nil
}

// Render a page using a specific site config and layout file.
func (page Page) Render(site Site) (string, error) {
	layoutTemplate := template.New("layout")
	layoutTemplate, err := layoutTemplate.Parse(site.Layout)
	if err != nil {
		return "", err
	}

	err = addTemplates(layoutTemplate, site.Templates)
	if err != nil {
		return "", err
	}

	pageTemplate := template.New("page")
	pageTemplate, err = pageTemplate.Parse(page.Template)
	if err != nil {
		return "", err
	}

	err = addTemplates(pageTemplate, site.Templates)
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
