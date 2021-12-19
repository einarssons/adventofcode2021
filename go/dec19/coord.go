package main

import (
	u "github.com/einarssons/adventofcode2021/go/utils"
)

type coord struct {
	x, y, z int
}

func (c coord) rotZ(nr int) coord {
	if nr >= 0 {
		for i := 0; i < nr; i++ {
			c.x, c.y = c.y, -c.x
		}
	} else {
		for i := -1; i >= nr; i-- {
			c.x, c.y = -c.y, c.x
		}
	}
	return c
}

func (c coord) rotX(nr int) coord {
	if nr >= 0 {
		for i := 0; i < nr; i++ {
			c.y, c.z = c.z, -c.y
		}
	} else {
		for i := -1; i >= nr; i-- {
			c.y, c.z = -c.x, c.y
		}
	}
	return c
}

func (c coord) rotate(rotNr int) coord {
	rotZ := rotNr / 6
	upSide := rotNr - rotZ*6
	switch upSide {
	case 0:
		// c = c
	case 1:
		c = c.rotX(1)
	case 2:
		c = c.rotZ(1).rotX(1)
	case 3:
		c = c.rotZ(2).rotX(1)
	case 4:
		c = c.rotZ(3).rotX(1)
	case 5:
		c = c.rotX(2)
	}
	switch rotZ {
	case 0:
		// c == c
	case 1:
		c = c.rotZ(1)
	case 2:
		c = c.rotZ(2)
	case 3:
		c = c.rotZ(3)
	}
	return c
}

func coordDiff(a, b coord) coord {
	return coord{a.x - b.x, a.y - b.y, a.z - b.z}
}

// SetCoords - mathematical set with operations
// Empty struct requires zero bytes so is more efficient than bool
type SetCoords map[coord]struct{}

// CreateSetInts - create an empty set
func CreateSetCoords() SetCoords {
	return SetCoords(make(map[coord]struct{}))
}

// Contains - check if elem in set
func (s SetCoords) Contains(elem coord) bool {
	_, ok := s[elem]
	return ok
}

// Add - add elem to set
func (s SetCoords) Add(elem coord) {
	s[elem] = struct{}{}
}

// Remove - remove elem from set (does not need to be in set)
func (s SetCoords) Remove(elem coord) {
	delete(s, elem)
}

// Extend - extend set s with all elements in other (the result is union)
func (s SetCoords) Extend(other SetCoords) {
	for k := range other {
		s[k] = struct{}{}
	}
}

// Subtract - remove all elements from s that are in other
func (s SetCoords) Subtract(other SetCoords) {
	for k := range other {
		_, ok := s[k]
		if ok {
			delete(s, k)
		}
	}
}

// Intersect - only keep elements in s which are also in other
func (s SetCoords) Intersect(other SetCoords) {
	deleteList := make([]coord, 0, len(s))
	for k := range s {
		_, inOther := other[k]
		if !inOther {
			deleteList = append(deleteList, k)
		}
	}
	for _, k := range deleteList {
		delete(s, k)
	}
}

// SetCoordsIntersection - only keep elements in s which are also in other
func SetCoordsIntersection(a, b SetCoords) SetCoords {
	isc := CreateSetCoords()
	for elem := range a {
		if b.Contains(elem) {
			isc.Add(elem)
		}
	}
	return isc
}

func findMaxDist(scanners []scanner) int {

	maxDist := 0

	for i := 0; i < len(scanners); i++ {
		for j := i + 1; j < len(scanners); j++ {
			dist := manhattan(scanners[i].pos, scanners[j].pos)
			if dist > maxDist {
				maxDist = dist
			}
		}
	}
	return maxDist
}

func manhattan(a, b coord) int {
	return u.Abs(a.x-b.x) + u.Abs(a.y-b.y) + u.Abs(a.z-b.z)
}
