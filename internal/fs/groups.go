package fs

import (
	"os"
	"path/filepath"
	"sort"
)

// ListGroups returns subdirectory names under .pm/, sorted with "core" first,
// others alphabetically, and "custom" last.
func ListGroups(root string) ([]string, error) {
	pmPath := filepath.Join(root, PMDir)
	entries, err := os.ReadDir(pmPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	var groups []string
	for _, e := range entries {
		if e.IsDir() {
			groups = append(groups, e.Name())
		}
	}

	sort.SliceStable(groups, func(i, j int) bool {
		return groupSortKey(groups[i]) < groupSortKey(groups[j])
	})

	return groups, nil
}

// groupSortKey returns a sort key that places "core" first, "custom" last,
// and everything else alphabetically in between.
func groupSortKey(name string) string {
	switch name {
	case "core":
		return "\x00" // sorts first
	case "custom":
		return "\xff" // sorts last
	default:
		return "\x01" + name // between core and custom
	}
}
