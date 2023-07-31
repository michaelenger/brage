package files

import (
	"os"
	"path"
	"path/filepath"
)

// A file type.
type FileType uint8

const (
	UnknownFile FileType = iota
	HtmlFile
	MarkdownFile
)

// A file which was read from the file system.
type File struct {
	Type    FileType
	Path    string
	Content []byte
}

// Recursively read files and returns a map of their path to their content, relative to the directory path.
func ReadFiles(directoryPath string, pathPrefix string) (map[string]File, error) {
	pages := map[string]File{}

	files, err := os.ReadDir(directoryPath)
	if err != nil {
		return pages, err
	}

	for _, file := range files {
		filename := file.Name()
		if filename[0] == '.' {
			continue
		}

		fullPath := path.Join(directoryPath, filename)

		if file.IsDir() {
			subpages, err := ReadFiles(fullPath, path.Join(pathPrefix, filename))
			if err != nil {
				return pages, err
			}
			for page, file := range subpages {
				pages[page] = file
			}

			continue
		}

		fileExtension := filepath.Ext(filename)
		filetype := UnknownFile

		switch fileExtension[1:] {
		case "html", "htm":
			filetype = HtmlFile
		case "markdown", "md":
			filetype = MarkdownFile
		default:
			continue
		}

		fileContents, err := os.ReadFile(fullPath)
		if err != nil {
			return pages, err
		}

		name := path.Join(pathPrefix, filename[:len(filename)-len(fileExtension)])
		pages[name] = File{
			filetype,
			fullPath,
			fileContents,
		}
	}

	return pages, nil
}
