package day6

import (
	"fmt"

	"github.com/shiftynick/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 6, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
		},
	}
}

func a(challenge *challenge.Input) int {
	totTot := 0

	groupSet := make(map[string]bool)
	for line := range challenge.Lines() {
		if line == "" {
			totTot += len(groupSet)
			groupSet = make(map[string]bool)
		} else {
			for _, c := range line {
				groupSet[fmt.Sprintf("%c", c)] = true
			}
		}
	}

	totTot += len(groupSet)
	// final set of

	return totTot
}
