package site

import (
	"path"
	"strings"

	"github.com/cbroglie/mustache"
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

// Create the context used when rendering a page.
func (page Page) makeContext(site Site) map[string]interface{} {
	siteContext := map[string]interface{}{
		"title":       site.Config.Title,
		"description": site.Config.Description,
		"image":       site.Config.Image,
		"root_url":    site.Config.RootUrl,
		"redirects":   site.Config.Redirects,
	}

	pageContext := map[string]string{
		"path":     page.Path,
		"template": page.Template,
		"title":    page.Title(),
	}

	return map[string]interface{}{
		"site": siteContext,
		"page": pageContext,
		"data": site.Config.Data,
	}
}

// Render a page using a specific site config and layout file.
func (page Page) Render(site Site) (string, error) {
	context := page.makeContext(site)

	partialsProvider := &mustache.StaticProvider{site.Partials}

	return mustache.RenderInLayoutPartials(page.Template, site.Layout, partialsProvider, context)
}
