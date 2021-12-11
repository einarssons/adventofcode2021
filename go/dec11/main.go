package main

import (
	"fmt"

	u "github.com/einarssons/adventofcode2021/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("test.txt")
	fmt.Printf("### TEST DATA ###\n")
	task1(lines)
	task2(lines)
	lines = u.ReadLinesFromFile("data.txt")
	fmt.Printf("### MY DATA ###\n")
	task1(lines)
	task2(lines)
}

func task1(lines []string) {
	grid := u.CreateDigitGridFromLines(lines)
	nrSteps := 100
	totNrFlashes := 0
	for i := 0; i < nrSteps; i++ {
		totNrFlashes += stepOctopuses(grid)
	}
	fmt.Printf("Task1: total nr flashes: %d\n", totNrFlashes)
}

func task2(lines []string) {
	grid := u.CreateDigitGridFromLines(lines)
	maxNrSteps := 1000
	for i := 0; i < maxNrSteps; i++ {
		nrFlashes := stepOctopuses(grid)
		if nrFlashes == grid.Width*grid.Height {
			fmt.Printf("Task2: complete flash after step %d\n", i+1)
			break
		}
	}
}

func stepOctopuses(g u.DigitGrid) int {
	for j := 0; j < g.Width; j++ {
		for i := 0; i < g.Height; i++ {
			g.Grid[i][j]++
		}
	}
	nrFlashes := 0
	for {
		newUpdates := 0
		for i := 0; i < g.Height; i++ {
			for j := 0; j < g.Width; j++ {
				if g.Grid[i][j] > 9 {
					// Flash, reset and spread to neighbors which are non-zero
					nrFlashes++
					g.Grid[i][j] = 0
					for ni := i - 1; ni <= i+1; ni++ {
						for nj := j - 1; nj <= j+1; nj++ {
							if ni == i && nj == j { // The mid point itself
								continue
							}
							if g.InBounds(ni, nj) && g.Grid[ni][nj] != 0 {
								g.Grid[ni][nj]++
								newUpdates++
							}
						}
					}
				}
			}
		}
		if newUpdates == 0 {
			break
		}
	}
	return nrFlashes
}
