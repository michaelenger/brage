package site

import (
	"os"
	"path"
	"regexp"
	"testing"
)

var testConfig = SiteConfig{
	"Test Site",
	"This is just a test.",
	"test.jpg",
	"https://example.org/",
	DataMap{
		"Skills": []string{
			"one", "two", "three",
		},
	},
}

var whitespacePattern = regexp.MustCompile(`\s`)

func makeTempFile(filePath string, contents string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(contents)
	if err != nil {
		return err
	}

	return nil
}

func TestPageRender(t *testing.T) {
	temporaryDirectory, err := os.MkdirTemp("", "examplesite")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	layoutTemplate := `<head>
			<title>{{ .Site.Title }}</title>
		</head>
		<body>
		{{ .Content }}
		</body>`
	layoutFilePath := path.Join(temporaryDirectory, "layout.html")
	if err := makeTempFile(layoutFilePath, layoutTemplate); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	pageTemplate := `<h1>{{ .Page.Title }}</h1>
		{{ range .Data.Skills }}
			<p>{{ . }}</p>
		{{ end }}`
	pageFilePath := path.Join(temporaryDirectory, "page.html")
	if err := makeTempFile(pageFilePath, pageTemplate); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	page := Page{
		"/example",
		pageFilePath,
	}

	expected := `<head>
			<title>Test Site</title>
		</head>
		<body>
		<h1>Example</h1>

			<p>one</p>

			<p>two</p>

			<p>three</p>

		</body>`

	site := SiteDescription{
		testConfig,
		temporaryDirectory,
		[]Page{},
	}

	result, err := page.Render(site)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	result = whitespacePattern.ReplaceAllString(result, "")
	expected = whitespacePattern.ReplaceAllString(expected, "")

	if result != expected {
		t.Fatalf("Result:\n%v\nExpected:\n%v", result, expected)
	}
}
