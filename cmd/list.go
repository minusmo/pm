package cmd

import (
	"os"

	"github.com/hojooneum/pm/internal/cli"
	"github.com/hojooneum/pm/internal/fs"
	"github.com/hojooneum/pm/internal/manual"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list [core|custom]",
	Aliases: []string{"ls"},
	Short:   "List available sections",
	Args:    cobra.MaximumNArgs(1),
	RunE:    runList,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func runList(cmd *cobra.Command, args []string) error {
	root, _ := os.Getwd()
	w := cmd.OutOrStdout()

	if !fs.DetectPMDir(root) {
		cli.PrintNoPMDir(w)
		return nil
	}

	groups := []string{"core", "custom"}
	if len(args) == 1 {
		groups = []string{args[0]}
	}

	var sections []manual.Section
	for _, group := range groups {
		names, err := fs.ListMarkdownFiles(root, group)
		if err != nil {
			return err
		}
		for _, name := range names {
			raw, err := fs.ReadFile(root, group+"/"+name+".md")
			if err != nil {
				return err
			}
			s := manual.ParseSection(name, group, raw)
			sections = append(sections, s)
		}
	}

	cli.PrintSectionList(w, sections)
	return nil
}
