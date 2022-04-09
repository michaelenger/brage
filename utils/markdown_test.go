package utils

import (
	"testing"
)

func TestRenderMarkdown(t *testing.T) {
	var expected string
	var result string

	result = RenderMarkdown([]byte("just some text"))
	expected = `<p>just some text</p>
`
	if result != expected {
		t.Fatalf("Expected: '%s'\nReceived: '%s'", expected, result)
	}

	result = RenderMarkdown([]byte("just _some_ text"))
	expected = `<p>just <em>some</em> text</p>
`
	if result != expected {
		t.Fatalf("Expected: '%s'\nReceived: '%s'", expected, result)
	}

	result = RenderMarkdown([]byte("just _some_ [text](https://example.org)"))
	expected = `<p>just <em>some</em> <a href="https://example.org">text</a></p>
`
	if result != expected {
		t.Fatalf("Expected: '%s'\nReceived: '%s'", expected, result)
	}
}
