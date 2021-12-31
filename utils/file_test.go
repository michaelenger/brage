package utils

import (
	"os"
	"path"
	"reflect"
	"testing"
)

func TestListHTMLFiles(t *testing.T) {
	// Create a temporary hierarchy
	temporaryDirectory, err := os.MkdirTemp("", "pages")
	if err != nil {
		t.Fatalf("%v", err)
	}
	pageFile, err := os.Create(path.Join(temporaryDirectory, "index.html"))
	if err != nil {
		t.Fatalf("%v", err)
	}
	pageFile.Close()
	pageFile, err = os.Create(path.Join(temporaryDirectory, "some-page.html"))
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
	pageFile.Close()
	pageFile, err = os.Create(path.Join(subPath, "subsubsub.html"))
	if err != nil {
		t.Fatalf("%v", err)
	}
	pageFile.Close()

	result, err := ListHTMLFiles(temporaryDirectory, "derp")
	expected := map[string]string{
		"derp/index":         path.Join(temporaryDirectory, "index.html"),
		"derp/some-page":     path.Join(temporaryDirectory, "some-page.html"),
		"derp/sub/index":     path.Join(subPath, "index.html"),
		"derp/sub/subsubsub": path.Join(subPath, "subsubsub.html"),
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Received:\n%+v\nExpected:\n%+v", result, expected)
	}
}
