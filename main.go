package main

import (
	"os"

	"github.com/shiftynick/aoc2020/challenge/cmd"
)

//go:generate go run ./gen
func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
