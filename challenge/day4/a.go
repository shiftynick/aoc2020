package day4

import (
	"fmt"
	"strings"

	"github.com/shiftynick/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 4, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
		},
	}
}

func a(challenge *challenge.Input) int {
	req := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	m := make(map[string]bool)
	valid := 0
	for line := range challenge.Lines() {
		if line == "" {
			ct := 0
			for _, r := range req {
				if b, prs := m[r]; prs && b {
					ct++
				}
			}
			if ct == 7 {
				valid++
			}
			m = make(map[string]bool)
		} else {
			splits1 := strings.Split(line, " ")
			for _, s := range splits1 {
				splits2 := strings.Split(s, ":")
				m[splits2[0]] = true
			}
		}
	}

	// check final map
	ct := 0
	for _, r := range req {
		if b, prs := m[r]; prs && b {
			ct++
		}
	}
	if ct == 7 {
		valid++
	}

	return valid
}
