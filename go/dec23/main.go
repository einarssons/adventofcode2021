package main

import (
	"fmt"
	"strings"

	u "github.com/einarssons/adventofcode2021/go/utils"
)

var myData = `#############
#...........#
###D#C#D#B###
  #B#A#A#C#
  #########`

var myData2 = `#############
#...........#
###D#C#D#B###
  #D#C#B#A#
  #D#B#A#C#
  #B#A#A#C#
  #########`

func main() {
	lines := strings.Split(myData, "\n")
	task1Mamual(lines)
	task1(lines)
	lines = strings.Split(myData2, "\n")
	task2(lines)
}

func task1Mamual(lines []string) {
	apods := parse(lines)
	c := cave{apods, 0, 2}
	fmt.Printf("%d\n%s\n\n", c.cost, c.String())
	c.moveApod(3, 3, 0) // B
	fmt.Printf("%d\n%s\n\n", c.cost, c.String())
	c.moveApod(7, 5, 0) // C
	fmt.Printf("%d\n%s\n\n", c.cost, c.String())
	c.moveApod(2, 8, 2) // D-> D home
	fmt.Printf("%d\n%s\n\n", c.cost, c.String())
	c.moveApod(6, 10, 0) // A -> right
	fmt.Printf("%d\n%s\n\n", c.cost, c.String())
	c.moveApod(7, 6, 2) // C -> home 2
	fmt.Printf("%d\n%s\n\n", c.cost, c.String())
	c.moveApod(1, 6, 1) // C -> home 1
	fmt.Printf("%d\n%s\n\n", c.cost, c.String())
	c.moveApod(5, 9, 0) // A-> right
	fmt.Printf("%d\n%s\n\n", c.cost, c.String())
	c.moveApod(3, 4, 2) // B -> home
	fmt.Printf("%d\n%s\n\n", c.cost, c.String())
	c.moveApod(0, 8, 1) // D -> D home
	c.moveApod(4, 4, 1) // B -> home
	c.moveApod(5, 2, 2) // A -> home
	c.moveApod(6, 2, 1)
	fmt.Printf("%d\n%s\n\n", c.cost, c.String())
	fmt.Printf("manual lowest cost = %d\n", c.cost)
}

func task1(lines []string) {
	maxCost := 15000
	apods := parse(lines)
	c := cave{apods, 0, 2}
	lowestCost := play1(c, 0, maxCost, 0)
	fmt.Printf("task1: lowest cost = %d\n", lowestCost)
}

func task2(lines []string) {
	maxCost := 100000
	apods := parse(lines)
	c := cave{apods, 0, 4}
	lowestCost := play1(c, 0, maxCost, 0)
	fmt.Printf("task2: lowest cost = %d\n", lowestCost)
}

func play1(c cave, prevCost, maxCost, level int) (lowestCost int) {
	lowestCost = maxCost
	//fmt.Printf("Level = %d\n%s\n", level, c.String())
	for aNr := 0; aNr < len(c.apods); aNr++ {
		moves := c.possibleMoves(aNr, prevCost, maxCost)
		//fmt.Printf("aNr %d spec %d: %+v\n", aNr, c.apods[aNr].spec, moves)
		for _, m := range moves {
			newC := copyCave(c)
			newC.apods[aNr].x, newC.apods[aNr].y = m.x, m.y
			newC.cost += m.cost
			if newC.solved() {
				if newC.cost < lowestCost {
					lowestCost = newC.cost
					//fmt.Printf("Lowest cost = %d, level = %d\n", lowestCost, level)
					//fmt.Printf("%s\n", &c)
				}
				return
			}
			newLowestCost := play1(newC, prevCost+m.cost, maxCost, level+1)
			if newLowestCost < lowestCost {
				lowestCost = newLowestCost
			}
			//fmt.Printf("Back at level %d, aNr = %d\n", level, aNr)
			//fmt.Printf("Back at level = %d, aNr = %d\n%s\n", level, aNr, c.String())
		}
	}
	return lowestCost
}

func copyCave(c cave) cave {
	newApods := make([]apod, len(c.apods))
	for i := range c.apods {
		newApods[i] = c.apods[i]
	}
	nc := cave{
		apods: newApods,
		cost:  c.cost,
		depth: c.depth,
	}
	return nc
}

type apod struct {
	spec int // 0-3 = A-D
	cost int
	x    int
	y    int
}

func (a apod) homeX() int {
	return 2 + 2*a.spec
}

type cave struct {
	apods []apod
	cost  int
	depth int
}

func (c *cave) String() string {
	msg := "#############\n"
	msg += "#"
	y := 0
	for x := 0; x <= 10; x++ {
		msg += c.occChar(x, y)
	}
	msg += "#\n###"
	y = 1
	for x := 2; x <= 8; x += 2 {
		msg += c.occChar(x, y)
		msg += "#"
	}
	msg += "##\n"
	for y := 2; y <= c.depth; y++ {
		msg += "  #"
		for x := 2; x <= 8; x += 2 {
			msg += c.occChar(x, y)
			msg += "#"
		}
		msg += "\n"
	}
	msg += "  #########"
	return msg
}

func (c *cave) solved() bool {
	for _, a := range c.apods {
		if a.x != a.homeX() {
			return false
		}
	}
	return true
}

func (c *cave) occChar(x, y int) string {
	for _, a := range c.apods {
		if a.x == x && a.y == y {
			return string(rune(a.spec) + 65)
		}
	}
	return "."
}

type target struct {
	x, y, cost int
}

