package day2

import (
	"fmt"
	"strconv"

	"github.com/shiftynick/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 2, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
		},
	}
}

func b(challenge *challenge.Input) int {
	ct := 0
	ctr := 0
	for v := range challenge.Lines() {
		fmt.Printf("%d: ", ctr)
		matches := reggie.FindAllStringSubmatch(v, -1)
		idx1, _ := strconv.Atoi(matches[0][1])
		idx2, _ := strconv.Atoi(matches[0][2])
		letter := matches[0][3]
		passwd := matches[0][4]
		valid := false
		for i, c := range passwd {
			if fmt.Sprintf("%c", c) == letter && (i+1 == idx1 || i+1 == idx2) {
				if valid {
					valid = false
				} else {
					valid = true
				}
				fmt.Printf("%d ", i+1)
			}
		}
		if valid {
			ct++
			fmt.Printf(" | %s -> valid\n", v)
		} else {
			fmt.Printf(" | %s -> invalid\n", v)
		}
		ctr++
	}
	return ct
}
