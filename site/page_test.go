package site

import (
	"os"
	"path"
	"regexp"
	"testing"

	"github.com/michaelenger/brage/utils"
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
	if err := utils.WriteFile(layoutFilePath, layoutTemplate); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	templatesPath := path.Join(temporaryDirectory, "templates")
	if err := os.Mkdir(templatesPath, 0755); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	templateTemplate := `<em>This is from a template</em>`
	templateFilePath := path.Join(templatesPath, "temp.html")
	if err := utils.WriteFile(templateFilePath, templateTemplate); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	pageTemplate := `<h1>{{ .Page.Title }}</h1>
		{{ range .Data.Skills }}
			<p>{{ . }}</p>
		{{ end }}

		{{ template "temp" . }}`
	pageFilePath := path.Join(temporaryDirectory, "page.html")
	if err := utils.WriteFile(pageFilePath, pageTemplate); err != nil {
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

			<em>This is from a template</em>

		</body>`

	site := Site{
		testConfig,
		temporaryDirectory,
		[]Page{},
		map[string]string{
			"temp": templateFilePath,
		},
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
