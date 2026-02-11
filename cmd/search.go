package cmd

import (
	"os"

	"github.com/hojooneum/pm/internal/cli"
	"github.com/hojooneum/pm/internal/fs"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search <keyword>",
	Short: "Search for a keyword across all sections",
	Args:  cobra.ExactArgs(1),
	RunE:  runSearch,
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

func runSearch(cmd *cobra.Command, args []string) error {
	root, _ := os.Getwd()
	w := cmd.OutOrStdout()

	if !fs.DetectPMDir(root) {
		cli.PrintNoPMDir(w)
		return nil
	}

	results, err := fs.Search(root, args[0])
	if err != nil {
		return err
	}

	cli.PrintSearchResults(w, results)
	return nil
}
