package day3

import (
	"fmt"

	"github.com/shiftynick/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 3, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
		},
	}
}

func b(challenge *challenge.Input) int {
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

	treeCts := []int{}
	rights := []int{1, 3, 5, 7, 1}
	downs := []int{1, 1, 1, 1, 2}
	for m := 0; m < 5; m++ {
		j, treeCt := 0, 0
		fmt.Printf("\n\nNEXT %d,%d\n\n", rights[m], downs[m])
		for l := 0; l < len(grid); l++ {
			for k := 0; k < len(grid[l]); k++ {
				if k == j && l%downs[m] == 0 {
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
			if l%downs[m] == 0 {
				j += rights[m]
				if j > len(grid[l])-1 {
					j = j - len(grid[l])
				}
			}
			fmt.Println()
		}
		fmt.Printf("\n\n%d,%d => %d\n\n", rights[m], downs[m], treeCt)
		treeCts = append(treeCts, treeCt)
	}

	ttreeCt := treeCts[0]
	for i := 1; i < 5; i++ {
		ttreeCt *= treeCts[i]
	}
	return ttreeCt
}
