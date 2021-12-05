package main

import (
	"fmt"
	"log"

	"github.com/einarssons/adventofcode2021/go/utils"
	"github.com/oriser/regroup"
)

func main() {
	lines := utils.ReadLinesFromFile("test.txt")
	fmt.Printf("Nr test lines: %d\n", len(lines))
	task1(lines, 10)
	task2(lines, 10)
	lines = utils.ReadLinesFromFile("data.txt")
	fmt.Printf("Nr data lines: %d\n", len(lines))
	task1(lines, 1000)
	task2(lines, 1000)
}

func task1(data []string, side int) {
	g := createGrid(side)
	for _, l := range data {
		line := parseLine(l)
		g.AddLine(line, false)
	}
	if side < 12 {
		g.Print()
	}
	nrTwos := g.countTwos()
	fmt.Printf("task1, twos or bigger = %d\n", nrTwos)
}

func task2(data []string, side int) {
	g := createGrid(side)
	for _, l := range data {
		line := parseLine(l)
		g.AddLine(line, true)
		//g.Print()
	}
	nrTwos := g.countTwos()
	fmt.Printf("task2, twos or bigger = %d\n", nrTwos)
}

type Grid struct {
	side int
	p    []int
}

func createGrid(side int) Grid {
	g := Grid{
		side,
		make([]int, side*side),
	}
	return g
}

func (g Grid) Print() {
	for j := 0; j < g.side; j++ {
		for i := 0; i < g.side; i++ {
			p := g.p[j*g.side+i]
			if p == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", p)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (g Grid) countTwos() int {
	c := 0
	for _, x := range g.p {
		if x >= 2 {
			c++
		}
	}
	return c
}

type Line struct {
	X1 int `regroup:"x1"`
	Y1 int `regroup:"y1"`
	X2 int `regroup:"x2"`
	Y2 int `regroup:"y2"`
}

func (g Grid) AddLine(l Line, includeDiagonal bool) {
	//fmt.Printf("%v\n", l)
	if l.X1 == l.X2 {
		min, max := utils.MinMax(l.Y1, l.Y2)
		for i := min; i <= max; i++ {
			g.p[l.X1+i*g.side]++
		}
		return
	}
	if l.Y1 == l.Y2 {
		min, max := utils.MinMax(l.X1, l.X2)
		for i := min; i <= max; i++ {
			g.p[i+l.Y1*g.side]++
		}
		return
	}
	if includeDiagonal {
		xmin, xmax := utils.MinMax(l.X1, l.X2)
		ymin, ymax := utils.MinMax(l.Y1, l.Y2)
		if xmax-xmin != ymax-ymin {
			return
		}
		xstep := 1
		if l.X2 < l.X1 {
			xstep = -1
		}
		ystep := 1
		if l.Y2 < l.Y1 {
			ystep = -1
		}
		nrsteps := xmax - xmin + 1
		x, y := l.X1, l.Y1
		for j := 0; j < nrsteps; j++ {
			g.p[x+y*g.side]++
			x += xstep
			y += ystep
		}
	}
}

var rex = regroup.MustCompile(`(?P<x1>\d+),(?P<y1>\d+) -> (?P<x2>\d+),(?P<y2>\d+)`)

func parseLine(line string) Line {
	l := Line{}
	if err := rex.MatchToTarget(line, &l); err != nil {
		log.Fatal(err)
	}
	return l
}
