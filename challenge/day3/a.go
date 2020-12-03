package day3

import (
	"fmt"

	"github.com/shiftynick/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 3, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
		},
	}
}

func a(challenge *challenge.Input) int {
	grid := [][]rune{}

	// build grid
	i := 0
	for line := range challenge.Lines() {
		grid = append(grid, []rune{})
		for _, square := range line {
			grid[i] = append(grid[i], square)
		}
		i++
	}

	j, treeCt := 0, 0
	for l := 0; l < len(grid); l++ {
		for k := 0; k < len(grid[l]); k++ {
			if k == j {
				if grid[l][k] == '.' {
					fmt.Printf("O")
				} else {
					treeCt++
					fmt.Printf("X")
				}
			} else {
				fmt.Printf("%c", grid[l][k])
			}
		}
		j += 3
		if j > len(grid[l])-1 {
			j = j - len(grid[l])
		}
		fmt.Println()
	}

	return treeCt
}
