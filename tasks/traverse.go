package tasks

import (
	"io/fs"
	"path/filepath"
)

// FileInfo holds the basic details of a file
type FileInfo struct {
	Path string
	Info fs.FileInfo
}

// TraverseDirectory traverses a directory recursively and returns information about all files.
func TraverseDirectory(dir string) ([]FileInfo, error) {
	var files []FileInfo

	// Walk through the directory
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// Return error if any issues occur during traversal
			return err
		}

		// Only process files, not directories
		if !d.IsDir() {
			info, err := d.Info()

			if err != nil {
				return err
			}

			files = append(files, FileInfo{
				Path: path,
				Info: info,
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
