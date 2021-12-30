package site

import (
	"os"
	"path"
	"testing"
)

func makeTempFile(filename string, contents string) (string, error) {
	filepath := path.Join(os.TempDir(), filename)

	file, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.WriteString(contents)
	if err != nil {
		return "", err
	}

	return filepath, nil
}

func TestRenderPage(t *testing.T) {
	var examplePage = `
<h1>{{ .Site.Title }}</h1>
<h2>{{ .Page.Title }}</h1>
{{ range .Data.Skills }}
	<p>{{ . }}</p>
{{ end }}
`
	exampleConfig := SiteConfig{
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
	var expected = `
<h1>Test Site</h1>
<h2>Example</h1>

	<p>one</p>

	<p>two</p>

	<p>three</p>

`

	filepath, err := makeTempFile("example.html", examplePage)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	result, err := RenderPage("/example", exampleConfig, filepath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Fatalf("Result:\n%v\nExpected:\n%v", result, expected)
	}
}
