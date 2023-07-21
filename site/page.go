package site

import (
	"bytes"
	"path"
	"strings"
	"text/template"

	"brage/utils"
)

type Page struct {
	Path     string
	Template string
}

// Get the title of a page.
func (page Page) Title() string {
	if page.Path == "/" {
		return "Home"
	}

	return strings.Title(
		strings.ReplaceAll(
			strings.ReplaceAll(
				path.Base(page.Path), "_", " "),
			"-", " "))
}

// Add functions and partial templates to the main template.
func addPartials(mainTemplate *template.Template, partials map[string]string) error {
	mainTemplate.Funcs(template.FuncMap{
		"markdown": func(text string) string {
			return utils.RenderMarkdown([]byte(text))
		},
	})

	for name, content := range partials {
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
	err := addPartials(layoutTemplate, site.Partials)
	if err != nil {
		return "", err
	}
	layoutTemplate, err = layoutTemplate.Parse(site.Layout)
	if err != nil {
		return "", err
	}

	pageTemplate := template.New("page")
	err = addPartials(pageTemplate, site.Partials)
	if err != nil {
		return "", err
	}
	pageTemplate, err = pageTemplate.Parse(page.Template)
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
