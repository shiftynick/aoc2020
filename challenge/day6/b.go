package day6

import (
	"fmt"

	"github.com/shiftynick/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 6, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
		},
	}
}

func b(challenge *challenge.Input) int {
	totTot := 0

	groupSet := make(map[string]int)
	groupCt := 0
	for line := range challenge.Lines() {
		if line == "" {
			for k := range groupSet { // Loop
				if ct, _ := groupSet[k]; ct == groupCt {
					totTot++
				}
			}
			groupSet = make(map[string]int)
			groupCt = 0
		} else {
			for _, c := range line {
				letter := fmt.Sprintf("%c", c)
				ct, exs := groupSet[letter]
				if !exs {
					ct = 0
				}
				groupSet[letter] = ct + 1
			}
			groupCt++
		}
	}

	for k := range groupSet { // Loop
		if ct, _ := groupSet[k]; ct == groupCt {
			totTot++
		}
	}

	return totTot
}
