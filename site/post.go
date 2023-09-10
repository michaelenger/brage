package site

import (
	"log"
	"time"

	"brage/files"
	"github.com/cbroglie/mustache"
)

// A blog post.
type Post struct {
	Path     string
	Title    string
	Date     time.Time
	Template string
}

// Type alias for sorting posts by date.
type byPostDate []Post

func (s byPostDate) Len() int {
	return len(s)
}
func (s byPostDate) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byPostDate) Less(i, j int) bool {
	return s[i].Date.Before(s[j].Date)
}

// Make a post out the given File.
func MakePost(file files.File, pathName string) Post {
	var content string
	var metadata map[string]interface{}
	switch file.Type {
	case files.MarkdownFile:
		metadata, content = files.ParseMarkdown(file.Content)
	default:
		content = string(file.Content)
	}

	title := files.PathToTitle(file.Path)
	if val, ok := metadata["title"]; ok {
		title = val.(string)
	}

	publishedDate := time.Now()
	if val, ok := metadata["date"]; ok {
		parsedTime, err := time.Parse("2006-01-02", val.(string))
		if err != nil {
			logger := log.Default()
			logger.Printf("Unable to parse published date: %v", val)
		}

		publishedDate = parsedTime
	}

	return Post{
		pathName,
		title,
		publishedDate,
		content,
	}
}

// Create the context used when rendering the post.
func (post Post) makeContext(site Site) map[string]interface{} {
	postContext := map[string]interface{}{
		"path":     post.Path,
		"template": post.Template,
		"title":    post.Title,
		"date":     post.Date.Format("2006-01-02"),
	}

	return map[string]interface{}{
		"site": site.MakeContext(),
		"post": postContext,
		"data": site.Config.Data,
	}
}

// Render a post using a specific site config and layout file.
func (post Post) Render(site Site) (string, error) {
	context := post.makeContext(site)

	partialsProvider := &mustache.StaticProvider{site.Partials}

	return mustache.RenderInLayoutPartials(post.Template, site.Layouts[PostLayout], partialsProvider, context)
}
