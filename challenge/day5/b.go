package day5

import (
	"fmt"
	"math"

	//"sort"

	"github.com/shiftynick/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 5, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
		},
	}
}

func b(challenge *challenge.Input) int {

	foundSeats := make(map[float64]bool)
	maxID, minID := 0.0, 1025.0
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
					cMax = math.Floor(cMax - (cMax-cMin)/2.0)
				} else {
					cMin = math.Ceil(cMin + ((cMax - cMin) / 2.0))
				}
			}
		}
		seatID := (rMin * 8) + cMin
		maxID = math.Max(maxID, seatID)
		minID = math.Min(minID, seatID)
		foundSeats[seatID] = true
	}

	// keys := make([]float64, 0, len(foundSeats))
	// for k := range foundSeats {
	// 	keys = append(keys, k)
	// }
	// sort.Float64s(keys)

	// for _, k := range keys {
	// 	fmt.Println(k, foundSeats[k])
	// }

	freeID := 0.0
	for i := 0.0; i < 128.0; i++ {
		for j := 0.0; j < 8.0; j++ {
			checkSeatID := (i * 8.0) + j
			if checkSeatID > minID && checkSeatID < maxID {
				if _, prs := foundSeats[checkSeatID]; !prs {
					freeID = checkSeatID
					fmt.Printf("FREE [%d,%d] %d, %d => %d\n", int(minID), int(maxID), int(i), int(j), int(freeID))
				} else {
					//fmt.Printf("TAKE [%d,%d] %d, %d => %d\n", int(minID), int(maxID), int(i), int(j), int(checkSeatID))
				}
			}
		}
	}

	return int(freeID)
}
