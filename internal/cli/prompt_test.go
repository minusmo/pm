package cli

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestConfirmYesNo(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		defaultYes bool
		want       bool
	}{
		{"explicit yes", "y\n", true, true},
		{"explicit YES", "YES\n", true, true},
		{"explicit no", "n\n", true, false},
		{"explicit No", "No\n", false, false},
		{"empty default yes", "\n", true, true},
		{"empty default no", "\n", false, false},
		{"EOF default yes", "", true, true},
		{"EOF default no", "", false, false},
		{"invalid then yes", "maybe\ny\n", true, true},
		{"invalid then no", "x\nn\n", true, false},
		{"three invalids uses default", "a\nb\nc\n", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.input))
			var out bytes.Buffer
			got, err := ConfirmYesNo(scanner, &out, "Continue?", tt.defaultYes)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectOption(t *testing.T) {
	options := []string{"default", "minimal"}
	descriptions := []string{"Standard runbook", "Minimal runbook"}

	tests := []struct {
		name       string
		input      string
		defaultIdx int
		want       int
	}{
		{"select first", "1\n", 0, 0},
		{"select second", "2\n", 0, 1},
		{"empty uses default 0", "\n", 0, 0},
		{"empty uses default 1", "\n", 1, 1},
		{"EOF uses default", "", 0, 0},
		{"invalid then valid", "abc\n2\n", 0, 1},
		{"out of range then valid", "5\n1\n", 0, 0},
		{"zero then valid", "0\n2\n", 0, 1},
		{"three invalids uses default", "x\ny\nz\n", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.input))
			var out bytes.Buffer
			got, err := SelectOption(scanner, &out, "Select:", options, descriptions, tt.defaultIdx)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
