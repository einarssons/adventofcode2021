package main

import (
	"fmt"
	"strconv"
	"strings"

	u "github.com/einarssons/adventofcode2021/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("test.txt")
	fmt.Printf("--- TEST: %d lines ---\n", len(lines))
	task1(lines)
	lines = u.ReadLinesFromFile("data.txt")
	fmt.Printf("### MY DATA: %d lines ###\n", len(lines))
	task1(lines)
}

type coord struct {
	x, y int
}

type fold struct {
	dir string
	nr  int
}

type paper struct {
	dots   Set
	width  int
	height int
}

func (p *paper) add(c coord) {
	if c.x+1 > p.width {
		p.width = c.x + 1
	}
	if c.y+1 > p.height {
		p.height = c.y + 1
	}
	p.dots.Add(c)
}

func createPaper() *paper {
	p := paper{}
	p.dots = CreateSet()
	return &p
}

// fold - fold paper along x column or y row
func (p *paper) fold(f fold) {
	if f.dir == "x" {
		for d := range p.dots {
			if d.x > f.nr {
				dx := d.x - f.nr
				p.dots.Delete(d)
				p.add(coord{d.x - 2*dx, d.y})
			}
		}
		min, max := 100000000, -1000000
		for d := range p.dots {
			if d.x < min {
				min = d.x
			}
			if d.x > max {
				max = d.x
			}
		}
		newSet := CreateSet()
		for d := range p.dots {
			newSet.Add(coord{d.x - min, d.y})
		}
		p.dots = newSet
		p.width = max - min + 1
	} else {
		for d := range p.dots {
			if d.y > f.nr {
				dy := d.y - f.nr
				p.dots.Delete(d)
				p.add(coord{d.x, d.y - 2*dy})
			}
		}
		min, max := 100000000, -1000000
		for d := range p.dots {
			if d.y < min {
				min = d.y
			}
			if d.y > max {
				max = d.y
			}
		}
		newSet := CreateSet()
		for d := range p.dots {
			newSet.Add(coord{d.x - min, d.y})
		}
		p.dots = newSet
		p.height = max - min + 1
	}
}

func (p *paper) Print() {
	// fmt.Printf("paper width=%d height=%d\n", p.width, p.height)
	for y := 0; y < p.height; y++ {
		for x := 0; x < p.width; x++ {
			c := p.dots.Contains(coord{x, y})
			if c {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func task1(lines []string) {
	instrPart := false
	var coords []coord
	var folds []fold
	for _, l := range lines {
		if l == "" {
			instrPart = true
			continue
		}
		if !instrPart {
			c2 := u.SplitToInts(l)
			coords = append(coords, coord{c2[0], c2[1]})
		} else {
			x := strings.TrimPrefix(l, "fold along ")
			c2 := strings.Split(x, "=")
			n, _ := strconv.Atoi(c2[1])
			folds = append(folds, fold{c2[0], n})
		}
	}
	// fmt.Printf("%d %d\n", len(coords), len(folds))
	//nrs := u.SplitToInts(d)
	p := createPaper()
	for _, c := range coords {
		p.add(c)
	}
	// fmt.Printf("width=%d height=%d\n", p.width, p.height)
	p.fold(folds[0])
	//fmt.Printf("fold %+v: width=%d height=%d\n", folds[0], p.width, p.height)
	fmt.Printf("Task 1: Nr dots after one fold: %d\n", len(p.dots))
	for i := 1; i < len(folds); i++ {
		p.fold(folds[i])
	}
	fmt.Printf("Task 2: Secret code:\n")
	p.Print()
}

type Set map[coord]struct{}

func CreateSet() Set {
	return Set(make(map[coord]struct{}))
}

func (s Set) Contains(elem coord) bool {
	_, ok := s[elem]
	return ok
}

func (s Set) Add(elem coord) {
	s[elem] = struct{}{}
}

func (s Set) Delete(elem coord) {
	delete(s, elem)
}
