package day7

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/shiftynick/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 7, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
		},
	}
}

func a(challenge *challenge.Input) int {
	bagRules := make(map[string]map[string]int)

	for line := range challenge.Lines() {
		splits1 := strings.Split(line[0:len(line)-1], " contain ")
		splits2 := strings.Split(splits1[1], ", ")

		for _, cbag := range splits2 {
			if cbag != "no other bags" {
				if _, exists := bagRules[splits1[0]]; !exists {
					bagRules[splits1[0]] = make(map[string]int)
				}
				bagRules[splits1[0]][cbag[2:]], _ = strconv.Atoi(cbag[0:1])
			}
		}

	}

	matches := make(map[string]bool)
	keys := getKeysFromMap(bagRules)
	recurse2(bagRules, matches, keys, "shiny gold bags")

	//totCt := recurse(bagRules, , "shiny gold bags", "")
	return len(matches)
}

func recurse2(bagRules map[string]map[string]int, matches map[string]bool, searchBags []string, myBag string) bool {

	nextKeys := []string{}
	foundMatch := false
	for _, searchBag := range searchBags {
		for rule := range bagRules[searchBag] {
			if rule == myBag {
				matches[searchBag] = true
				foundMatch = true
			} else {
				foundMatch = recurse2(bagRules, matches, searchBags, rule)
				if foundMatch {
					matches[searchBag] = true
				} else {
					nextKeys = append(nextKeys, rule)
				}
			}
		}
		if !foundMatch && len(nextKeys) > 0 {
			foundMatch = recurse2(bagRules, matches, nextKeys, myBag)
			if foundMatch {
				matches[searchBag] = true
			}
		}
	}
	return foundMatch
}

func recurse(bagRules map[string]map[string]int, searchBags []string, myBag string, searchPath string) int {

	ct := 0
	for _, searchBag := range searchBags {
		searchPath += " - " + searchBag
		sBagCan := false
		// see if outer bag can straightup carry my bag
		for rule := range bagRules[searchBag] {
			if rule == myBag {
				sBagCan = true
			}
		}
		// if not recurse, else add to the ct
		keys2 := getKeysFromMap2(bagRules[searchBag])
		if !sBagCan && len(keys2) > 0 {
			sct := recurse(bagRules, keys2, myBag, searchPath)
			if sct > 0 {
				ct++
			}
		} else if sBagCan {
			ct++
		}
	}
	// if ct == 0 {
	// 	fmt.Printf("%s - dead end\n", searchPath)
	// } else {
	// 	fmt.Printf("%s - can\n", searchPath)
	// }
	// if len(strings.Split(searchPath, " - ")) == 2 {
	// 	fmt.Println("Asdsadoaisdoasidnosai")
	// }
	return ct
}

func getKeysFromMap(m map[string]map[string]int) []string {
	keys := make([]string, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	return keys
}

func getKeysFromMap2(m map[string]int) []string {
	keys := make([]string, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	return keys
}
