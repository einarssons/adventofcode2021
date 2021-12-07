package main

import (
	"fmt"

	"github.com/einarssons/adventofcode2021/go/utils"
)

func main() {
	lines := utils.ReadLinesFromFile("test.txt")
	fmt.Printf("Nr test lines: %d\n", len(lines))
	task1(lines)
	task2(lines)
	lines = utils.ReadLinesFromFile("data.txt")
	fmt.Printf("Nr data lines: %d\n", len(lines))
	task1(lines)
	task2(lines)
}

func task1(input []string) {
	numbers := utils.SplitToInts(input[0])
	findBest(numbers)
}

func task2(input []string) {
	numbers := utils.SplitToInts(input[0])
	findBest2(numbers)

}

func findBest(numbers []int) {
	min, max := utils.MinMaxInts(numbers)
	minFuel := 1000000000
	bestA := -1
	for a := min; a <= max; a++ {
		fuel := 0
		for _, n := range numbers {
			fuel += utils.Abs(a - n)
		}
		if fuel < minFuel {
			bestA = a
			minFuel = fuel
		}
	}
	fmt.Printf("Task 1: pos: %d fuel %d\n", bestA, minFuel)
}
func findBest2(numbers []int) {
	min, max := utils.MinMaxInts(numbers)
	minFuel := 1000000000
	bestA := -1
	for a := min; a <= max; a++ {
		fuel := 0
		for _, n := range numbers {
			fuel += utils.Triangle(utils.Abs(a - n))
		}
		if fuel < minFuel {
			bestA = a
			minFuel = fuel
		}
	}
	fmt.Printf("Task 2: pos: %d fuel %d\n", bestA, minFuel)
}
