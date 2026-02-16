package manual

import (
	"strings"
	"testing"
)

func TestDefaultTemplates_AllPresent(t *testing.T) {
	for _, name := range CoreSectionOrder {
		tmpl, ok := DefaultTemplates[name]
		if !ok {
			t.Errorf("missing template for core section %q", name)
			continue
		}
		if len(tmpl) == 0 {
			t.Errorf("empty template for %q", name)
		}
	}
}

func TestDefaultTemplates_HaveFrontmatter(t *testing.T) {
	for name, tmpl := range DefaultTemplates {
		if !strings.HasPrefix(tmpl, "---\n") {
			t.Errorf("template %q should start with frontmatter delimiter", name)
		}
		// Should have closing ---
		rest := tmpl[4:]
		if !strings.Contains(rest, "\n---\n") {
			t.Errorf("template %q missing closing frontmatter delimiter", name)
		}
	}
}

func TestDefaultTemplates_Parseable(t *testing.T) {
	for name, tmpl := range DefaultTemplates {
		s := ParseSection(name, "core", tmpl)
		if s.Title == "" {
			t.Errorf("template %q should have a title after parsing", name)
		}
		if len(s.Tags) == 0 {
			t.Errorf("template %q should have tags after parsing", name)
		}
	}
}

func TestCoreSectionOrder_SubsetOfTemplates(t *testing.T) {
	for _, name := range CoreSectionOrder {
		if _, ok := DefaultTemplates[name]; !ok {
			t.Errorf("CoreSectionOrder entry %q not found in DefaultTemplates", name)
		}
	}
}
