package manual

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestListPresets(t *testing.T) {
	presets := ListPresets()
	if len(presets) != 2 {
		t.Fatalf("expected 2 presets, got %d", len(presets))
	}
	// Should be sorted by name
	if presets[0].Name != "default" || presets[1].Name != "minimal" {
		t.Errorf("unexpected preset order: %s, %s", presets[0].Name, presets[1].Name)
	}
}

func TestLoadPreset_Default(t *testing.T) {
	tmpl, err := LoadPreset("default")
	if err != nil {
		t.Fatal(err)
	}
	if len(tmpl.Sections) != 7 {
		t.Fatalf("expected 7 sections, got %d", len(tmpl.Sections))
	}
	for _, s := range tmpl.Sections {
		if s.Group != "core" {
			t.Errorf("expected all default sections to be core, got %q for %q", s.Group, s.Name)
		}
	}
	// Verify names match CoreSectionOrder
	for i, s := range tmpl.Sections {
		if s.Name != CoreSectionOrder[i] {
			t.Errorf("section[%d]: expected %q, got %q", i, CoreSectionOrder[i], s.Name)
		}
	}
}

func TestLoadPreset_Minimal(t *testing.T) {
	tmpl, err := LoadPreset("minimal")
	if err != nil {
		t.Fatal(err)
	}
	if len(tmpl.Sections) != 3 {
		t.Fatalf("expected 3 sections, got %d", len(tmpl.Sections))
	}
	expected := []string{"overview", "deploy", "contacts"}
	for i, s := range tmpl.Sections {
		if s.Name != expected[i] {
			t.Errorf("section[%d]: expected %q, got %q", i, expected[i], s.Name)
		}
	}
}

func TestLoadPreset_Unknown(t *testing.T) {
	_, err := LoadPreset("nonexistent")
	if err == nil {
		t.Fatal("expected error for unknown preset")
	}
	if !strings.Contains(err.Error(), "unknown preset") {
		t.Errorf("expected 'unknown preset' in error, got: %s", err)
	}
}

func TestGenerateSectionContent_KnownName(t *testing.T) {
	def := SectionDef{Name: "deploy", Group: "core", Title: "Deployment Guide"}
	content := GenerateSectionContent(def)
	if content != DefaultTemplates["deploy"] {
		t.Error("expected verbatim DefaultTemplates content for known section name")
	}
}

func TestGenerateSectionContent_UnknownName(t *testing.T) {
	def := SectionDef{
		Name:        "api-docs",
		Group:       "infra",
		Title:       "API Documentation",
		Description: "REST API reference",
		Tags:        []string{"api", "docs"},
	}
	content := GenerateSectionContent(def)

	if !strings.Contains(content, "title: API Documentation") {
		t.Error("expected title in frontmatter")
	}
	if !strings.Contains(content, "description: REST API reference") {
		t.Error("expected description in frontmatter")
	}
	if !strings.Contains(content, "tags: api, docs") {
		t.Error("expected tags in frontmatter")
	}
	if !strings.Contains(content, "# API Documentation") {
		t.Error("expected H1 heading")
	}
	if !strings.Contains(content, "<!-- TODO: Document this section -->") {
		t.Error("expected TODO placeholder")
	}
}

func TestGenerateSectionContent_Parseable(t *testing.T) {
	def := SectionDef{
		Name:        "custom-section",
		Group:       "custom",
		Title:       "Custom Section",
		Description: "A custom section",
		Tags:        []string{"custom", "test"},
	}
	content := GenerateSectionContent(def)
	s := ParseSection(def.Name, def.Group, content)

	if s.Title != "Custom Section" {
		t.Errorf("expected title 'Custom Section', got %q", s.Title)
	}
	if len(s.Tags) != 2 {
		t.Errorf("expected 2 tags, got %d", len(s.Tags))
	}
	if s.Name != "custom-section" {
		t.Errorf("expected name 'custom-section', got %q", s.Name)
	}
}