func (c *cave) possibleMoves(aNr int, prevCost, maxCost int) []target {
	if c.alreadyInHome(aNr) {
		return nil
	}
	a := c.apods[aNr]

	// Next move to home if possible
	if cost, ok := c.checkMoveToHome(aNr); ok {
		t := target{a.homeX(), c.spaceInHome(aNr), cost}
		return []target{t}
	}

	// If at y == 0 level, only move to home is allowed. Therefore quit
	if a.y == 0 {
		return nil
	}

	// Need to move from y > 0 to y == 0 and an x position
	// The x-position must be 0<= x <= 10 and not 2, 4, 6,  8
	// Use step to step in either positive or negative direction
	var targets []target
	maxStep := 10
	switch a.spec {
	case 3:
		maxStep = 2
	case 2:
		maxStep = 4
	case 1:
		maxStep = 8
	case 0:
		maxStep = 8
	}
	for step := 1; step <= maxStep; step++ {
		for dir := -1; dir <= 1; dir += 2 {
			xn := a.x + step*dir
			if xn < 0 || xn == 2 || xn == 4 || xn == 6 || xn == 8 || xn > 10 {
				continue
			}
			if cost, ok := c.checkMoveUp(aNr, xn); ok {
				if prevCost+cost < maxCost {
					targets = append(targets, target{xn, 0, cost})
				}
			}
		}

	}
	return targets
}

func (c *cave) moveApod(aNr, x, y int) bool {
	a := c.apods[aNr]
	nrSteps := u.Abs(a.x-x) + y + a.y
	c.apods[aNr].x, c.apods[aNr].y = x, y
	c.cost += nrSteps * a.cost
	return true
}

func (c *cave) checkMoveToHome(aNr int) (cost int, ok bool) {
	if c.alreadyInHome(aNr) {
		return 0, false
	}
	a := c.apods[aNr]
	homeLevel := c.spaceInHome(aNr)
	if homeLevel == 0 {
		return 0, false
	}
	if a.y > 1 && !c.isFree(a.x, a.y-1) {
		return 0, false //Blocked from moving up

	}
	minX, maxX := u.MinMax(a.x, a.homeX())
	for x := minX; x < maxX; x++ { // Check intermediate positions
		if a.y == 0 && a.x == x {
			continue
		}
		if !c.isFree(x, 0) {
			return 0, false
		}
	}
	nrStepsX := maxX - minX
	// Steps up, sidewise, down
	cost = (a.y + nrStepsX + homeLevel) * a.cost
	return cost, true
}

// checkMoveUp, check if can move from inside a home up to y == 0
func (c *cave) checkMoveUp(aNr, x int) (cost int, ok bool) {
	a := c.apods[aNr]
	if a.y > 1 && !c.isFree(a.x, a.y-1) { // Cannot move up
		return 0, false
	}
	xMin, xMax := u.MinMax(a.x, x)
	for nx := xMin; nx <= xMax; nx++ {
		if !c.isFree(nx, 0) {
			return 0, false
		}
	}
	return (a.y + xMax - xMin) * a.cost, true
}

func (c *cave) isFree(x, y int) bool {
	for _, apod := range c.apods {
		if apod.x == x && apod.y == y {
			return false
		}
	}
	return true
}

func (c *cave) occupied(x, y int) (spec int, ok bool) {
	for _, apod := range c.apods {
		if apod.x == x && apod.y == y {
			return apod.spec, true
		}
	}
	return -1, false
}

// alreadyInHome -  check if apod aNr is in home
func (c *cave) alreadyInHome(aNr int) bool {
	apod := c.apods[aNr]
	homeX := apod.homeX()
	inHome := apod.x == homeX
	homeSpec := apod.spec
	if !inHome {
		return false
	}
	// Check below for right apods
	for lvl := c.depth; lvl > apod.y; lvl-- {
		if occSpec, ok := c.occupied(homeX, lvl); ok {
			if occSpec != homeSpec {
				return false
			}
		}
	}
	return true
}

// Check if home is available. Return level where 0 means not available
func (c *cave) spaceInHome(aNr int) (level int) {
	apod := c.apods[aNr]
	homeX := apod.homeX()
	homeSpec := apod.spec
	for lvl := c.depth; lvl > 0; lvl-- {
		if c.isFree(homeX, lvl) { // Nothing on this level or above
			return lvl
		}
		// Check if other spec in this level. If so, home is unavailable
		occSpec, ok := c.occupied(homeX, lvl)
		if !ok {
			panic("This should be occupied since not free")
		}
		if occSpec != homeSpec {
			return 0
		}
	}
	return 0
}

func parse(lines []string) []apod {
	apods := make([]apod, 0, 8)
	nrLines := len(lines)
	for i, c := range u.SplitToChars(lines[1]) {
		xPos := i - 1
		switch c {
		case "A":
			apods = append(apods, apod{0, 1, xPos, 0})
		case "B":
			apods = append(apods, apod{1, 10, xPos, 0})
		case "C":
			apods = append(apods, apod{2, 100, xPos, 0})
		case "D":
			apods = append(apods, apod{3, 1000, xPos, 0})
		}
	}
	for i, l := range lines[2 : nrLines-1] {
		cs := u.SplitToChars(l)
		xPos := 2
		for _, c := range cs {
			switch c {
			case "A":
				apods = append(apods, apod{0, 1, xPos, i + 1})
				xPos += 2
			case "B":
				apods = append(apods, apod{1, 10, xPos, i + 1})
				xPos += 2
			case "C":
				apods = append(apods, apod{2, 100, xPos, i + 1})
				xPos += 2
			case "D":
				apods = append(apods, apod{3, 1000, xPos, i + 1})
				xPos += 2
			case ".":
				xPos += 2
			}
		}
	}
	return apods
}
