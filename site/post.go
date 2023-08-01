package site

import (
	"log"
	"path"
	"regexp"
	"strings"
	"time"

	"brage/files"
)

// A blog post.
type Post struct {
	Path          string
	Title         string
	Tags          []string
	PublishedDate time.Time
	Content       string
}

// Make a post out the given File.
func MakePost(file files.File, pathPrefix string) Post {
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

	tags := []string{}
	if val, ok := metadata["tags"]; ok {
		for _, tag := range strings.Split(val.(string), " ") {
			tags = append(tags, strings.ToLower(tag))
		}
	}

	publishedDate := time.Now()
	if val, ok := metadata["published_date"]; ok {
		parsedTime, err := time.Parse("2006-01-02", val.(string))
		if err != nil {
			logger := log.Default()
			logger.Printf("Unable to parse published date: %v", val)
		}

		publishedDate = parsedTime
	}

	var nonWordCharacters = regexp.MustCompile(`[^\w\s]`)
	postPath := strings.ToLower(strings.ReplaceAll(
		nonWordCharacters.ReplaceAllString(title, ""),
		" ", "-"))

	return Post{
		path.Join(pathPrefix, postPath),
		title,
		tags,
		publishedDate,
		content,
	}
}
