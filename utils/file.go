// File utilities.
package utils

import (
	"os"
	"path"
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
