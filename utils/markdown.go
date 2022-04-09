package utils

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

// Render markdown to HTML.
func RenderMarkdown(text []byte) string {
	markdownParser := parser.NewWithExtensions(parser.CommonExtensions)
	htmlText := markdown.ToHTML(text, markdownParser, nil)

	return string(htmlText)
}
