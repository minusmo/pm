package manual

import (
	"strings"
)

// Section represents a parsed .pm/ markdown document.
type Section struct {
	Name  string   // filename without .md extension
	Group string   // "core" or "custom"
	Title string   // from frontmatter "title:" field
	Tags  []string // from frontmatter "tags:" field
	Body  string   // content after frontmatter
}

// ParseSection parses raw markdown content into a Section.
// Frontmatter is delimited by "---" lines. Supported keys: title, description, tags.
func ParseSection(name, group, raw string) Section {
	s := Section{
		Name:  name,
		Group: group,
	}

	lines := strings.Split(raw, "\n")
	if len(lines) == 0 || strings.TrimSpace(lines[0]) != "---" {
		s.Body = raw
		return s
	}

	// Find closing ---
	closeIdx := -1
	for i := 1; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "---" {
			closeIdx = i
			break
		}
	}

	if closeIdx == -1 {
		s.Body = raw
		return s
	}

	// Parse frontmatter key: value pairs
	for _, line := range lines[1:closeIdx] {
		idx := strings.Index(line, ":")
		if idx == -1 {
			continue
		}
		key := strings.TrimSpace(line[:idx])
		val := strings.TrimSpace(line[idx+1:])

		switch strings.ToLower(key) {
		case "title":
			s.Title = val
		case "tags":
			for _, t := range strings.Split(val, ",") {
				t = strings.TrimSpace(t)
				if t != "" {
					s.Tags = append(s.Tags, t)
				}
			}
		}
	}

	// Body is everything after closing ---
	if closeIdx+1 < len(lines) {
		s.Body = strings.Join(lines[closeIdx+1:], "\n")
		s.Body = strings.TrimLeft(s.Body, "\n")
	}

	return s
}
