package tasks

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"
	"sync"
	"time"

	"github.com/tejastn10/halcyon/utils"
)

// FileInfo holds detailed information about a file
type FileInfo struct {
	Path      string
	Canonical string
	Info      fs.FileInfo
	Hash      string    // For content-based comparison
	Modified  time.Time // Last modified time
}

// TraverseOptions configures the directory traversal
type TraverseOptions struct {
	IgnorePatterns []string // Patterns to ignore
	MaxSize        int64    // Maximum file size to process
	MinSize        int64    // Minimum file size to process
	Concurrent     bool     // Enable concurrent processing
	Extensions     []string // File extensions to process
}

// TraverseResult contains the results and any errors
type TraverseResult struct {
	Files map[string][]FileInfo
	Stats struct {
		TotalFiles     int64
		ProcessedFiles int64
		SkippedFiles   int64
	}
}

// TraverseDirectory traverses a directory with improved handling and options
func TraverseDirectory(ctx context.Context, dir string, opts TraverseOptions) (*TraverseResult, error) {
	if dir == "" {
		return nil, fmt.Errorf("directory path cannot be empty")
	}

	result := &TraverseResult{
		Files: make(map[string][]FileInfo),
	}

	var mu sync.Mutex
	sem := make(chan struct{}, 10) // Limit concurrent operations

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to access path %s: %w", path, err)
		}

		// Check for context cancellation
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if d.IsDir() {
			return nil
		}

		// Filter by extension if specified
		if len(opts.Extensions) > 0 {
			ext := filepath.Ext(path)
			if !utils.Contains(opts.Extensions, ext) {
				result.Stats.SkippedFiles++
				return nil
			}
		}

		info, err := d.Info()
		if err != nil {
			return fmt.Errorf("failed to get file info for %s: %w", path, err)
		}

		// Size filters
		if (opts.MaxSize > 0 && info.Size() > opts.MaxSize) ||
			(opts.MinSize > 0 && info.Size() < opts.MinSize) {
			result.Stats.SkippedFiles++
			return nil
		}

		if opts.Concurrent {
			sem <- struct{}{} // Acquire semaphore
			go func() {
				defer func() { <-sem }() // Release semaphore
				fileInfo := processFile(path, info)
				mu.Lock()
				key := fmt.Sprintf("%s_%d", fileInfo.Canonical, info.Size())
				result.Files[key] = append(result.Files[key], fileInfo)
				result.Stats.ProcessedFiles++
				mu.Unlock()
			}()
		} else {
			fileInfo := processFile(path, info)
			key := fmt.Sprintf("%s_%d", fileInfo.Canonical, info.Size())
			result.Files[key] = append(result.Files[key], fileInfo)
			result.Stats.ProcessedFiles++
		}

		result.Stats.TotalFiles++
		return nil
	})

	// Wait for all goroutines to finish if concurrent
	if opts.Concurrent {
		for i := 0; i < cap(sem); i++ {
			sem <- struct{}{}
		}
	}

	if err != nil {
		return nil, fmt.Errorf("traverse error: %w", err)
	}

	return result, nil
}

func processFile(path string, info fs.FileInfo) FileInfo {
	return FileInfo{
		Path:      path,
		Canonical: utils.GetCanonicalName(info.Name()),
		Info:      info,
		Modified:  info.ModTime(),
	}
}
