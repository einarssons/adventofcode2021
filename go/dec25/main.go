package main

import (
	"fmt"

	u "github.com/einarssons/adventofcode2021/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("data.txt")
	task1(lines)
	task2(lines)
}

func task1(lines []string) {
	floor := CreateCucumberGridFromLines(lines)
	fmt.Printf("%s\n", &floor)
	nrSteps := 0
	for {
		nrMoves := floor.Moov()
		nrSteps++
		//fmt.Printf("steps: %d nrMoves %d\n", nrSteps, nrMoves)
		//fmt.Printf("%s\n", &floor)
		if nrMoves == 0 {
			break
		}
		if nrSteps > 1_000_000_000 {
			fmt.Printf("Did not stop after %d steps\n", nrSteps)
			break
		}
	}
	fmt.Printf("Stopped after %d steps\n", nrSteps)
}

func task2(lines []string) {
	fmt.Printf("2: Not implemented\n")
}

type CucumberGrid struct {
	Grid   [][]string
	Width  int
	Height int
}

func CreateCucumberGridFromLines(lines []string) CucumberGrid {
	g := CucumberGrid{}
	for i, line := range lines {
		if i == 0 {
			g.Width = len(line)
		}
		if len(line) != g.Width {
			panic("non-rectangular grid")
		}
		charsRow := u.SplitToChars(line)
		g.Grid = append(g.Grid, charsRow)
		g.Height++
	}
	return g
}

func (c *CucumberGrid) String() string {
	msg := ""
	for row := 0; row < c.Height; row++ {
		for col := 0; col < c.Width; col++ {
			msg += c.Grid[row][col]
		}
		msg += "\n"
	}
	return msg
}

func (c *CucumberGrid) Moov() int {
	nrMoves := 0
	b := u.CreateBoolGrid(c.Width, c.Height)
	// Check Move
	for row := 0; row < c.Height; row++ {
		for col := 0; col < c.Width; col++ {
			if c.Grid[row][col] == ">" {
				col2 := (col + 1) % c.Width
				if c.Grid[row][col2] == "." {
					b.Grid[row][col] = true
					nrMoves++
				}
			}
		}
	}
	// Move
	for row := 0; row < c.Height; row++ {
		for col := 0; col < c.Width; col++ {
			if b.Grid[row][col] {
				col2 := (col + 1) % c.Width
				c.Grid[row][col] = "."
				c.Grid[row][col2] = ">"
			}
		}
	}
	b = u.CreateBoolGrid(c.Width, c.Height)
	// Check Move
	for row := 0; row < c.Height; row++ {
		for col := 0; col < c.Width; col++ {
			if c.Grid[row][col] == "v" {
				row2 := (row + 1) % c.Height
				if c.Grid[row2][col] == "." {
					b.Grid[row][col] = true
					nrMoves++
				}
			}
		}
	}
	// Move
	for row := 0; row < c.Height; row++ {
		for col := 0; col < c.Width; col++ {
			if b.Grid[row][col] {
				row2 := (row + 1) % c.Height
				c.Grid[row][col] = "."
				c.Grid[row2][col] = "v"
			}
		}
	}
	return nrMoves
}
