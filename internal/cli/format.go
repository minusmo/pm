package cli

import (
	"fmt"
	"io"
	"strings"

	"github.com/hojooneum/pm/internal/fs"
	"github.com/hojooneum/pm/internal/manual"
)

// PrintSectionList writes a grouped list of sections to w.
func PrintSectionList(w io.Writer, sections []manual.Section) {
	coreItems := []manual.Section{}
	customItems := []manual.Section{}

	for _, s := range sections {
		switch s.Group {
		case "core":
			coreItems = append(coreItems, s)
		case "custom":
			customItems = append(customItems, s)
		}
	}

	if len(coreItems) > 0 {
		fmt.Fprintln(w, "Core sections:")
		for _, s := range coreItems {
			title := s.Title
			if title == "" {
				title = s.Name
			}
			fmt.Fprintf(w, "  %-16s %s\n", s.Name, title)
		}
	}

	if len(customItems) > 0 {
		if len(coreItems) > 0 {
			fmt.Fprintln(w)
		}
		fmt.Fprintln(w, "Custom sections:")
		for _, s := range customItems {
			title := s.Title
			if title == "" {
				title = s.Name
			}
			fmt.Fprintf(w, "  %-16s %s\n", s.Name, title)
		}
	}

	if len(coreItems) == 0 && len(customItems) == 0 {
		fmt.Fprintln(w, "No sections found.")
	}
}

// PrintSearchResults writes search results in grep-like format to w.
func PrintSearchResults(w io.Writer, results []fs.SearchResult) {
	if len(results) == 0 {
		fmt.Fprintln(w, "No matches found.")
		return
	}

	for _, r := range results {
		fmt.Fprintf(w, "%s:%d: %s\n", r.File, r.Line, r.Content)
	}

	fmt.Fprintf(w, "\n%d match(es) found.\n", len(results))
}

// PrintSectionContent writes a section's content to w.
func PrintSectionContent(w io.Writer, s manual.Section) {
	header := s.Title
	if header == "" {
		header = s.Name
	}
	fmt.Fprintf(w, "[%s/%s]\n", s.Group, s.Name)
	fmt.Fprintln(w, strings.Repeat("-", len(header)+len(s.Group)+len(s.Name)+3))
	fmt.Fprintln(w)
	fmt.Fprintln(w, s.Body)
}

// PrintProjectSummary writes a brief project summary with available sections.
func PrintProjectSummary(w io.Writer, sections []manual.Section) {
	fmt.Fprintln(w, "Project manual (.pm/) detected.")
	fmt.Fprintln(w)
	PrintSectionList(w, sections)
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w, "  pm open <section>    Open a section")
	fmt.Fprintln(w, "  pm list              List all sections")
	fmt.Fprintln(w, "  pm search <keyword>  Search across sections")
}

// PrintNoPMDir writes a message when no .pm/ directory is found.
func PrintNoPMDir(w io.Writer) {
	fmt.Fprintln(w, "No .pm/ directory found in the current directory.")
	fmt.Fprintln(w, "Run 'pm init' to create one.")
}
