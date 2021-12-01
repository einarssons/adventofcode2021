package main

import (
	"fmt"

	"github.com/einarssons/adventofcode2021/go/utils"
)

func main() {
	numbers := utils.ReadNumbersFromFile("data.txt")
	part1(numbers)
	part2(numbers)
}

func part1(numbers []int) {
	lastNr := 0
	nrIncreases := 0
	for i, nr := range numbers {
		if i > 0 {
			if nr > lastNr {
				nrIncreases++
			}
		}
		lastNr = nr
	}
	fmt.Println(nrIncreases)
}

func part2(numbers []int) {
	nrIncreases := 0
	for i := range numbers {
		if i >= 2 && i <= len(numbers)-2 {
			slideSum1 := numbers[i-2] + numbers[i-1] + numbers[i]
			slideSum2 := numbers[i-1] + numbers[i] + numbers[i+1]
			if slideSum2 > slideSum1 {
				nrIncreases++
			}
		}
	}
	fmt.Println(nrIncreases)
}