func TestValidateTemplate_Valid(t *testing.T) {
	tmpl := Template{
		Name: "test",
		Sections: []SectionDef{
			{Name: "overview", Group: "core", Title: "Overview"},
		},
	}
	if err := ValidateTemplate(tmpl); err != nil {
		t.Errorf("expected valid template, got error: %s", err)
	}
}

func TestValidateTemplate_EmptySections(t *testing.T) {
	tmpl := Template{Name: "test", Sections: []SectionDef{}}
	err := ValidateTemplate(tmpl)
	if err == nil {
		t.Fatal("expected error for empty sections")
	}
	if !strings.Contains(err.Error(), "at least one section") {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestValidateTemplate_DuplicateSection(t *testing.T) {
	tmpl := Template{
		Name: "test",
		Sections: []SectionDef{
			{Name: "deploy", Group: "core", Title: "Deploy"},
			{Name: "deploy", Group: "core", Title: "Deploy Again"},
		},
	}
	err := ValidateTemplate(tmpl)
	if err == nil {
		t.Fatal("expected error for duplicate section")
	}
	if !strings.Contains(err.Error(), "duplicate") {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestValidateTemplate_InvalidName(t *testing.T) {
	tmpl := Template{
		Name: "test",
		Sections: []SectionDef{
			{Name: "Invalid Name", Group: "core", Title: "Bad"},
		},
	}
	err := ValidateTemplate(tmpl)
	if err == nil {
		t.Fatal("expected error for invalid section name")
	}
	if !strings.Contains(err.Error(), "must match") {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestLoadTemplateFromFile_Valid(t *testing.T) {
	tmpl := Template{
		Name: "file-test",
		Sections: []SectionDef{
			{Name: "api", Group: "infra", Title: "API Docs", Tags: []string{"api"}},
		},
	}
	data, _ := json.Marshal(tmpl)

	dir := t.TempDir()
	path := filepath.Join(dir, "template.json")
	os.WriteFile(path, data, 0o644)

	loaded, err := LoadTemplateFromFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.Name != "file-test" {
		t.Errorf("expected name 'file-test', got %q", loaded.Name)
	}
	if len(loaded.Sections) != 1 {
		t.Fatalf("expected 1 section, got %d", len(loaded.Sections))
	}
	if loaded.Sections[0].Name != "api" {
		t.Errorf("expected section name 'api', got %q", loaded.Sections[0].Name)
	}
}

func TestLoadTemplateFromFile_InvalidJSON(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.json")
	os.WriteFile(path, []byte("{not valid json"), 0o644)

	_, err := LoadTemplateFromFile(path)
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
	if !strings.Contains(err.Error(), "parsing template JSON") {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestResolveTemplate_PresetName(t *testing.T) {
	tmpl, err := ResolveTemplate("default")
	if err != nil {
		t.Fatal(err)
	}
	if tmpl.Name != "default" {
		t.Errorf("expected 'default', got %q", tmpl.Name)
	}
}

func TestResolveTemplate_FilePath(t *testing.T) {
	tmpl := Template{
		Name: "from-file",
		Sections: []SectionDef{
			{Name: "docs", Group: "core", Title: "Docs"},
		},
	}
	data, _ := json.Marshal(tmpl)

	dir := t.TempDir()
	path := filepath.Join(dir, "t.json")
	os.WriteFile(path, data, 0o644)

	loaded, err := ResolveTemplate(path)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.Name != "from-file" {
		t.Errorf("expected 'from-file', got %q", loaded.Name)
	}
}

func TestResolveTemplate_Neither(t *testing.T) {
	_, err := ResolveTemplate("no-such-thing-xyz")
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "available presets") {
		t.Errorf("expected available presets in error, got: %s", err)
	}
}

func TestResolveTemplate_Empty(t *testing.T) {
	tmpl, err := ResolveTemplate("")
	if err != nil {
		t.Fatal(err)
	}
	if tmpl.Name != "default" {
		t.Errorf("expected 'default' for empty input, got %q", tmpl.Name)
	}
}
