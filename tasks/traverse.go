package tasks

import (
	"io/fs"
	"path/filepath"
	"strconv"

	"github.com/tejastn10/halcyon/utils"
)

// FileInfo holds the basic details of a file
type FileInfo struct {
	Path      string
	Canonical string
	Info      fs.FileInfo
}

// TraverseDirectory traverses a directory recursively and returns information about all files.
func TraverseDirectory(dir string) (map[string][]FileInfo, error) {
	fileMap := make(map[string][]FileInfo)

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

			canonical := utils.GetCanonicalName(info.Name())
			key := canonical + "_" + strconv.FormatInt(info.Size(), 10) // Group by name and size

			fileMap[key] = append(fileMap[key], FileInfo{
				Path:      path,
				Canonical: canonical,
				Info:      info,
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileMap, nil
}
