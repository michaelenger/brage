package site

import (
	"bytes"
	"html/template"
	"path"
	"strings"
)

type PageDefinition struct {
	Title string
	Path  string
}

type PageTemplateData struct {
	Site SiteConfig
	Page PageDefinition
	Data DataMap
}

// Render a page.
func RenderPage(pagePath string, siteConfig SiteConfig, filePath string) (string, error) {
	pageTemplate, err := template.ParseFiles(filePath)
	if err != nil {
		return "", err
	}

	title := strings.Title(
		strings.ReplaceAll(
			strings.ReplaceAll(
				path.Base(pagePath), "_", " "),
			"-", " "))

	data := PageTemplateData{
		siteConfig,
		PageDefinition{
			title,
			filePath,
		},
		siteConfig.Data,
	}

	var buffer bytes.Buffer
	err = pageTemplate.Execute(&buffer, data)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
