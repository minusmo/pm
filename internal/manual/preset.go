package manual

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// SectionDef describes a section to scaffold in a template.
type SectionDef struct {
	Name        string   `json:"name"`
	Group       string   `json:"group"`
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

// Template describes a set of sections to scaffold with pm init.
type Template struct {
	Name        string       `json:"name"`
	Description string       `json:"description,omitempty"`
	Sections    []SectionDef `json:"sections"`
}

var sectionNamePattern = regexp.MustCompile(`^[a-z0-9][a-z0-9-]*$`)

// built-in presets keyed by name
var builtinPresets = map[string]Template{
	"default": {
		Name:        "default",
		Description: "Standard runbook with 7 core sections",
		Sections: []SectionDef{
			{Name: "overview", Group: "core", Title: "Project Overview", Description: "High-level summary of this project", Tags: []string{"overview", "architecture"}},
			{Name: "deploy", Group: "core", Title: "Deployment Guide", Description: "Step-by-step deployment procedures", Tags: []string{"deploy", "release"}},
			{Name: "troubleshoot", Group: "core", Title: "Troubleshooting Guide", Description: "Common issues and resolution steps", Tags: []string{"troubleshoot", "incident", "debug"}},
			{Name: "backup", Group: "core", Title: "Backup & Recovery", Description: "Backup procedures and disaster recovery", Tags: []string{"backup", "recovery", "disaster"}},
			{Name: "maintenance", Group: "core", Title: "Maintenance Procedures", Description: "Routine maintenance tasks and schedules", Tags: []string{"maintenance", "cron", "cleanup"}},
			{Name: "monitoring", Group: "core", Title: "Monitoring & Alerts", Description: "Monitoring setup, dashboards, and alert runbooks", Tags: []string{"monitoring", "alerts", "metrics"}},
			{Name: "contacts", Group: "core", Title: "Contacts & Escalation", Description: "Team contacts and escalation procedures", Tags: []string{"contacts", "oncall", "escalation"}},
		},
	},
	"minimal": {
		Name:        "minimal",
		Description: "Minimal runbook with 3 essential sections",
		Sections: []SectionDef{
			{Name: "overview", Group: "core", Title: "Project Overview", Description: "High-level summary of this project", Tags: []string{"overview", "architecture"}},
			{Name: "deploy", Group: "core", Title: "Deployment Guide", Description: "Step-by-step deployment procedures", Tags: []string{"deploy", "release"}},
			{Name: "contacts", Group: "core", Title: "Contacts & Escalation", Description: "Team contacts and escalation procedures", Tags: []string{"contacts", "oncall", "escalation"}},
		},
	},
	"onboarding": {
		Name:        "onboarding",
		Description: "New developer onboarding with 6 sections",
		Sections: []SectionDef{
			{Name: "overview", Group: "core", Title: "Project Overview", Description: "High-level summary of this project", Tags: []string{"overview", "architecture"}},
			{Name: "setup-guide", Group: "core", Title: "Setup Guide", Description: "Local development environment setup instructions", Tags: []string{"setup", "install", "environment"}},
			{Name: "codebase-walkthrough", Group: "core", Title: "Codebase Walkthrough", Description: "Guided tour of the project structure and key modules", Tags: []string{"codebase", "structure", "architecture"}},
			{Name: "dev-workflow", Group: "core", Title: "Development Workflow", Description: "Day-to-day development process and branch strategy", Tags: []string{"workflow", "git", "branching"}},
			{Name: "coding-conventions", Group: "core", Title: "Coding Conventions", Description: "Code style, naming, and best practices for this project", Tags: []string{"conventions", "style", "standards"}},
			{Name: "contacts", Group: "core", Title: "Contacts & Escalation", Description: "Team contacts and escalation procedures", Tags: []string{"contacts", "oncall", "escalation"}},
		},
	},
	"microservice": {
		Name:        "microservice",
		Description: "Microservice runbook with 9 sections",
		Sections: []SectionDef{
			{Name: "overview", Group: "core", Title: "Project Overview", Description: "High-level summary of this project", Tags: []string{"overview", "architecture"}},
			{Name: "service-dependencies", Group: "core", Title: "Service Dependencies", Description: "Upstream and downstream service dependencies", Tags: []string{"dependencies", "services", "integration"}},
			{Name: "api-contracts", Group: "core", Title: "API Contracts", Description: "API endpoints, schemas, and contract specifications", Tags: []string{"api", "contracts", "schema"}},
			{Name: "health-checks", Group: "core", Title: "Health Checks", Description: "Service health check endpoints and liveness/readiness probes", Tags: []string{"health", "probes", "liveness", "readiness"}},
			{Name: "scaling", Group: "core", Title: "Scaling Guide", Description: "Horizontal and vertical scaling strategies", Tags: []string{"scaling", "performance", "capacity"}},
			{Name: "deploy", Group: "core", Title: "Deployment Guide", Description: "Step-by-step deployment procedures", Tags: []string{"deploy", "release"}},
			{Name: "monitoring", Group: "core", Title: "Monitoring & Alerts", Description: "Monitoring setup, dashboards, and alert runbooks", Tags: []string{"monitoring", "alerts", "metrics"}},
			{Name: "troubleshoot", Group: "core", Title: "Troubleshooting Guide", Description: "Common issues and resolution steps", Tags: []string{"troubleshoot", "incident", "debug"}},
			{Name: "contacts", Group: "core", Title: "Contacts & Escalation", Description: "Team contacts and escalation procedures", Tags: []string{"contacts", "oncall", "escalation"}},
		},
	},
	"library": {
		Name:        "library",
		Description: "Library/package documentation with 7 sections",
		Sections: []SectionDef{
			{Name: "overview", Group: "core", Title: "Project Overview", Description: "High-level summary of this project", Tags: []string{"overview", "architecture"}},
			{Name: "api-reference", Group: "core", Title: "API Reference", Description: "Public API surface and usage documentation", Tags: []string{"api", "reference", "documentation"}},
			{Name: "usage-examples", Group: "core", Title: "Usage Examples", Description: "Practical examples and common use cases", Tags: []string{"examples", "usage", "quickstart"}},
			{Name: "versioning", Group: "core", Title: "Versioning", Description: "Version strategy, changelog, and compatibility policy", Tags: []string{"versioning", "semver", "changelog"}},
			{Name: "publishing", Group: "core", Title: "Publishing", Description: "Release and publishing procedures", Tags: []string{"publishing", "release", "distribution"}},
			{Name: "contributing", Group: "core", Title: "Contributing", Description: "Guidelines for contributing to this project", Tags: []string{"contributing", "guidelines", "community"}},
			{Name: "contacts", Group: "core", Title: "Contacts & Escalation", Description: "Team contacts and escalation procedures", Tags: []string{"contacts", "oncall", "escalation"}},
		},
	},
	"framework": {
		Name:        "framework",
		Description: "Framework documentation with 7 sections",
		Sections: []SectionDef{
			{Name: "overview", Group: "core", Title: "Project Overview", Description: "High-level summary of this project", Tags: []string{"overview", "architecture"}},
			{Name: "getting-started", Group: "core", Title: "Getting Started", Description: "Quick start guide for new users of this framework", Tags: []string{"getting-started", "quickstart", "tutorial"}},
			{Name: "plugin-system", Group: "core", Title: "Plugin System", Description: "Plugin architecture, extension points, and authoring guide", Tags: []string{"plugins", "extensions", "hooks"}},
			{Name: "migration-guide", Group: "core", Title: "Migration Guide", Description: "Upgrade paths and breaking change migration instructions", Tags: []string{"migration", "upgrade", "breaking-changes"}},
			{Name: "contributing", Group: "core", Title: "Contributing", Description: "Guidelines for contributing to this project", Tags: []string{"contributing", "guidelines", "community"}},
			{Name: "versioning", Group: "core", Title: "Versioning", Description: "Version strategy, changelog, and compatibility policy", Tags: []string{"versioning", "semver", "changelog"}},
			{Name: "contacts", Group: "core", Title: "Contacts & Escalation", Description: "Team contacts and escalation procedures", Tags: []string{"contacts", "oncall", "escalation"}},
		},
	},
}

// ListPresets returns all built-in presets sorted by name.
func ListPresets() []Template {
	names := make([]string, 0, len(builtinPresets))
	for n := range builtinPresets {
		names = append(names, n)
	}
	sort.Strings(names)

	presets := make([]Template, len(names))
	for i, n := range names {
		presets[i] = builtinPresets[n]
	}
	return presets
}

// LoadPreset returns a built-in preset by name.
func LoadPreset(name string) (Template, error) {
	t, ok := builtinPresets[name]
	if !ok {
		available := make([]string, 0, len(builtinPresets))
		for n := range builtinPresets {
			available = append(available, n)
		}
		sort.Strings(available)
		return Template{}, fmt.Errorf("unknown preset %q (available: %s)", name, strings.Join(available, ", "))
	}
	return t, nil
}

// LoadTemplateFromFile reads and validates a JSON template from a file.
func LoadTemplateFromFile(path string) (Template, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Template{}, fmt.Errorf("reading template file: %w", err)
	}

	var t Template
	if err := json.Unmarshal(data, &t); err != nil {
		return Template{}, fmt.Errorf("parsing template JSON: %w", err)
	}

	if err := ValidateTemplate(t); err != nil {
		return Template{}, fmt.Errorf("invalid template: %w", err)
	}

	return t, nil
}

