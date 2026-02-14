package cli

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const maxRetries = 3

// ConfirmYesNo prints a Y/n confirmation prompt and reads a response.
// defaultYes controls the default when the user presses Enter.
// Returns the user's choice. After maxRetries invalid inputs, returns the default.
func ConfirmYesNo(scanner *bufio.Scanner, w io.Writer, prompt string, defaultYes bool) (bool, error) {
	hint := "[Y/n]"
	if !defaultYes {
		hint = "[y/N]"
	}

	for range maxRetries {
		fmt.Fprintf(w, "%s %s ", prompt, hint)

		if !scanner.Scan() {
			// EOF or error — return default
			return defaultYes, scanner.Err()
		}
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			return defaultYes, nil
		}

		switch strings.ToLower(input) {
		case "y", "yes":
			return true, nil
		case "n", "no":
			return false, nil
		default:
			fmt.Fprintf(w, "  Please enter y or n.\n")
		}
	}

	return defaultYes, nil
}

// SelectOption prints a numbered list of options and reads a selection.
// options and descriptions must have the same length.
// defaultIdx is the 0-based index used when the user presses Enter.
// After maxRetries invalid inputs, returns defaultIdx.
func SelectOption(scanner *bufio.Scanner, w io.Writer, prompt string, options []string, descriptions []string, defaultIdx int) (int, error) {
	fmt.Fprintln(w, prompt)
	for i, opt := range options {
		fmt.Fprintf(w, "  %d) %-16s %s\n", i+1, opt, descriptions[i])
	}
	fmt.Fprintln(w)

	for range maxRetries {
		fmt.Fprintf(w, "Enter choice [%d]: ", defaultIdx+1)

		if !scanner.Scan() {
			return defaultIdx, scanner.Err()
		}
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			return defaultIdx, nil
		}

		n, err := strconv.Atoi(input)
		if err != nil || n < 1 || n > len(options) {
			fmt.Fprintf(w, "  Please enter a number between 1 and %d.\n", len(options))
			continue
		}
		return n - 1, nil
	}

	return defaultIdx, nil
}
