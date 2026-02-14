package cmd

import (
	"bufio"
	"fmt"
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
		w := cmd.OutOrStdout()
		fmt.Fprintln(w, "No .pm/ directory found in the current directory.")
		fmt.Fprintln(w)

		if isInteractive() {
			return runInteractiveInit(cmd, root)
		}

		fmt.Fprintln(w, "Run 'pm init' to create one.")
		return nil
	}

	sections, err := loadAllSections(root)
	if err != nil {
		return err
	}

	cli.PrintProjectSummary(cmd.OutOrStdout(), sections)
	return nil
}

// isInteractive returns true when stdin is a terminal (not piped/redirected).
func isInteractive() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}

// runInteractiveInit drives the interactive init flow: confirm, pick template, scaffold.
func runInteractiveInit(cmd *cobra.Command, root string) error {
	w := cmd.OutOrStdout()
	scanner := bufio.NewScanner(os.Stdin)

	ok, err := cli.ConfirmYesNo(scanner, w, "Would you like to create one?", true)
	if err != nil {
		return err
	}
	if !ok {
		fmt.Fprintln(w, "Run 'pm init' to create one when you're ready.")
		return nil
	}

	fmt.Fprintln(w)

	presets := manual.ListPresets()
	options := make([]string, len(presets))
	descriptions := make([]string, len(presets))
	for i, p := range presets {
		options[i] = p.Name
		descriptions[i] = p.Description
	}

	idx, err := cli.SelectOption(scanner, w, "Select a template:", options, descriptions, 0)
	if err != nil {
		return err
	}

	fmt.Fprintln(w)
	return doInit(w, root, presets[idx])
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
