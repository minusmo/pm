package cli

import (
	"fmt"
	"io"
	"strings"

	"github.com/hojooneum/pm/internal/fs"
	"github.com/hojooneum/pm/internal/manual"
)

// PrintSectionList writes a grouped list of sections to w.
// Groups are printed in the order they first appear in the input.
func PrintSectionList(w io.Writer, sections []manual.Section) {
	// Collect groups in first-appearance order
	var groupOrder []string
	grouped := make(map[string][]manual.Section)
	for _, s := range sections {
		if _, exists := grouped[s.Group]; !exists {
			groupOrder = append(groupOrder, s.Group)
		}
		grouped[s.Group] = append(grouped[s.Group], s)
	}

	if len(groupOrder) == 0 {
		fmt.Fprintln(w, "No sections found.")
		return
	}

	for i, g := range groupOrder {
		if i > 0 {
			fmt.Fprintln(w)
		}
		fmt.Fprintf(w, "%s sections:\n", capitalize(g))
		for _, s := range grouped[g] {
			title := s.Title
			if title == "" {
				title = s.Name
			}
			fmt.Fprintf(w, "  %-16s %s\n", s.Name, title)
		}
	}
}

// PrintTemplateList writes the available templates to w.
func PrintTemplateList(w io.Writer, templates []manual.Template) {
	fmt.Fprintln(w, "Available templates:")
	fmt.Fprintln(w)
	for _, t := range templates {
		desc := t.Description
		if desc == "" {
			desc = fmt.Sprintf("%d section(s)", len(t.Sections))
		}
		fmt.Fprintf(w, "  %-16s %s\n", t.Name, desc)

		for _, s := range t.Sections {
			fmt.Fprintf(w, "                     - %s/%s\n", s.Group, s.Name)
		}
		fmt.Fprintln(w)
	}
	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w, "  pm init --template <name>        Use a built-in preset")
	fmt.Fprintln(w, "  pm init --template <path.json>   Use a custom template file")
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
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
