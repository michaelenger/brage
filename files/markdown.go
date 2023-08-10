package files

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

// Parse markdown, rendering it to HTML and returning the metadata as a map.
func ParseMarkdown(text []byte) (map[string]interface{}, string) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
			extension.Strikethrough,
		),
	)

	var buf bytes.Buffer
	context := parser.NewContext()
	if err := markdown.Convert(text, &buf, parser.WithContext(context)); err != nil {
		panic(err)
	}

	return meta.Get(context), buf.String()
}

// Render markdown to HTML.
func RenderMarkdown(text []byte) string {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			extension.Strikethrough,
		),
	)

	var buf bytes.Buffer
	context := parser.NewContext()
	if err := markdown.Convert(text, &buf, parser.WithContext(context)); err != nil {
		panic(err)
	}

	return buf.String()
}
