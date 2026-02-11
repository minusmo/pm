package fs

import (
	"os"
	"path/filepath"
	"testing"
)

func setupTestPM(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	for _, sub := range []string{"core", "custom"} {
		if err := os.MkdirAll(filepath.Join(dir, PMDir, sub), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}

func writeTestFile(t *testing.T, root, relPath, content string) {
	t.Helper()
	path := filepath.Join(root, PMDir, relPath)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestDetectPMDir(t *testing.T) {
	t.Run("exists", func(t *testing.T) {
		dir := setupTestPM(t)
		if !DetectPMDir(dir) {
			t.Error("expected .pm/ to be detected")
		}
	})

	t.Run("missing", func(t *testing.T) {
		dir := t.TempDir()
		if DetectPMDir(dir) {
			t.Error("expected .pm/ to not be detected")
		}
	})

	t.Run("file not dir", func(t *testing.T) {
		dir := t.TempDir()
		os.WriteFile(filepath.Join(dir, PMDir), []byte("not a dir"), 0o644)
		if DetectPMDir(dir) {
			t.Error("expected .pm file (not dir) to not be detected")
		}
	})
}

func TestListMarkdownFiles(t *testing.T) {
	dir := setupTestPM(t)
	writeTestFile(t, dir, "core/deploy.md", "# Deploy")
	writeTestFile(t, dir, "core/backup.md", "# Backup")
	writeTestFile(t, dir, "core/readme.txt", "not markdown")

	files, err := ListMarkdownFiles(dir, "core")
	if err != nil {
		t.Fatal(err)
	}

	if len(files) != 2 {
		t.Fatalf("expected 2 files, got %d: %v", len(files), files)
	}

	// Should not include .txt file
	for _, f := range files {
		if f == "readme" {
			t.Error("should not include non-md files")
		}
	}
}

func TestListMarkdownFiles_MissingDir(t *testing.T) {
	dir := t.TempDir()
	files, err := ListMarkdownFiles(dir, "core")
	if err != nil {
		t.Fatal(err)
	}
	if files != nil {
		t.Errorf("expected nil for missing dir, got %v", files)
	}
}

func TestReadFile(t *testing.T) {
	dir := setupTestPM(t)
	writeTestFile(t, dir, "core/test.md", "hello world")

	content, err := ReadFile(dir, "core/test.md")
	if err != nil {
		t.Fatal(err)
	}
	if content != "hello world" {
		t.Errorf("expected 'hello world', got %q", content)
	}
}

func TestFindSection(t *testing.T) {
	dir := setupTestPM(t)
	writeTestFile(t, dir, "core/deploy.md", "# Deploy")
	writeTestFile(t, dir, "custom/myapp.md", "# MyApp")

	t.Run("finds core", func(t *testing.T) {
		group, relPath, err := FindSection(dir, "deploy")
		if err != nil {
			t.Fatal(err)
		}
		if group != "core" || relPath != "core/deploy.md" {
			t.Errorf("unexpected: group=%s path=%s", group, relPath)
		}
	})

	t.Run("finds custom", func(t *testing.T) {
		group, relPath, err := FindSection(dir, "myapp")
		if err != nil {
			t.Fatal(err)
		}
		if group != "custom" || relPath != "custom/myapp.md" {
			t.Errorf("unexpected: group=%s path=%s", group, relPath)
		}
	})

	t.Run("case insensitive", func(t *testing.T) {
		_, _, err := FindSection(dir, "DEPLOY")
		if err != nil {
			t.Error("expected case-insensitive match")
		}
	})

	t.Run("not found", func(t *testing.T) {
		_, _, err := FindSection(dir, "nonexistent")
		if err == nil {
			t.Error("expected error for missing section")
		}
	})
}

func TestSearch(t *testing.T) {
	dir := setupTestPM(t)
	writeTestFile(t, dir, "core/deploy.md", "line 1\nTODO: fix this\nline 3")
	writeTestFile(t, dir, "core/backup.md", "no match here")
	writeTestFile(t, dir, "custom/app.md", "another TODO item")

	results, err := Search(dir, "TODO")
	if err != nil {
		t.Fatal(err)
	}

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	// Verify first result
	if results[0].Line != 2 {
		t.Errorf("expected line 2, got %d", results[0].Line)
	}

	// Verify case insensitive
	results2, _ := Search(dir, "todo")
	if len(results2) != 2 {
		t.Errorf("expected case-insensitive match, got %d results", len(results2))
	}
}

func TestWriteFileIfNotExists(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "sub", "file.md")

	t.Run("creates new file", func(t *testing.T) {
		created, err := WriteFileIfNotExists(path, "content")
		if err != nil {
			t.Fatal(err)
		}
		if !created {
			t.Error("expected file to be created")
		}
		data, _ := os.ReadFile(path)
		if string(data) != "content" {
			t.Errorf("unexpected content: %q", data)
		}
	})

	t.Run("skips existing", func(t *testing.T) {
		created, err := WriteFileIfNotExists(path, "new content")
		if err != nil {
			t.Fatal(err)
		}
		if created {
			t.Error("expected file to be skipped")
		}
		data, _ := os.ReadFile(path)
		if string(data) != "content" {
			t.Error("original content should be preserved")
		}
	})
}
