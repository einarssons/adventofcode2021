package utils

import "strconv"

type DigitGrid struct {
	Grid   [][]int
	Width  int
	Height int
}

func CreateDigitGridFromLines(lines []string) DigitGrid {
	g := DigitGrid{}
	for i, line := range lines {
		if i == 0 {
			g.Width = len(line)
		}
		if len(line) != g.Width {
			panic("non-rectangular grid")
		}
		row := make([]int, 0, g.Width)
		digits := SplitToChars(line)
		for _, digit := range digits {
			nr, err := strconv.Atoi(digit)
			if err != nil {
				panic(err)
			}
			row = append(row, nr)
		}
		g.Grid = append(g.Grid, row)
		g.Height++
	}
	return g
}

// InBounds - is (i, j) in grid
func (g DigitGrid) InBounds(i, j int) bool {
	return 0 <= i && i < g.Height && 0 <= j && j < g.Width
}
