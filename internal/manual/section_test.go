package manual

import (
	"testing"
)

func TestParseSection_WithFrontmatter(t *testing.T) {
	raw := `---
title: Deployment Guide
description: How to deploy
tags: deploy, release, ci
---

# Deployment

Steps here.`

	s := ParseSection("deploy", "core", raw)

	if s.Name != "deploy" {
		t.Errorf("expected name 'deploy', got %q", s.Name)
	}
	if s.Group != "core" {
		t.Errorf("expected group 'core', got %q", s.Group)
	}
	if s.Title != "Deployment Guide" {
		t.Errorf("expected title 'Deployment Guide', got %q", s.Title)
	}
	if len(s.Tags) != 3 {
		t.Fatalf("expected 3 tags, got %d: %v", len(s.Tags), s.Tags)
	}
	if s.Tags[0] != "deploy" || s.Tags[1] != "release" || s.Tags[2] != "ci" {
		t.Errorf("unexpected tags: %v", s.Tags)
	}
	if s.Body != "# Deployment\n\nSteps here." {
		t.Errorf("unexpected body: %q", s.Body)
	}
}

func TestParseSection_NoFrontmatter(t *testing.T) {
	raw := "# Just a document\n\nSome content."

	s := ParseSection("readme", "custom", raw)

	if s.Title != "" {
		t.Errorf("expected empty title, got %q", s.Title)
	}
	if s.Body != raw {
		t.Errorf("expected body to be raw content")
	}
}

func TestParseSection_EmptyFrontmatter(t *testing.T) {
	raw := "---\n---\n\nContent after empty frontmatter."

	s := ParseSection("test", "core", raw)

	if s.Title != "" {
		t.Errorf("expected empty title, got %q", s.Title)
	}
	if s.Body != "Content after empty frontmatter." {
		t.Errorf("unexpected body: %q", s.Body)
	}
}

func TestParseSection_UnclosedFrontmatter(t *testing.T) {
	raw := "---\ntitle: Unclosed\nNo closing delimiter"

	s := ParseSection("test", "core", raw)

	// Should treat entire content as body since frontmatter is unclosed
	if s.Title != "" {
		t.Errorf("expected empty title for unclosed frontmatter, got %q", s.Title)
	}
	if s.Body != raw {
		t.Errorf("expected raw content as body")
	}
}

func TestParseSection_SingleTag(t *testing.T) {
	raw := "---\ntags: solo\n---\nBody"

	s := ParseSection("test", "core", raw)

	if len(s.Tags) != 1 || s.Tags[0] != "solo" {
		t.Errorf("expected single tag 'solo', got %v", s.Tags)
	}
}
