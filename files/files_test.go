package files

import (
	"os"
	"path"
	"reflect"
	"testing"
)

func TestPathToIdentifier(t *testing.T) {
	var tests map[string]string = map[string]string{
		"/":               "index",
		"/about":          "about",
		"/this is a test": "this-is-a-test",
		"/one/two/three":  "one-two-three",
	}

	var result string

	for filePath, expected := range tests {
		result = PathToIdentifier(filePath)
		if result != expected {
			t.Fatalf("Result:\n%v\nExpected:\n%v", result, expected)
		}
	}
}

func TestPathToTitle(t *testing.T) {
	var tests map[string]string = map[string]string{
		"/":               "Home",
		"/about":          "About",
		"/this-is-a-test": "This Is A Test",
		"/one/two/three":  "Three",
	}

	var result string

	for filePath, expected := range tests {
		result = PathToTitle(filePath)
		if result != expected {
			t.Fatalf("Result:\n%v\nExpected:\n%v", result, expected)
		}
	}
}

func TestReadFiles(t *testing.T) {
	test_files := map[string]string{
		".hidden.html":       "",
		"index.html":         "index file",
		"some-page.markdown": "some _markdown_ file",
		"notatemplate.real":  "",
	}
	test_sub_files := map[string]string{
		"index.htm":      "sub index file",
		"subsubsub.md":   "subsubsub _markdown_ file",
		"ignorethis.dat": "",
	}

	temporaryDirectory, err := os.MkdirTemp("", "pages")
	if err != nil {
		t.Fatalf("%v", err)
	}
	for filename, contents := range test_files {
		file, err := os.Create(path.Join(temporaryDirectory, filename))
		if err != nil {
			t.Fatalf("%v", err)
		}
		_, err = file.WriteString(contents)
		if err != nil {
			t.Fatalf("%v", err)
		}
		file.Close()
	}

	subDirectory := path.Join(temporaryDirectory, "sub")
	err = os.Mkdir(subDirectory, 0755)
	if err != nil {
		t.Fatalf("%v", err)
	}
	for filename, contents := range test_sub_files {
		file, err := os.Create(path.Join(subDirectory, filename))
		if err != nil {
			t.Fatalf("%v", err)
		}
		_, err = file.WriteString(contents)
		if err != nil {
			t.Fatalf("%v", err)
		}
		file.Close()
	}

	result, err := ReadFiles(temporaryDirectory, "derp")
	expected := map[string]File{
		"derp/index":         {HtmlFile, path.Join(temporaryDirectory, "index.html"), []byte("index file")},
		"derp/some-page":     {MarkdownFile, path.Join(temporaryDirectory, "some-page.markdown"), []byte("some _markdown_ file")},
		"derp/sub/index":     {HtmlFile, path.Join(subDirectory, "index.htm"), []byte("sub index file")},
		"derp/sub/subsubsub": {MarkdownFile, path.Join(subDirectory, "subsubsub.md"), []byte("subsubsub _markdown_ file")},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Received:\n%+v\nExpected:\n%+v", result, expected)
	}
}
