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
		"skills": []string{
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

	page := Page{
		"/example",
		`<h1>{{ page.title }}</h1>
		{{ #data.skills }}
			<p>{{ . }}</p>
		{{ /data.skills }}

		{{> temp }}`,
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
		map[LayoutType]string{
			PageLayout: `<head>
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
