package utils

import (
	"os"
	"path"
	"reflect"
	"testing"
)

func TestLoadTemplateFiles(t *testing.T) {
	// Create a temporary hierarchy
	temporaryDirectory, err := os.MkdirTemp("", "pages")
	if err != nil {
		t.Fatalf("%v", err)
	}
	pageFile, err := os.Create(path.Join(temporaryDirectory, ".hidden.html"))
	if err != nil {
		t.Fatalf("%v", err)
	}
	pageFile.Close()
	pageFile, err = os.Create(path.Join(temporaryDirectory, "index.html"))
	if err != nil {
		t.Fatalf("%v", err)
	}
	_, err = pageFile.WriteString("index file")
	if err != nil {
		t.Fatalf("%v", err)
	}
	pageFile.Close()
	pageFile, err = os.Create(path.Join(temporaryDirectory, "some-page.markdown"))
	if err != nil {
		t.Fatalf("%v", err)
	}
	_, err = pageFile.WriteString("some _markdown_ file")
	if err != nil {
		t.Fatalf("%v", err)
	}
	pageFile.Close()
	pageFile, err = os.Create(path.Join(temporaryDirectory, "notatemplate.real"))
	if err != nil {
		t.Fatalf("%v", err)
	}
	pageFile.Close()
	subPath := path.Join(temporaryDirectory, "sub")
	err = os.Mkdir(subPath, 0755)
	if err != nil {
		t.Fatalf("%v", err)
	}
	pageFile, err = os.Create(path.Join(subPath, "index.html"))
	if err != nil {
		t.Fatalf("%v", err)
	}
	_, err = pageFile.WriteString("sub index file")
	if err != nil {
		t.Fatalf("%v", err)
	}
	pageFile.Close()
	pageFile, err = os.Create(path.Join(subPath, "subsubsub.markdown"))
	if err != nil {
		t.Fatalf("%v", err)
	}
	_, err = pageFile.WriteString("subsubsub _markdown_ file")
	if err != nil {
		t.Fatalf("%v", err)
	}
	pageFile.Close()

	result, err := LoadTemplateFiles(temporaryDirectory, "derp")
	expected := map[string]string{
		"derp/index":         "index file",
		"derp/some-page":     "<p>some <em>markdown</em> file</p>\n",
		"derp/sub/index":     "sub index file",
		"derp/sub/subsubsub": "<p>subsubsub <em>markdown</em> file</p>\n",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Received:\n%+v\nExpected:\n%+v", result, expected)
	}
}
