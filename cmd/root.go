package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tejastn10/halcyon/tasks"
)

var rootCmd = &cobra.Command{
	Use:   "halcyon",
	Short: "A tool to find and manage duplicate files",
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("path")
		if dir == "" {
			fmt.Println("Please provide a directory path using the --path flag.")
			return
		}

		files, err := tasks.TraverseDirectory(dir)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		for _, file := range files {
			fmt.Printf("Path: %s, Size: %d bytes\n", file.Path, file.Size)
		}
	},
}

func init() {
	// Register the --path flag here so it's recognized
	rootCmd.Flags().String("path", ".", "Directory path to traverse") // Default is current directory
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
