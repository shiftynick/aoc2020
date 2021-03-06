package day1

import (
	"fmt"

	"github.com/shiftynick/aoc2020/challenge"
	"github.com/shiftynick/aoc2020/util"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 1, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
		},
	}
}

func b(challenge *challenge.Input) int {
	var entries []int

	for v := range challenge.Lines() {
		entries = append(entries, util.MustAtoI(v))
	}

	for i := range entries {
		for j := i + 1; j < len(entries); j++ {
			for k := j + 1; k < len(entries); k++ {
				if entries[i]+entries[j]+entries[k] == 2020 {
					return entries[i] * entries[j] * entries[k]
				}
			}
		}
	}

	panic("no solution")
}
