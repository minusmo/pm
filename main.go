package main

import (
	"os"

	"github.com/hojooneum/pm/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
