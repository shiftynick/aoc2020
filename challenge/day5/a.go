package day5

import (
	"fmt"
	"math"

	"github.com/shiftynick/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 5, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
		},
	}
}

func a(challenge *challenge.Input) int {

	maxId := 0.0
	for line := range challenge.Lines() {
		rMax, rMin, cMax, cMin := 127.0, 0.0, 7.0, 0.0
		for i, c := range line {
			if i < 7 {
				if fmt.Sprintf("%c", c) == "F" {
					rMax = math.Floor(rMax - (rMax-rMin)/2.0)
				} else {
					rMin = math.Ceil(rMin + ((rMax - rMin) / 2.0))
				}
			} else {
				if fmt.Sprintf("%c", c) == "L" {
					cMax = math.Floor((cMax - cMin) / 2.0)
				} else {
					cMin = math.Ceil(cMin + ((cMax - cMin) / 2.0))
				}
			}
		}
		tot := (rMin * 8) + cMin
		maxId = math.Max(maxId, tot)
	}

	return int(maxId)
}
