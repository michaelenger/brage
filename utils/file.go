// File utilities.
package utils

import (
	"fmt"
	"io"
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

// Copy a directory into another.
func CopyDirectory(sourceDirectory string, targetDirectory string) (int, error) {
	count := 0

	files, err := os.ReadDir(sourceDirectory)
	if err != nil {
		return count, err
	}

	targetDirectory = path.Join(targetDirectory, path.Base(sourceDirectory))

	os.MkdirAll(targetDirectory, 0755)

	for _, file := range files {
		sourcePath := path.Join(sourceDirectory, file.Name())

		if file.IsDir() {
			subcount, err := CopyDirectory(sourcePath, targetDirectory)
			if err != nil {
				return count, err
			}

			count += subcount
		} else {
			targetPath := path.Join(targetDirectory, file.Name())
			if err := CopyFile(sourcePath, targetPath); err != nil {
				return count, err
			}

			count += 1
		}
	}

	return count, nil
}

// Copy a file from a source to a destination.
// Taken from: https://opensource.com/article/18/6/copying-files-go
func CopyFile(source string, destination string) error {
	sourceFileStat, err := os.Stat(source)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", source)
	}

	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
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
