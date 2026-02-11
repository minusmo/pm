package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hojooneum/pm/internal/fs"
	"github.com/hojooneum/pm/internal/manual"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a .pm/ directory with default runbook templates",
	RunE:  runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, args []string) error {
	root, _ := os.Getwd()
	pmPath := fs.PMPath(root)
	w := cmd.OutOrStdout()

	// Create directory structure
	for _, dir := range []string{
		filepath.Join(pmPath, "core"),
		filepath.Join(pmPath, "custom"),
	} {
		if err := fs.EnsureDir(dir); err != nil {
			return fmt.Errorf("creating directory %s: %w", dir, err)
		}
	}

	// Write template files
	createdCount := 0
	skippedCount := 0

	for _, name := range manual.CoreSectionOrder {
		tmpl, ok := manual.DefaultTemplates[name]
		if !ok {
			continue
		}
		path := filepath.Join(pmPath, "core", name+".md")
		created, err := fs.WriteFileIfNotExists(path, tmpl)
		if err != nil {
			return fmt.Errorf("writing %s: %w", name+".md", err)
		}
		if created {
			createdCount++
			fmt.Fprintf(w, "  created: core/%s.md\n", name)
		} else {
			skippedCount++
			fmt.Fprintf(w, "  exists:  core/%s.md (skipped)\n", name)
		}
	}

	fmt.Fprintln(w)
	fmt.Fprintf(w, "Initialized .pm/ — %d file(s) created, %d skipped.\n", createdCount, skippedCount)
	fmt.Fprintln(w, "Edit the files in .pm/core/ to document your project.")
	return nil
}
