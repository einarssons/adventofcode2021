package main

import (
	"fmt"

	"github.com/einarssons/adventofcode2021/go/utils"
)

func main() {
	lines := utils.ReadLinesFromFile("test.txt")
	fmt.Printf("Nr test points: %d\n", len(lines))
	task1(lines)
	task2(lines)
	lines = utils.ReadLinesFromFile("data.txt")
	fmt.Printf("Nr data points: %d\n", len(lines))
	task1(lines)
	task2(lines)
}

func task1(lines []string) {
	var counter []int
	for _, line := range lines {
		bits := getBits(line)
		if counter == nil {
			counter = make([]int, len(line))
		}
		for i, b := range bits {
			counter[i] += int(b)
		}
	}
	//fmt.Printf("%v\n", counter)
	binMax := makeFrequencyBinary(counter, len(lines), true)
	binMin := makeFrequencyBinary(counter, len(lines), false)
	fmt.Printf("binMax=%d binMin=%d max+min=%d answer: %d\n", binMax, binMin, binMax+binMin, binMax*binMin)
}

func task2(lines []string) {
	nrBits := len(lines[0])
	nrLines := len(lines)
	matches := make([]bool, nrLines)
	patterns := make([][]byte, 0, nrLines)
	for _, line := range lines {
		patterns = append(patterns, getBits(line))
	}
	var oxGenRating int
	var co2ScrubRating int
	mostCommon := true
	for i := range matches {
		matches[i] = true
	}
	for i := 0; i < nrBits; i++ {
		nrLeft, bits := updateMatches(i, patterns, matches, mostCommon)
		if nrLeft == 1 {
			fmt.Printf("%v  %d\n", bits, makeDec(bits))
			oxGenRating = makeDec(bits)
		}
	}
	for i := range matches {
		matches[i] = true
	}
	mostCommon = false
	for i := 0; i < nrBits; i++ {
		nrLeft, bits := updateMatches(i, patterns, matches, mostCommon)
		if nrLeft == 1 {
			fmt.Printf("%v  %d\n", bits, makeDec(bits))
			co2ScrubRating = makeDec(bits)
		}
	}
	fmt.Printf("oxGenRating=%d, co2ScrubRating=%d, answer: %d\n", oxGenRating, co2ScrubRating, oxGenRating*co2ScrubRating)
}

func getBits(line string) []byte {
	bits := make([]byte, 0, len(line))
	for i := range line {
		bit := line[i] - 48
		bits = append(bits, bit)
	}
	return bits
}

func updateMatches(pos int, patterns [][]byte, matches []bool, mostCommon bool) (nrLeft int, bits []byte) {
	nrZeros := 0
	nrOnes := 0
	nrMatches := 0
	for i, pattern := range patterns {
		if matches[i] {
			nrMatches++
			if pattern[pos] == 1 {
				nrOnes++
			} else {
				nrZeros++
			}
		}
	}
	var bit byte = 0
	if mostCommon {
		if nrOnes >= nrZeros {
			bit = 1
		}
	} else {
		bit = 1
		if nrOnes >= nrZeros {
			bit = 0
		}
	}
	//fmt.Printf("nrmatches=%d\n", nrMatches)
	nrMatches = 0
	var matchIdx int
	for i, pattern := range patterns {
		if matches[i] {
			if pattern[pos] == bit {
				nrMatches++
				matchIdx = i
			} else {
				matches[i] = false
			}
		}
	}
	//fmt.Printf("nrmatches=%d\n", nrMatches)
	return nrMatches, patterns[matchIdx]
}

func makeDec(bits []byte) int {
	number := 0
	for _, bit := range bits {
		number = (number << 1) + int(bit)
	}
	return number
}

func makeFrequencyBinary(counter []int, nrEntries int, countMax bool) int {
	number := 0
	for _, count := range counter {
		bit := 0
		if countMax {
			if count > nrEntries/2 {
				bit = 1
			}
		} else {
			if count < nrEntries/2 {
				bit = 1
			}
		}
		number = (number << 1) + bit
	}
	return number
}
