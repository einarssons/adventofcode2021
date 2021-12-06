package main

import (
	"fmt"

	"github.com/einarssons/adventofcode2021/go/utils"
)

func main() {
	lines := utils.ReadLinesFromFile("test.txt")
	fmt.Printf("TEST: Nr test lines: %d\n", len(lines))
	task1(lines)
	task2(lines)
	lines = utils.ReadLinesFromFile("data.txt")
	fmt.Printf("DATA: Nr data lines: %d\n", len(lines))
	task1(lines)
	task2(lines)

}

func task1(data []string) {
	fish := utils.SplitToInts(data[0])
	nrDays := 80
	for d := 0; d < nrDays; d++ {
		fish = stepFish(fish)
		//fmt.Printf("%v\n", fish)
	}
	fmt.Printf("Task1: Nr fish after %d days is %d\n", nrDays, len(fish))
}

// task2 is same as task1, but much bigger numbers, so we need to use
// bins of states instead of individual fish
func task2(data []string) {
	fish := utils.SplitToInts(data[0])
	bins := make([]int, 9)
	for _, f := range fish {
		bins[f]++
	}
	nrDays := 256
	for d := 0; d < nrDays; d++ {
		bins = stepBins(bins)
		//fmt.Printf("%v\n", bins)
	}
	nrFish := 0
	for _, d := range bins {
		nrFish += d
	}
	fmt.Printf("Task 2: Nr fish after %d days is %d\n", nrDays, nrFish)
}

func stepBins(bins []int) []int {
	newBins := make([]int, 9)
	for i := 0; i <= 8; i++ {
		switch i {
		case 0:
			newBins[6] += bins[0]
			newBins[8] += bins[0]
		case 1, 2, 3, 4, 5, 6, 7, 8:
			newBins[i-1] += bins[i]
		}
	}
	return newBins
}

// stepFish - step one day of fish
func stepFish(fish []int) []int {
	var new []int
	nrNew := 0
	for _, f := range fish {
		switch f {
		case 1, 2, 3, 4, 5, 6, 7, 8:
			new = append(new, f-1)
		case 0:
			new = append(new, 6)
			nrNew++
		}
	}
	for i := 0; i < nrNew; i++ {
		new = append(new, 8)
	}
	return new
}
