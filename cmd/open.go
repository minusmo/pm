package cmd

import (
	"fmt"
	"os"

	"github.com/hojooneum/pm/internal/cli"
	"github.com/hojooneum/pm/internal/fs"
	"github.com/hojooneum/pm/internal/manual"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open <section>",
	Short: "Open and display a section",
	Args:  cobra.ExactArgs(1),
	RunE:  runOpen,
}

func init() {
	rootCmd.AddCommand(openCmd)
}

func runOpen(cmd *cobra.Command, args []string) error {
	root, _ := os.Getwd()
	w := cmd.OutOrStdout()

	if !fs.DetectPMDir(root) {
		cli.PrintNoPMDir(w)
		return nil
	}

	group, relPath, err := fs.FindSection(root, args[0])
	if err != nil {
		fmt.Fprintf(w, "Error: %v\n\n", err)
		fmt.Fprintln(w, "Available sections:")
		sections, _ := loadAllSections(root)
		cli.PrintSectionList(w, sections)
		return nil
	}

	raw, err := fs.ReadFile(root, relPath)
	if err != nil {
		return err
	}

	s := manual.ParseSection(args[0], group, raw)
	cli.PrintSectionContent(w, s)
	return nil
}
