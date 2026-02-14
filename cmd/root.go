package cmd

import (
	"os"

	"github.com/hojooneum/pm/internal/cli"
	"github.com/hojooneum/pm/internal/fs"
	"github.com/hojooneum/pm/internal/manual"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pm",
	Short: "Project manual — manage and browse runbooks from .pm/",
	Long:  "pm is a CLI tool for managing project-specific runbooks and manuals stored in .pm/ directories.",
	RunE:  runRoot,
}

func runRoot(cmd *cobra.Command, args []string) error {
	root, _ := os.Getwd()

	if !fs.DetectPMDir(root) {
		cli.PrintNoPMDir(cmd.OutOrStdout())
		return nil
	}

	sections, err := loadAllSections(root)
	if err != nil {
		return err
	}

	cli.PrintProjectSummary(cmd.OutOrStdout(), sections)
	return nil
}

// loadAllSections reads and parses all sections from all groups under .pm/.
func loadAllSections(root string) ([]manual.Section, error) {
	groups, err := fs.ListGroups(root)
	if err != nil {
		return nil, err
	}

	var sections []manual.Section
	for _, group := range groups {
		names, err := fs.ListMarkdownFiles(root, group)
		if err != nil {
			return nil, err
		}
		for _, name := range names {
			raw, err := fs.ReadFile(root, group+"/"+name+".md")
			if err != nil {
				return nil, err
			}
			s := manual.ParseSection(name, group, raw)
			sections = append(sections, s)
		}
	}

	return sections, nil
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}
