package day2

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/shiftynick/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 2, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
		},
	}
}

var reggie, _ = regexp.Compile(`([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)`)

func a(challenge *challenge.Input) int {
	ct := 0
	ctr := 0
	for v := range challenge.Lines() {
		matches := reggie.FindAllStringSubmatch(v, -1)
		minCt, _ := strconv.Atoi(matches[0][1])
		maxCt, _ := strconv.Atoi(matches[0][2])
		letter := matches[0][3]
		passwd := matches[0][4]
		letterCt := 0
		for _, c := range passwd {
			if fmt.Sprintf("%c", c) == letter {
				letterCt++
			}
		}
		if letterCt >= minCt && letterCt <= maxCt {
			ct++
			fmt.Printf("%d: %s = %d -> valid\n", ctr, v, letterCt)
		} else {
			fmt.Printf("%d: %s = %d -> invalid\n", ctr, v, letterCt)
		}
		ctr++
	}
	return ct
}
