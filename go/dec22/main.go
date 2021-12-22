package main

import (
	"fmt"
	"strings"

	u "github.com/einarssons/adventofcode2021/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("data.txt")
	task1(lines)
	totVolume := task2(lines)
	fmt.Printf("Task 2: Total cubes on: %d\n", totVolume)
}

func task1(lines []string) {
	c := createCuboid(101, 101, 101, -50, -50, -50)
	for _, line := range lines {
		step := parseLine(line)
		c.step(step)
	}
	fmt.Printf("Task 1: Total cubes on: %d\n", c.nrOn())
}

// task2 - strategy is to split into cuboids with on state
func task2(lines []string) int {
	var rects []Cuboid
	step := parseLine(lines[0])
	c := CreateCuboid(step)
	rects = append(rects, c)
	for _, line := range lines[1:] {
		step := parseLine(line)
		c := CreateCuboid(step)
		inRects := []Cuboid{c}
		if step.on {
			for _, r := range rects {
				var outRects []Cuboid
				for _, ic := range inRects {
					outRects = append(outRects, GetProtruding(r, ic)...)
				}
				inRects = outRects
			}
			rects = append(rects, inRects...)
		} else {
			var outRects []Cuboid
			for _, r := range rects {
				outRects = append(outRects, GetProtruding(c, r)...)
			}
			rects = outRects
		}
	}
	totVolume := 0
	for _, r := range rects {
		totVolume += r.Volume()
	}
	return totVolume
}

type step struct {
	on                     bool
	x0, x1, y0, y1, z0, z1 int
}

func parseLine(line string) step {
	v := true
	if strings.HasPrefix(line, "off") {
		v = false
	}
	q := u.SplitToInts(line)
	return step{v, q[0], q[1], q[2], q[3], q[4], q[5]}
}

type cuboid struct {
	xoff int
	yoff int
	zoff int
	xlen int
	ylen int
	zlen int
	c    [][][]bool
}

func createCuboid(xlen, ylen, zlen, xoff, yoff, zoff int) *cuboid {
	cu := cuboid{xoff, yoff, zoff, xlen, ylen, zlen, nil}
	c := make([][][]bool, xlen)
	for x := 0; x < xlen; x++ {
		cy := make([][]bool, ylen)
		for y := 0; y < ylen; y++ {
			cz := make([]bool, zlen)
			cy[y] = cz
		}
		c[x] = cy
	}
	cu.c = c
	return &cu
}

func (c *cuboid) step(s step) {
	zmin := max(c.zoff, s.z0)
	zmax := min(c.zoff+c.zlen-1, s.z1)
	ymin := max(c.yoff, s.y0)
	ymax := min(c.yoff+c.ylen-1, s.y1)
	xmin := max(c.xoff, s.x0)
	xmax := min(c.xoff+c.xlen-1, s.x1)

	for z := zmin; z <= zmax; z++ {
		for y := ymin; y <= ymax; y++ {
			for x := xmin; x <= xmax; x++ {
				dx, dy, dz := x-c.xoff, y-c.yoff, z-c.zoff
				if dx < 0 || dx >= c.xlen {
					panic("Bad x")
				}
				if dy < 0 || dy >= c.ylen {
					panic("Bad y")
				}
				if dz < 0 || dz >= c.zlen {
					panic("Bad z")
				}
				c.c[dx][dy][dz] = s.on
			}
		}
	}
}

func (c *cuboid) nrOn() int {
	count := 0
	for z := 0; z < c.zlen; z++ {
		for y := 0; y < c.ylen; y++ {
			for x := 0; x < c.xlen; x++ {
				if c.c[x][y][z] {
					count++
				}
			}
		}
	}
	return count
}

// Cuboid - x_min, x_max, y_min, y_max, z_min, z_max
type Cuboid [6]int

func CreateCuboid(s step) Cuboid {
	return Cuboid{s.x0, s.x1, s.y0, s.y1, s.z0, s.z1}
}

func (r Cuboid) Volume() int {
	return (r[1] - r[0] + 1) * (r[3] - r[2] + 1) * (r[5] - r[4] + 1)
}

func (r Cuboid) Contains(o Cuboid) bool {
	for i := 0; i < 3; i++ {
		if r[2*i] > o[2*i] || r[2*i+1] < o[2*i+1] {
			return false // min of r > min of o or max of r < max of o
		}
	}
	return true
}

// Separate - true if no overlap
func (r Cuboid) Separate(o Cuboid) bool {
	for i := 0; i < 3; i++ {
		// min of one is bigger than  max of other
		if r[2*i] > o[2*i+1] || o[2*i] > r[2*i+1] {
			return true
		}
	}
	return false
}

func (r Cuboid) CutOverlap(r2, cutC Cuboid) []Cuboid {
	var cuts []Cuboid
	for i := 0; i < 3; i++ {
		if cutC[2*i] < r[2*i] { //Something protuding on lower side r2[2*i] = cutC[2*i] < r[2*i]
			// First set protruding part of r2 to be limited by cutC and r2
			prot := Cuboid{max(r2[0], cutC[0]), min(r2[1], cutC[1]), max(r2[2], cutC[2]), min(r2[3], cutC[3]), max(r2[4], cutC[4]), min(r2[5], cutC[5])}
			prot[2*i], prot[2*i+1] = r2[2*i], r[2*i]-1 // Cut from low r2 to low r-1
			cutC[2*i] = r[2*i]                         // Change cut to avoid protruding part
			cuts = append(cuts, prot)
		}
		if cutC[2*i+1] > r[2*i+1] { //Something protuding on higher side r2[2*i+1] = cutC[2*i+1] > r[2*i+1]
			prot := Cuboid{max(r2[0], cutC[0]), min(r2[1], cutC[1]), max(r2[2], cutC[2]), min(r2[3], cutC[3]), max(r2[4], cutC[4]), min(r2[5], cutC[5])}
			prot[2*i], prot[2*i+1] = r[2*i+1]+1, r2[2*i+1] // Cut from high r+1 to high r2
			cutC[2*i+1] = r[2*i+1]                         // Change cut to avoid protruding in this direction
			cuts = append(cuts, prot)
		}
	}
	return cuts
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// GetProtruding - get everything from r2 which is not in r1
func GetProtruding(r1, r2 Cuboid) []Cuboid {
	if len(r2) == 0 {
		return nil
	}
	if r1.Separate(r2) {
		return []Cuboid{r2}
	}
	if r1.Contains(r2) {
		return []Cuboid{}
	}
	// Try to cut off parts and return them
	cutCubd := Cuboid{}
	for i := 0; i < 3; i++ {
		cutCubd[2*i] = min(r1[2*i], r2[2*i])
		cutCubd[2*i+1] = max(r1[2*i+1], r2[2*i+1])
	}
	cds := r1.CutOverlap(r2, cutCubd)
	return cds
}
