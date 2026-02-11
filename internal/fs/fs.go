package fs

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const PMDir = ".pm"

// SearchResult represents a single match from a keyword search.
type SearchResult struct {
	File    string // relative path within .pm/, e.g. "core/deploy.md"
	Line    int    // 1-based line number
	Content string // matched line content (trimmed)
}

// DetectPMDir checks whether a .pm/ directory exists under the given root.
func DetectPMDir(root string) bool {
	info, err := os.Stat(filepath.Join(root, PMDir))
	if err != nil {
		return false
	}
	return info.IsDir()
}

// PMPath returns the absolute path to .pm/ under the given root.
func PMPath(root string) string {
	return filepath.Join(root, PMDir)
}

// ListMarkdownFiles returns .md filenames (without extension) under .pm/<group>/.
// group should be "core" or "custom".
func ListMarkdownFiles(root, group string) ([]string, error) {
	dir := filepath.Join(root, PMDir, group)
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("reading %s: %w", dir, err)
	}

	var names []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		if strings.HasSuffix(e.Name(), ".md") {
			names = append(names, strings.TrimSuffix(e.Name(), ".md"))
		}
	}
	return names, nil
}

// ReadFile reads the full content of a file under .pm/.
func ReadFile(root, relPath string) (string, error) {
	data, err := os.ReadFile(filepath.Join(root, PMDir, relPath))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// FindSection looks for <name>.md in core/ first, then custom/.
// Returns the group ("core" or "custom") and the relative path within .pm/.
// Comparison is case-insensitive.
func FindSection(root, name string) (group, relPath string, err error) {
	lower := strings.ToLower(name)
	for _, g := range []string{"core", "custom"} {
		files, err := ListMarkdownFiles(root, g)
		if err != nil {
			return "", "", err
		}
		for _, f := range files {
			if strings.ToLower(f) == lower {
				return g, filepath.Join(g, f+".md"), nil
			}
		}
	}
	return "", "", fmt.Errorf("section %q not found", name)
}

// Search scans all .md files under .pm/ for lines containing keyword (case-insensitive).
func Search(root, keyword string) ([]SearchResult, error) {
	pmRoot := filepath.Join(root, PMDir)
	lowerKW := strings.ToLower(keyword)

	var results []SearchResult
	err := filepath.Walk(pmRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".md") {
			return nil
		}

		rel, _ := filepath.Rel(pmRoot, path)

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		lineNum := 0
		for scanner.Scan() {
			lineNum++
			line := scanner.Text()
			if strings.Contains(strings.ToLower(line), lowerKW) {
				results = append(results, SearchResult{
					File:    rel,
					Line:    lineNum,
					Content: strings.TrimSpace(line),
				})
			}
		}
		return scanner.Err()
	})

	return results, err
}

// WriteFileIfNotExists creates a file only if it doesn't already exist.
// Parent directories are created as needed.
func WriteFileIfNotExists(path, content string) (created bool, err error) {
	if _, err := os.Stat(path); err == nil {
		return false, nil // already exists
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return false, fmt.Errorf("creating directory: %w", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return false, fmt.Errorf("writing file: %w", err)
	}
	return true, nil
}

// EnsureDir creates a directory (and parents) if it doesn't exist.
func EnsureDir(path string) error {
	return os.MkdirAll(path, 0o755)
}