// ResolveTemplate resolves a template by name (built-in preset) or file path.
// An empty string resolves to the "default" preset.
func ResolveTemplate(nameOrPath string) (Template, error) {
	if nameOrPath == "" {
		return LoadPreset("default")
	}

	// Try as preset name first
	if t, err := LoadPreset(nameOrPath); err == nil {
		return t, nil
	}

	// Try as file path
	if _, err := os.Stat(nameOrPath); err == nil {
		return LoadTemplateFromFile(nameOrPath)
	}

	available := make([]string, 0, len(builtinPresets))
	for n := range builtinPresets {
		available = append(available, n)
	}
	sort.Strings(available)
	return Template{}, fmt.Errorf("%q is not a known preset or valid file path (available presets: %s)", nameOrPath, strings.Join(available, ", "))
}

// ValidateTemplate checks that a template has all required fields.
func ValidateTemplate(t Template) error {
	if t.Name == "" {
		return fmt.Errorf("template name is required")
	}
	if len(t.Sections) == 0 {
		return fmt.Errorf("template must have at least one section")
	}

	seen := make(map[string]bool)
	for i, s := range t.Sections {
		if s.Name == "" {
			return fmt.Errorf("section[%d]: name is required", i)
		}
		if !sectionNamePattern.MatchString(s.Name) {
			return fmt.Errorf("section[%d]: name %q must match %s", i, s.Name, sectionNamePattern.String())
		}
		if s.Group == "" {
			return fmt.Errorf("section[%d]: group is required", i)
		}
		if s.Title == "" {
			return fmt.Errorf("section[%d]: title is required", i)
		}

		key := s.Group + "/" + s.Name
		if seen[key] {
			return fmt.Errorf("duplicate section %q in group %q", s.Name, s.Group)
		}
		seen[key] = true
	}
	return nil
}

// GenerateSectionContent returns markdown content for a section definition.
// If the section name matches a DefaultTemplate, that content is returned verbatim.
// Otherwise, a generic placeholder is generated from the definition's metadata.
func GenerateSectionContent(def SectionDef) string {
	if tmpl, ok := DefaultTemplates[def.Name]; ok {
		return tmpl
	}

	var b strings.Builder
	b.WriteString("---\n")
	b.WriteString("title: " + def.Title + "\n")
	if def.Description != "" {
		b.WriteString("description: " + def.Description + "\n")
	}
	if len(def.Tags) > 0 {
		b.WriteString("tags: " + strings.Join(def.Tags, ", ") + "\n")
	}
	b.WriteString("---\n\n")
	b.WriteString("# " + def.Title + "\n\n")
	b.WriteString("<!-- TODO: Document this section -->\n")
	return b.String()
}
