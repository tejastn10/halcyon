package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

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

		files, err := tasks.TraverseDirectory(dir)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		// Creating a tab writer for formatted output
		writer := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
		defer writer.Flush()

		// Print the header
		fmt.Fprintln(writer, "Path\tName\tSize\tMode\tModTime")

		// Print file information in a structured format
		for _, file := range files {
			fmt.Fprintf(writer, "%s\t%s\t%d bytes\t%s\t%v\n",
				file.Path,
				file.Info.Name(),
				file.Info.Size(),
				file.Info.Mode(),
				file.Info.ModTime().Format(time.RFC3339),
			)
		}
	},
}

func init() {
	// Register the --path flag here, but it's optional and defaults to current directory
	rootCmd.Flags().String("path", ".", "Directory path to traverse")
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
