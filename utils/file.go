// File utilities.
package utils

import (
	"os"
	"path"
	"regexp"
)

// Convert a relative path to an absolute path, relative to the current
// working directory.
func AbsolutePath(relativePath string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		return relativePath // fail silently
	}

	return path.Join(currentDir, relativePath)
}

// Recursively build a list of pages.
func ListPages(dirPath string, prefixPath string) (map[string]string, error) {
	pages := map[string]string{}
	htmlFilePattern := regexp.MustCompile(`^(.+?)\.html$`)

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return pages, err
	}

	for _, file := range files {
		filename := file.Name()
		filepath := path.Join(dirPath, filename)

		if file.IsDir() {
			subpages, err := ListPages(filepath, path.Join(prefixPath, filename))
			if err != nil {
				return pages, err
			}
			for page, file := range subpages {
				pages[page] = file
			}
		}

		filenameMatch := htmlFilePattern.FindStringSubmatch(filename)
		if len(filenameMatch) == 0 {
			continue
		}

		page := filenameMatch[1]
		if page == "index" {
			page = ""
		}

		pages[path.Join(prefixPath, page)] = filepath
	}

	return pages, nil
}

// Write the contents of the string to a file.
func WriteFile(filePath string, contents string) error {
	targetDirectory := path.Dir(filePath)
	if _, err := os.Stat(targetDirectory); os.IsNotExist(err) {
		if err = os.MkdirAll(targetDirectory, 0755); err != nil {
			return err
		}
	}

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
