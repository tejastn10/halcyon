package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/tejastn10/halcyon/tasks"
)

var rootCmd = &cobra.Command{
	Use:   "halcyon",
	Short: "A tool to find and manage duplicate files",
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("path")
		if dir == "" {
			dir = "."
			fmt.Println("No path provided, defaulting to the current directory.")
		}

		fileGroups, err := tasks.TraverseDirectory(dir)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		// Creating a tab writer for formatted output
		writer := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
		defer writer.Flush()

		// Print the header
		fmt.Fprintln(writer, "Group\tPath\tSize\tName\tCanonical Name")

		// Print file information in a structured format
		for groupKey, files := range fileGroups {
			if len(files) > 1 { // Only show duplicates
				for _, file := range files {
					fmt.Fprintf(writer, "%s\t%s\t%d bytes\t%s\t%s\n",
						groupKey,
						file.Path,
						file.Info.Size(),
						file.Info.Name(),
						file.Canonical,
					)
				}
			}
		}
	},
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
