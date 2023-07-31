package files

import (
	"fmt"
	"io"
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

// Load the contents of markdown file.
func readMarkdownFile(filePath string) (string, error) {
	contents, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return RenderMarkdown(contents), nil
}

// Convert a relative path to an absolute path, relative to the current
// working directory.
func AbsolutePath(relativePath string) string {
	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		return relativePath // fail silently
	}

	return absolutePath
}

// Copy a directory into another.
func CopyDirectory(sourceDirectory string, targetDirectory string) (int, error) {
	count := 0

	files, err := os.ReadDir(sourceDirectory)
	if err != nil {
		return count, err
	}

	targetDirectory = path.Join(targetDirectory, path.Base(sourceDirectory))

	err = os.MkdirAll(targetDirectory, 0755)
	if err != nil {
		return count, err
	}

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

// Write the contents of the string to a file.
func WriteFile(targetFilePath string, contents string) error {
	targetDirectory := path.Dir(targetFilePath)
	if _, err := os.Stat(targetDirectory); os.IsNotExist(err) {
		if err = os.MkdirAll(targetDirectory, 0755); err != nil {
			return err
		}
	}

	file, err := os.Create(targetFilePath)
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
