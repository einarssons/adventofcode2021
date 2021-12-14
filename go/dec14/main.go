package main

import (
	"fmt"
	"log"
	"strings"

	u "github.com/einarssons/adventofcode2021/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("test.txt")
	fmt.Printf("--- TASK 1 TEST: %d lines ---\n", len(lines))
	task1(lines)
	fmt.Printf("--- TASK 2 TEST: %d lines ---\n", len(lines))
	task2(lines)
	lines = u.ReadLinesFromFile("data.txt")
	fmt.Printf("### TASK 1 MY DATA: %d lines ###\n", len(lines))
	task1(lines)
	fmt.Printf("### TASK 2 MY DATA: %d lines ###\n", len(lines))
	task2(lines)
}

func task1(lines []string) {
	polymer := lines[0]
	rules := make(map[string]string)

	for _, line := range lines[2:] {
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1]
	}
	//fmt.Printf("%d %s\n", len(polymer), polymer)
	nrSteps := 10
	for i := 0; i < nrSteps; i++ {
		pairs := splitToPairs(polymer)
		polymer = joinPairs(pairs, rules)
		//fmt.Printf("%d %s\n", len(polymer), polymer)
	}
	counter := make(map[string]int)
	chars := u.SplitToChars(polymer)
	for _, c := range chars {
		counter[c] += 1
	}
	min, max := getMinMax(counter)
	fmt.Printf("max - min: %d - %d = %d\n", max, min, max-min)
}

func task2(lines []string) {
	polymer := lines[0]
	rules := make(map[string]string)

	for _, line := range lines[2:] {
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1]
	}
	pairs := splitToPairs(polymer)

	pairsCount := make(map[string]int)
	for _, p := range pairs {
		pairsCount[p]++
	}
	lastPair := pairs[len(pairs)-1]

	nrSteps := 40
	for i := 0; i < nrSteps; i++ {
		newPC := make(map[string]int)
		for pair, c := range pairsCount {
			newInsert := rules[pair]
			newPair1 := pair[0:1] + newInsert
			newPair2 := newInsert + pair[1:2]
			newPC[newPair1] += c
			newPC[newPair2] += c
		}
		// Add last pair
		lastInsert := rules[lastPair]
		newLast := lastInsert + lastPair[1:2]
		lastPair = newLast
		pairsCount = newPC
	}
	counter := make(map[string]int)
	for pair, v := range pairsCount {
		counter[pair[0:1]] += v
	}
	counter[lastPair[1:2]]++ // Last letter has not been counted yet
	//fmt.Printf("%+v\n", counter)
	min, max := getMinMax(counter)
	fmt.Printf("max - min: %d - %d = %d\n", max, min, max-min)
}

func splitToPairs(polymer string) []string {
	pairs := make([]string, 0, len(polymer)-1)
	for i := 0; i < len(polymer)-1; i++ {
		pairs = append(pairs, polymer[i:i+2])
	}
	return pairs
}

func joinPairs(pairs []string, rules map[string]string) string {
	out := ""
	for _, pair := range pairs {
		x, ok := rules[pair]
		if !ok {
			log.Fatalf("unknown pair %s\n", pair)
		}
		out += pair[0:1] + x
	}
	out += pairs[len(pairs)-1][1:2]
	return out
}

func getMinMax(counter map[string]int) (min, max int) {
	min = 1 << 50
	max = -1
	for _, v := range counter {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}
