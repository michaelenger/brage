package files

import (
	"reflect"
	"testing"
)

func TestParseMarkdown(t *testing.T) {
	test := []byte(`---
title: Test
test: true
number: 123
---

This is just a test`)

	meta, html := ParseMarkdown(test)
	expectedMeta := map[string]interface{}{
		"title":  "Test",
		"test":   true,
		"number": 123,
	}
	expectedHtml := `<p>This is just a test</p>
`
	if html != expectedHtml {
		t.Fatalf("Expected: '%s'\nReceived: '%s'", expectedHtml, html)
	}

	if !reflect.DeepEqual(meta, expectedMeta) {
		t.Fatalf("Received:\n%+v\nExpected:\n%+v", meta, expectedMeta)
	}
}

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

	result = RenderMarkdown([]byte("just ~~some~~ text"))
	expected = `<p>just <del>some</del> text</p>
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
