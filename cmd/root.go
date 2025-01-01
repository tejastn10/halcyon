package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/tejastn10/halcyon/tasks"
)

var (
	maxSize    int64
	minSize    int64
	concurrent bool
	extensions []string
	targetDir  string
)

var rootCmd = &cobra.Command{
	Use:   "halcyon",
	Short: "Find duplicate files in directories",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create context with cancellation
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Configure traverse options
		opts := tasks.TraverseOptions{
			MaxSize:    maxSize,
			MinSize:    minSize,
			Concurrent: concurrent,
			Extensions: extensions,
		}

		// Run directory traversal
		result, err := tasks.TraverseDirectory(ctx, targetDir, opts)
		if err != nil {
			return fmt.Errorf("failed to traverse directory: %w", err)
		}

		// Print results
		fmt.Printf("Total files: %d\n", result.Stats.TotalFiles)
		fmt.Printf("Processed: %d\n", result.Stats.ProcessedFiles)
		fmt.Printf("Skipped: %d\n", result.Stats.SkippedFiles)

		// Process potential duplicates
		for _, files := range result.Files {
			if len(files) > 1 {
				if err := handleDuplicates(files); err != nil {
					return err
				}
			}
		}

		return nil
	},
}

func displayDuplicatesTable(files []tasks.FileInfo) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Index", "Path", "Size", "Modified"})
	table.SetBorder(true)

	for i, f := range files {
		size := fmt.Sprintf("%.2f MB", float64(f.Info.Size())/1024/1024)
		table.Append([]string{
			strconv.Itoa(i),
			f.Path,
			size,
			f.Modified.Format("2006-01-02 15:04:05"),
		})
	}
	table.Render()
}

func handleDuplicates(files []tasks.FileInfo) error {
	for {
		displayDuplicatesTable(files)

		prompt := promptui.Select{
			Label: "Choose action",
			Items: []string{
				"Keep all files",
				"Delete specific files",
				"Delete all duplicates",
				"Move files to backup",
				"Skip these duplicates",
				"Exit",
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			return err
		}

		switch result {
		case "Keep all files", "Skip these duplicates", "Exit":
			return nil

		case "Delete specific files":
			indices, err := promptForIndices(len(files))
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			if err := deleteFiles(files, indices); err != nil {
				fmt.Printf("Error: %v\n", err)
			}

		case "Delete all duplicates":
			if confirmed := confirmAction("Delete all duplicates?"); confirmed {
				if err := deleteAllDuplicates(files); err != nil {
					fmt.Printf("Error: %v\n", err)
				}
			} else {
				fmt.Println("Action cancelled")
			}

		case "Move files to backup":
			dir, err := promptForBackupDir()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			if err := moveFilesToBackup(files, dir); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}
	}
}

func promptForIndices(max int) ([]int, error) {
	prompt := promptui.Prompt{
		Label: "Enter indices to delete (comma-separated)",
		Validate: func(input string) error {
			indices := strings.Split(input, ",")
			for _, idx := range indices {
				num, err := strconv.Atoi(strings.TrimSpace(idx))
				if err != nil || num < 0 || num >= max {
					return fmt.Errorf("invalid index: %s", idx)
				}
			}
			return nil
		},
	}

	result, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	var indices []int
	for _, idx := range strings.Split(result, ",") {
		num, _ := strconv.Atoi(strings.TrimSpace(idx))
		indices = append(indices, num)
	}
	return indices, nil
}

func confirmAction(message string) bool {
	prompt := promptui.Prompt{
		Label: message + " (type 'yes' or 'no')",
		Templates: &promptui.PromptTemplates{
			Success: "{{ . }} ",
			Valid:   "{{ . }} ",
			Invalid: "{{ . }} ",
			Confirm: "{{ . }} (yes/no) ",
		},
		Validate: func(input string) error {
			input = strings.ToLower(strings.TrimSpace(input))
			if input != "yes" && input != "no" {
				return fmt.Errorf("please type 'yes' or 'no'")
			}
			return nil
		},
	}

	result, err := prompt.Run()
	if err != nil {
		return false
	}
	return strings.ToLower(strings.TrimSpace(result)) == "yes"
}

func deleteFiles(files []tasks.FileInfo, indices []int) error {
	for _, idx := range indices {
		if err := os.Remove(files[idx].Path); err != nil {
			return fmt.Errorf("failed to delete %s: %w", files[idx].Path, err)
		}
		fmt.Printf("Deleted: %s\n", files[idx].Path)
	}
	return nil
}

func deleteAllDuplicates(files []tasks.FileInfo) error {
	// Keep the first file, delete the rest
	for i := 1; i < len(files); i++ {
		if err := os.Remove(files[i].Path); err != nil {
			return fmt.Errorf("failed to delete %s: %w", files[i].Path, err)
		}
		fmt.Printf("Deleted: %s\n", files[i].Path)
	}
	return nil
}

func promptForBackupDir() (string, error) {
	prompt := promptui.Prompt{
		Label: "Enter backup directory path",
	}

	dir, err := prompt.Run()
	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(dir, 0750); err != nil {
		return "", fmt.Errorf("failed to create backup directory: %w", err)
	}
	return dir, nil
}

func moveFilesToBackup(files []tasks.FileInfo, backupDir string) error {
	for _, f := range files {
		newPath := filepath.Join(backupDir, filepath.Base(f.Path))
		if err := os.Rename(f.Path, newPath); err != nil {
			return fmt.Errorf("failed to move %s: %w", f.Path, err)
		}
		fmt.Printf("Moved: %s -> %s\n", f.Path, newPath)
	}
	return nil
}

func init() {
	rootCmd.Flags().StringVarP(&targetDir, "dir", "d", ".", "Directory to scan")
	rootCmd.Flags().Int64VarP(&maxSize, "max-size", "M", 0, "Maximum file size in bytes")
	rootCmd.Flags().Int64VarP(&minSize, "min-size", "m", 0, "Minimum file size in bytes")
	rootCmd.Flags().BoolVarP(&concurrent, "concurrent", "c", true, "Enable concurrent processing")
	rootCmd.Flags().StringSliceVarP(&extensions, "ext", "e", nil, "File extensions to process (e.g., .txt,.pdf)")
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
