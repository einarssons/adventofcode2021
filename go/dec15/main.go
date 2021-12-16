package main

import (
	"fmt"

	u "github.com/einarssons/adventofcode2021/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("test.txt")
	fmt.Printf("--- TASK 1 TEST: %d lines ---\n", len(lines))
	task1(lines)
	lines = u.ReadLinesFromFile("data.txt")
	fmt.Printf("### TASK 1 MY DATA: %d lines ###\n", len(lines))
	task1(lines)
	lines = u.ReadLinesFromFile("test.txt")
	fmt.Printf("--- TASK 2 TEST: %d lines ---\n", len(lines))
	task2(lines)
	lines = u.ReadLinesFromFile("data.txt")
	fmt.Printf("### TASK 2 MY DATA: %d lines ###\n", len(lines))
	task2(lines)
}

const BIG = 10000000

type pos struct {
	y int
	x int
}

func task1(lines []string) {
	risks := u.CreateDigitGridFromLines(lines)
	accRisks := u.CreateZeroDigitGrid(risks.Width, risks.Height)
	visited := u.CreateBoolGrid(risks.Width, risks.Height)
	accRisks.SetValue(BIG)
	accRisks.Grid[0][0] = 0
	expandAccRisks(risks, accRisks, visited, 0, 0)
	for {
		if visited.Grid[risks.Height-1][risks.Width-1] {
			break
		}
		pos := findLowestUnvisited(visited, accRisks)
		expandAccRisks(risks, accRisks, visited, pos.y, pos.x)
	}
	fmt.Printf("total risk: %d\n", accRisks.Grid[risks.Height-1][risks.Width-1])
}

func task2(lines []string) {
	risk1 := u.CreateDigitGridFromLines(lines)
	risks := expandBy5(risk1)
	accRisks := u.CreateZeroDigitGrid(risks.Width, risks.Height)
	visited := u.CreateBoolGrid(risks.Width, risks.Height)
	accRisks.SetValue(BIG)
	accRisks.Grid[0][0] = 0
	expandAccRisks(risks, accRisks, visited, 0, 0)
	for {
		if visited.Grid[risks.Height-1][risks.Width-1] {
			break
		}
		pos := findLowestUnvisited(visited, accRisks)
		expandAccRisks(risks, accRisks, visited, pos.y, pos.x)
	}
	fmt.Printf("total risk: %d\n", accRisks.Grid[risks.Height-1][risks.Width-1])
}

func expandAccRisks(risks, accRisks u.DigitGrid, visited u.BoolGrid, y, x int) {
	if visited.Grid[y][x] {
		return
	}
	value := accRisks.Grid[y][x]
	dirs := []pos{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for _, dir := range dirs {
		ny, nx := y+dir.y, x+dir.x
		if accRisks.InBounds(ny, nx) && !visited.Grid[ny][nx] {
			nValue := value + risks.Grid[ny][nx]
			if nValue < accRisks.Grid[ny][nx] {
				accRisks.Grid[ny][nx] = nValue
			}
		}
	}
	visited.Grid[y][x] = true
}

func findLowestUnvisited(visited u.BoolGrid, accRisks u.DigitGrid) pos {
	lp := pos{}
	lowest := BIG
	for y := 0; y < visited.Height; y++ {
		for x := 0; x < visited.Width; x++ {
			if !visited.Grid[y][x] && accRisks.Grid[y][x] < lowest {
				lowest = accRisks.Grid[y][x]
				lp.x, lp.y = x, y
			}
		}
	}
	return lp
}

func digitMap(nr int, shift int) int {
	new := nr + shift
	if new > 9 {
		new -= 9
	}
	return new
}

func expandBy5(d u.DigitGrid) u.DigitGrid {
	new := u.CreateZeroDigitGrid(d.Width*5, d.Height*5)
	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			for y := 0; y < d.Height; y++ {
				for x := 0; x < d.Width; x++ {
					new.Grid[j*d.Height+y][i*d.Width+x] = digitMap(d.Grid[y][x], j+i)
				}
			}
		}
	}
	return new
}
