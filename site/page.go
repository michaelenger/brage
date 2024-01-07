package site

import (
	"github.com/cbroglie/mustache"
	"github.com/michaelenger/brage/files"
)

type Page struct {
	Path     string
	Template string
}

// Create the context used when rendering a page.
func (page Page) makeContext(site Site) map[string]interface{} {
	pageContext := map[string]string{
		"path":     page.Path,
		"template": page.Template,
		"title":    files.PathToTitle(page.Path),
	}

	return map[string]interface{}{
		"site": site.MakeContext(),
		"page": pageContext,
		"data": site.Config.Data,
	}
}

// Render a page using a specific site config and layout file.
func (page Page) Render(site Site) (string, error) {
	context := page.makeContext(site)

	partialsProvider := &mustache.StaticProvider{site.Partials}

	return mustache.RenderInLayoutPartials(page.Template, site.Layouts[PageLayout], partialsProvider, context)
}
