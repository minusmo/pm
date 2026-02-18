package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hojooneum/pm/internal/cli"
	"github.com/hojooneum/pm/internal/fs"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit <section>",
	Short: "Open a section in $EDITOR for editing",
	Args:  cobra.ExactArgs(1),
	RunE:  runEdit,
}

func init() {
	rootCmd.AddCommand(editCmd)
}

func runEdit(cmd *cobra.Command, args []string) error {
	root, _ := os.Getwd()
	w := cmd.OutOrStdout()

	if !fs.DetectPMDir(root) {
		cli.PrintNoPMDir(w)
		return nil
	}

	_, relPath, err := fs.FindSection(root, args[0])
	if err != nil {
		fmt.Fprintf(w, "Error: %v\n\n", err)
		fmt.Fprintln(w, "Available sections:")
		sections, _ := loadAllSections(root)
		cli.PrintSectionList(w, sections)
		return nil
	}

	absPath := filepath.Join(fs.PMPath(root), relPath)

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = os.Getenv("VISUAL")
	}
	if editor == "" {
		editor = "vi"
	}

	c := exec.Command(editor, absPath)
	c.Stdin = cmd.InOrStdin()
	c.Stdout = cmd.OutOrStdout()
	c.Stderr = cmd.ErrOrStderr()
	return c.Run()
}
