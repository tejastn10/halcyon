package tasks

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// createTestDirectory creates a temporary directory structure for testing
func createTestDirectory(t *testing.T) string {
	t.Helper()

	tempDir, err := os.MkdirTemp("", "traverse-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	// Create test files with different sizes and extensions
	files := map[string]int{
		"file1.txt":        100,
		"file2.txt":        200,
		"document.pdf":     300,
		"image.jpg":        400,
		"file1 copy.txt":   100,
		"file1_copy_1.txt": 100,
	}

	for name, size := range files {
		path := filepath.Join(tempDir, name)
		data := make([]byte, size)
		if err := os.WriteFile(path, data, 0644); err != nil {
			t.Fatalf("Failed to create test file %s: %v", name, err)
		}
	}

	return tempDir
}

func TestTraverseDirectory(t *testing.T) {
	tempDir := createTestDirectory(t)
	defer os.RemoveAll(tempDir)

	tests := []struct {
		name    string
		opts    TraverseOptions
		wantErr bool
		check   func(*testing.T, *TraverseResult)
	}{
		{
			name:    "Empty directory path",
			opts:    TraverseOptions{},
			wantErr: true,
		},
		{
			name: "Basic traversal",
			opts: TraverseOptions{},
			check: func(t *testing.T, r *TraverseResult) {
				if r.Stats.TotalFiles != 6 {
					t.Errorf("Expected 6 total files, got %d", r.Stats.TotalFiles)
				}
			},
		},
		{
			name: "Extension filtering",
			opts: TraverseOptions{
				Extensions: []string{".txt"},
			},
			check: func(t *testing.T, r *TraverseResult) {
				if r.Stats.ProcessedFiles != 4 {
					t.Errorf("Expected 4 processed txt files, got %d", r.Stats.ProcessedFiles)
				}
			},
		},
		{
			name: "Size filtering",
			opts: TraverseOptions{
				MinSize: 200,
				MaxSize: 350,
			},
			check: func(t *testing.T, r *TraverseResult) {
				if r.Stats.ProcessedFiles != 2 {
					t.Errorf("Expected 2 files within size range, got %d", r.Stats.ProcessedFiles)
				}
			},
		},
		{
			name: "Concurrent processing",
			opts: TraverseOptions{
				Concurrent: true,
			},
			check: func(t *testing.T, r *TraverseResult) {
				if r.Stats.TotalFiles != 6 {
					t.Errorf("Expected 6 total files, got %d", r.Stats.TotalFiles)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			var dir string
			if tt.name != "Empty directory path" {
				dir = tempDir
			}

			result, err := TraverseDirectory(ctx, dir, tt.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("TraverseDirectory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.check != nil {
				tt.check(t, result)
			}
		})
	}
}

func TestTraverseDirectoryContext(t *testing.T) {
	tempDir := createTestDirectory(t)
	defer os.RemoveAll(tempDir)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
	defer cancel()

	time.Sleep(2 * time.Millisecond) // Ensure context is cancelled

	_, err := TraverseDirectory(ctx, tempDir, TraverseOptions{})
	if err == nil {
		t.Error("Expected error due to context cancellation, got nil")
	}
}

func TestTraverseDirectoryInvalidPath(t *testing.T) {
	_, err := TraverseDirectory(context.Background(), "nonexistent/path", TraverseOptions{})
	if err == nil {
		t.Error("Expected error for invalid path, got nil")
	}
}
