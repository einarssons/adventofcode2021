package main

import (
	"fmt"
	"sort"
	"strconv"

	u "github.com/einarssons/adventofcode2021/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("test.txt")
	fmt.Printf("Nr test lines: %d\n", len(lines))
	task1(lines)
	task2(lines)
	lines = u.ReadLinesFromFile("data.txt")
	fmt.Printf("Nr data lines: %d\n", len(lines))
	task1(lines)
	task2(lines)
}

func task1(d []string) {
	m := makeMap(d)
	findRisk(m)
}

func task2(d []string) {
	m := makeMap(d)
	findBasins(m)
}

func makeMap(d []string) [][]byte {
	var m [][]byte
	for _, l := range d {
		row := make([]byte, 0, len(l))
		p := u.SplitToChars(l)
		for _, j := range p {
			x, _ := strconv.Atoi(j)
			row = append(row, byte(x))
		}
		m = append(m, row)
	}
	return m
}

func findRisk(m [][]byte) {
	nrRows := len(m)
	nrCols := len(m[0])
	riskSum := 0
	for i := 0; i < nrRows; i++ {
		for j := 0; j < nrCols; j++ {
			v := m[i][j]
			if i > 0 && m[i-1][j] <= v {
				continue
			}
			if i < nrRows-1 && m[i+1][j] <= v {
				continue
			}
			if j > 0 && m[i][j-1] <= v {
				continue
			}
			if j < nrCols-1 && m[i][j+1] <= v {
				continue
			}
			risk := int(v + 1)
			//fmt.Printf("%d %d %d\n", i, j, v)
			riskSum += risk
		}
		// fmt.Printf("En\n")
	}
	fmt.Printf("risk sum = %d\n", riskSum)
}

func findBasins(m [][]byte) {
	nrRows := len(m)
	nrCols := len(m[0])
	var basinSizes []int
	inB := make([][]bool, 0, nrRows)
	for i := 0; i < nrRows; i++ {
		inB = append(inB, make([]bool, nrCols))
	}
	for i := 0; i < nrRows; i++ {
		for j := 0; j < nrCols; j++ {
			v := m[i][j]
			if v != 9 && !inB[i][j] {
				//Inside basin
				bSize := extendBasin(m, inB, i, j, nrRows, nrCols)
				basinSizes = append(basinSizes, bSize)
			}
		}
	}
	sort.Ints(basinSizes)
	n := len(basinSizes)
	prod := basinSizes[n-3] * basinSizes[n-2] * basinSizes[n-1]
	fmt.Printf("basin size prod = %d\n", prod)
}

func extendBasin(m [][]byte, inB [][]bool, i, j, nrRows, nrCols int) int {
	if m[i][j] == 9 || inB[i][j] {
		return 0
	}
	size := 1
	inB[i][j] = true
	if i > 0 {
		size += extendBasin(m, inB, i-1, j, nrRows, nrCols)
	}
	if i < nrRows-1 {
		size += extendBasin(m, inB, i+1, j, nrRows, nrCols)
	}
	if j > 0 {
		size += extendBasin(m, inB, i, j-1, nrRows, nrCols)
	}
	if j < nrCols-1 {
		size += extendBasin(m, inB, i, j+1, nrRows, nrCols)
	}
	return size
}
