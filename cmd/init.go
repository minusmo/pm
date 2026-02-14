package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hojooneum/pm/internal/cli"
	"github.com/hojooneum/pm/internal/fs"
	"github.com/hojooneum/pm/internal/manual"
	"github.com/spf13/cobra"
)

var (
	templateFlag      string
	listTemplatesFlag bool
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a .pm/ directory with runbook templates",
	Long:  "Initialize a .pm/ directory with runbook templates.\nUse --template to select a preset or provide a JSON template file.\nUse --list-templates to see available presets.",
	RunE:  runInit,
}

func init() {
	initCmd.Flags().StringVar(&templateFlag, "template", "", "template preset name or path to JSON template file")
	initCmd.Flags().BoolVar(&listTemplatesFlag, "list-templates", false, "list available template presets")
	rootCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, args []string) error {
	w := cmd.OutOrStdout()

	if listTemplatesFlag {
		cli.PrintTemplateList(w, manual.ListPresets())
		return nil
	}

	tmpl, err := manual.ResolveTemplate(templateFlag)
	if err != nil {
		return err
	}

	root, _ := os.Getwd()
	pmPath := fs.PMPath(root)

	// Collect unique groups and ensure directories
	seen := make(map[string]bool)
	for _, s := range tmpl.Sections {
		if !seen[s.Group] {
			seen[s.Group] = true
			if err := fs.EnsureDir(filepath.Join(pmPath, s.Group)); err != nil {
				return fmt.Errorf("creating directory %s: %w", s.Group, err)
			}
		}
	}
	// Always ensure custom/ exists for user-added sections
	if !seen["custom"] {
		if err := fs.EnsureDir(filepath.Join(pmPath, "custom")); err != nil {
			return fmt.Errorf("creating directory custom: %w", err)
		}
	}

	createdCount := 0
	skippedCount := 0

	for _, def := range tmpl.Sections {
		content := manual.GenerateSectionContent(def)
		path := filepath.Join(pmPath, def.Group, def.Name+".md")
		created, err := fs.WriteFileIfNotExists(path, content)
		if err != nil {
			return fmt.Errorf("writing %s: %w", def.Name+".md", err)
		}
		if created {
			createdCount++
			fmt.Fprintf(w, "  created: %s/%s.md\n", def.Group, def.Name)
		} else {
			skippedCount++
			fmt.Fprintf(w, "  exists:  %s/%s.md (skipped)\n", def.Group, def.Name)
		}
	}

	fmt.Fprintln(w)
	fmt.Fprintf(w, "Initialized .pm/ with %q template — %d file(s) created, %d skipped.\n", tmpl.Name, createdCount, skippedCount)
	fmt.Fprintln(w, "Edit the files in .pm/ to document your project.")
	return nil
}
