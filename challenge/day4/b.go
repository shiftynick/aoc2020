package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/shiftynick/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 4, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
		},
	}
}

func b(challenge *challenge.Input) int {
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
				fmt.Printf("VALID")
				valid++
			}
			fmt.Println()
			m = make(map[string]bool)
		} else {
			splits1 := strings.Split(line, " ")
			for _, s := range splits1 {
				splits2 := strings.Split(s, ":")
				m[splits2[0]] = validate(splits2[0], splits2[1])
				fmt.Printf("%s:%s => %t\n", splits2[0], splits2[1], m[splits2[0]])
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

func validate(key string, val string) bool {

	if key == "byr" {
		if len(val) == 4 {
			if result, err := strconv.Atoi(val); err == nil {
				return result >= 1920 && result <= 2002
			}
		}
	} else if key == "iyr" {
		if len(val) == 4 {
			if result, err := strconv.Atoi(val); err == nil {
				return result >= 2010 && result <= 2020
			}
		}
	} else if key == "eyr" {
		if len(val) == 4 {
			if result, err := strconv.Atoi(val); err == nil {
				return result >= 2020 && result <= 2030
			}
		}
	} else if key == "hgt" {
		reggie, _ := regexp.Compile("([0-9]+)(cm|in)")
		match := reggie.FindString(val)
		if match != "" {
			if strings.HasSuffix(val, "cm") {
				if result, err := strconv.Atoi(strings.Replace(match, "cm", "", -1)); err == nil {
					return result >= 150 && result <= 193
				}
			} else {
				if result, err := strconv.Atoi(strings.Replace(match, "in", "", -1)); err == nil {
					return result >= 59 && result <= 76
				}
			}
		}
	} else if key == "hcl" {
		if len(val) == 7 {
			if match, _ := regexp.MatchString("#[0-9a-f]{6}", val); match {
				return true
			}
		}
	} else if key == "ecl" {
		if len(val) == 3 {
			if match, _ := regexp.MatchString("amb|blu|brn|gry|grn|hzl|oth", val); match {
				return true
			}
		}
	} else if key == "pid" {
		if len(val) == 9 {
			if match, _ := regexp.MatchString("[0-9]{9}", val); match {
				return true
			}
		}
	}
	return false
}
