package site

import (
	"os"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/michaelenger/brage/files"
)

func TestMakePost(t *testing.T) {
	file := files.File{
		files.MarkdownFile,
		"/tmp/test.md",
		[]byte(`---
title: Testing!
description: I am described.
image: foo.png
date: 2020-10-01
---

This is just a test.`),
	}
	expectedTime, _ := time.Parse("2006-01-02", "2020-10-01")
	expected := Post{
		"/blog/test",
		"Testing!",
		"I am described.",
		"foo.png",
		expectedTime,
		"<p>This is just a test.</p>\n",
	}

	result := MakePost(file, "/blog/test")

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Received:\n%+v\nExpected:\n%+v", result, expected)
	}
}

func TestMakePostWithDateTime(t *testing.T) {
	file := files.File{
		files.MarkdownFile,
		"/tmp/test.md",
		[]byte(`---
title: Testing!
description: I am described.
image: foo.png
date: 2020-10-01 12:13:14
---

This is just a test.`),
	}
	expectedTime, _ := time.Parse(time.DateTime, "2020-10-01 12:13:14")
	expected := Post{
		"/blog/test",
		"Testing!",
		"I am described.",
		"foo.png",
		expectedTime,
		"<p>This is just a test.</p>\n",
	}

	result := MakePost(file, "/blog/test")

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Received:\n%+v\nExpected:\n%+v", result, expected)
	}
}

func TestMakePostDefaultMetadata(t *testing.T) {
	file := files.File{
		files.MarkdownFile,
		"/tmp/some-test.md",
		[]byte("This is a test"),
	}
	expected := Post{
		"/blog/some-test",
		"Some Test",
		"",
		"",
		time.Now(),
		"<p>This is a test</p>\n",
	}

	result := MakePost(file, "/blog/some-test")

	result.Date = expected.Date // we'll be milliseconds off, so no point in checking

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Received:\n%+v\nExpected:\n%+v", result, expected)
	}
}

func TestMakePostHtmlFile(t *testing.T) {
	file := files.File{
		files.HtmlFile,
		"/tmp/another-test.html",
		[]byte("This is a test"),
	}
	expected := Post{
		"/another-test",
		"Another Test",
		"",
		"",
		time.Now(),
		"This is a test",
	}

	result := MakePost(file, "/another-test")

	result.Date = expected.Date // we'll be milliseconds off, so no point in checking

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Received:\n%+v\nExpected:\n%+v", result, expected)
	}
}

func TestPostRender(t *testing.T) {
	var whitespacePattern = regexp.MustCompile(`\s`)

	temporaryDirectory, err := os.MkdirTemp("", "examplesite")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	date, _ := time.Parse(time.DateOnly, "2010-09-08")

	post := Post{
		"/example",
		"This is a post",
		"Just a description.",
		"",
		date,
		`<h1>{{ post.title }}</h1>
		<h2>{{ post.date }}</h2>
		{{ #data.skills }}
			<p>{{ . }}</p>
		{{ /data.skills }}

		{{> temp }}`,
	}

	expected := `<head>
			<title>Test Site</title>
		</head>
		<body>
		<h1>This is a post</h1>
		<h2>2010-09-08</h2>

			<p>one</p>

			<p>two</p>

			<p>three</p>

			<em>This is from a template</em>

		</body>`

	site := Site{
		testConfig,
		temporaryDirectory,
		map[LayoutType]string{
			PostLayout: `<head>
				<title>{{ site.title }}</title>
			</head>
			<body>
				{{{ content }}}
			</body>`,
		},
		[]Page{},
		map[string]string{
			"temp": `<em>This is from a template</em>`,
		},
		[]Post{},
	}

	result, err := post.Render(site)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	result = whitespacePattern.ReplaceAllString(result, "")
	expected = whitespacePattern.ReplaceAllString(expected, "")

	if result != expected {
		t.Fatalf("Result:\n%v\nExpected:\n%v", result, expected)
	}
}

func TestPostRenderTemplate(t *testing.T) {
	var whitespacePattern = regexp.MustCompile(`\s`)

	temporaryDirectory, err := os.MkdirTemp("", "examplesite")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	date, _ := time.Parse(time.DateOnly, "2010-09-08")

	post := Post{
		"/example",
		"This is a post",
		"Just a description.",
		"",
		date,
		`<h1>{{ post.title }}</h1>
		<h2>{{ post.date }}</h2>
		{{ #data.skills }}
			<p>{{ . }}</p>
		{{ /data.skills }}

		{{> temp }}`,
	}

	expected := `<h1>This is a post</h1>
		<h2>2010-09-08</h2>

			<p>one</p>

			<p>two</p>

			<p>three</p>

			<em>This is from a template</em>`

	site := Site{
		testConfig,
		temporaryDirectory,
		map[LayoutType]string{
			PostLayout: `<head>
				<title>{{ site.title }}</title>
			</head>
			<body>
				{{{ content }}}
			</body>`,
		},
		[]Page{},
		map[string]string{
			"temp": `<em>This is from a template</em>`,
		},
		[]Post{},
	}

	result, err := post.RenderTemplate(site)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	result = whitespacePattern.ReplaceAllString(result, "")
	expected = whitespacePattern.ReplaceAllString(expected, "")

	if result != expected {
		t.Fatalf("Result:\n%v\nExpected:\n%v", result, expected)
	}
}
