package main

import (
	"fmt"
	"strconv"
	"strings"

	u "github.com/einarssons/adventofcode2021/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("data.txt")
	fmt.Printf("### TASK 1 MY DATA: %d lines ###\n", len(lines))
	task1(lines)
	fmt.Printf("### TASK 2 MY DATA: %d lines ###\n", len(lines))
	task2(lines)
}

func task1(lines []string) {
	sf := parse(u.SplitToChars(lines[0]))
	for _, l := range lines[1:] {
		sf = add(sf, parse(u.SplitToChars(l)), false)
		_, _ = sf.reduce(0)
		//fmt.Printf("%s\n", sf.String())
	}
	m := sf.magnitude()
	fmt.Printf("magnitude = %d\n", m)
}

func task2(lines []string) {
	maxMag := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			sfi := parse(u.SplitToChars(lines[i]))
			sfj := parse(u.SplitToChars(lines[j]))
			sf := add(sfi, sfj, false)
			_, _ = sf.reduce(0)
			m := sf.magnitude()
			if m > maxMag {
				maxMag = m
			}
		}
	}
	fmt.Printf("max magnitude is %d\n", maxMag)
}

type SnailFish struct {
	xn  int
	yn  int
	xnr int // nr in order
	ynr int // nr in order
	xp  *SnailFish
	yp  *SnailFish
}

func (sf *SnailFish) String() string {
	txt := "["
	if sf.xp != nil {
		txt += sf.xp.String()
	} else {
		txt += fmt.Sprintf("%d", sf.xn)
	}
	txt += ","
	if sf.yp != nil {
		txt += sf.yp.String()
	} else {
		txt += fmt.Sprintf("%d", sf.yn)
	}
	txt += "]"
	return txt
}

func parse(chars []string) *SnailFish {
	level := 0
	splitPos := 0
	for i, c := range chars {
		switch c {
		case "[":
			level++
		case "]":
			level--
		case ",":
			if level == 1 {
				splitPos = i
			}
		}
	}
	sf := SnailFish{}
	part1 := chars[1:splitPos]
	nr1, err := strconv.Atoi(strings.Join(part1, ""))
	if err != nil {
		sf.xp = parse(part1)
	} else {
		sf.xn = nr1
	}
	part2 := chars[splitPos+1 : len(chars)-1]
	nr2, err := strconv.Atoi(strings.Join(part2, ""))
	if err != nil {
		sf.yp = parse(part2)
	} else {
		sf.yn = nr2
	}
	return &sf
}

func add(sf1, sf2 *SnailFish, debug bool) *SnailFish {
	sf := SnailFish{xp: sf1, yp: sf2}
	for {
		if debug {
			fmt.Printf("Start: %s\n", &sf)
		}
		sf.setOrderNr(0)
		expl := sf.explode(0)
		if expl.happened {
			if debug {
				fmt.Printf("After explosion: %v\n", sf)
			}

			sf.addExplodedNumbers(&expl)
			if debug {
				fmt.Printf("After propagation: %v\n", sf)
			}

			continue
		}
		hasSplit := sf.split()
		if !hasSplit {
			break
		}
		if debug {
			fmt.Printf("After split: %v\n", sf)
		}

	}
	return &sf
}
func (sf *SnailFish) setOrderNr(nextOrder int) (next int) {
	if sf.xp == nil {
		sf.xnr = nextOrder
		//fmt.Printf("x = %d order %d\n", sf.xn, sf.xnr)
		nextOrder++
	} else {
		nextOrder = sf.xp.setOrderNr(nextOrder)
	}
	if sf.yp == nil {
		sf.ynr = nextOrder
		//fmt.Printf("y = %d order %d\n", sf.yn, sf.ynr)
		nextOrder++
	} else {
		nextOrder = sf.yp.setOrderNr(nextOrder)
	}
	return nextOrder
}

type explosion struct {
	happened     bool
	left         int
	right        int
	leftOrderNr  int
	rightOrderNr int
}

func (sf *SnailFish) addExplodedNumbers(expl *explosion) {
	if sf.xp == nil {
		if sf.xnr == expl.leftOrderNr-1 {
			sf.xn += expl.left
		}
		if sf.xnr == expl.rightOrderNr+1 {
			sf.xn += expl.right
		}
	} else {
		sf.xp.addExplodedNumbers(expl)
	}
	if sf.yp == nil {
		if sf.ynr == expl.leftOrderNr-1 {
			sf.yn += expl.left
		}
		if sf.ynr == expl.rightOrderNr+1 {
			sf.yn += expl.right
		}
	} else {
		sf.yp.addExplodedNumbers(expl)
	}
}

func (sf *SnailFish) explode(level int) explosion {
	if level == 4 {
		return explosion{true, sf.xn, sf.yn, sf.xnr, sf.ynr}
	}

	if sf.xp != nil {
		expl := sf.xp.explode(level + 1)
		if expl.happened {
			if level == 3 { // Back from level 4
				sf.xp = nil
				sf.xn = 0
				sf.xnr = -100000
			}
			return expl
		}
	}

	if sf.yp != nil {
		expl := sf.yp.explode(level + 1)
		if expl.happened {
			if level == 3 { // Back from level 4
				sf.yp = nil
				sf.yn = 0
				sf.ynr = -100000
			}
			return expl
		}
	}
	return explosion{}
}

func (sf *SnailFish) split() bool {
	if sf.xp != nil {
		hasSplit := sf.xp.split()
		if hasSplit {
			return true
		}
	} else {
		if sf.xn >= 10 {
			lv := sf.xn / 2
			rv := sf.xn - lv
			sf.xp = &SnailFish{xn: lv, yn: rv}
			return true
		}

	}
	if sf.yp != nil {
		hasSplit := sf.yp.split()
		if hasSplit {
			return true
		}
	} else {
		if sf.yn >= 10 {
			lv := sf.yn / 2
			rv := sf.yn - lv
			sf.yp = &SnailFish{xn: lv, yn: rv}
			return true
		}
	}
	return false
}

func (sf *SnailFish) reduce(level int) (*int, *int) {
	if level == 4 {
		return &sf.xn, &sf.yn
	}
	var xx, xy *int
	if sf.xp != nil {
		xx, xy = sf.xp.reduce(level + 1)
		if xy != nil {
			if level == 3 && xx != nil { // Back from level 4
				sf.xp = nil
				sf.xn = 0
			}
			if sf.yp == nil {
				sf.yn += *xy
				xy = nil
				if sf.yn >= 10 {
					lv := sf.yn / 2
					rv := sf.yn - lv
					sf.yp = &SnailFish{xn: lv, yn: rv}
				}
			}
		}
	}
	if sf.yp != nil {
		xx, xy = sf.yp.reduce(level + 1)
		if xx != nil {
			if level == 3 && xy != nil { // Back from level 4
				sf.xp = nil
				sf.xn = 0
			}
			if sf.xp == nil {
				sf.xn += *xx
				xx = nil
				if sf.xn >= 10 {
					lv := sf.xn / 2
					rv := sf.xn - lv
					sf.xp = &SnailFish{xn: lv, yn: rv}
				}
			}
		}
	}
	//fmt.Printf("%s\n", sf.Print())
	return xx, xy
}

func (sf *SnailFish) magnitude() int {
	m := 0
	if sf.xp != nil {
		m += 3 * sf.xp.magnitude()
	} else {
		m += 3 * sf.xn
	}
	if sf.yp != nil {
		m += 2 * sf.yp.magnitude()
	} else {
		m += 2 * sf.yn
	}
	return m
}
