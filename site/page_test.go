package site

import (
	"os"
	"regexp"
	"testing"
)

var testConfig = SiteConfig{
	"Test Site",
	"This is just a test.",
	"test.jpg",
	"https://example.org/",
	map[string]string{
		"/example": "https://example.org/",
	},
	DataMap{
		"Skills": []string{
			"one", "two", "three",
		},
	},
}

var whitespacePattern = regexp.MustCompile(`\s`)

func TestPageTitle(t *testing.T) {
	var tests map[string]string = map[string]string{
		"/":               "Home",
		"/about":          "About",
		"/this-is-a-test": "This Is A Test",
		"/one/two/three":  "Three",
	}

	var page Page
	var result string

	for path, expected := range tests {
		page = Page{
			path,
			"",
		}
		result = page.Title()
		if result != expected {
			t.Fatalf("Result:\n%v\nExpected:\n%v", result, expected)
		}
	}
}

func TestPageRender(t *testing.T) {
	temporaryDirectory, err := os.MkdirTemp("", "examplesite")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	page := Page{
		"/example",
		`<h1>{{ .Page.Title }}</h1>
		{{ range .Data.Skills }}
			<p>{{ . }}</p>
		{{ end }}

		{{ template "temp" . }}`,
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
		`<head>
			<title>{{ .Site.Title }}</title>
		</head>
		<body>
		{{ .Content }}
		</body>`,
		[]Page{},
		map[string]string{
			"temp": `<em>This is from a template</em>`,
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

func TestPageRenderWithMarkdown(t *testing.T) {
	temporaryDirectory, err := os.MkdirTemp("", "examplesite")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	page := Page{
		"/example",
		`{{ markdown "Now this is _podracing_!" }}`,
	}

	expected := `<p>Now this is <em>podracing</em>!</p>
		`

	site := Site{
		testConfig,
		temporaryDirectory,
		`{{ .Content }}`,
		[]Page{},
		map[string]string{},
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
