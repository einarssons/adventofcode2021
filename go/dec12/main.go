package main

import (
	"fmt"
	"strings"

	u "github.com/einarssons/adventofcode2021/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("data.txt")
	fmt.Printf("task1: nr paths = %d\n", task(lines, false))
	fmt.Printf("task2: nr paths = %d\n", task(lines, true))
}

//var allPaths []string  // Uncomment to get collect all paths

func task(lines []string, allowOneDoubleLow bool) (nrPaths int) {
	//allPaths = nil
	allLegs := getLegs(lines)

	for _, startLeg := range allLegs["start"] {
		path := "start-" + startLeg
		nrPaths += exploreFromPath(path, allLegs, allowOneDoubleLow)
	}
	return nrPaths
}

// exploreFromPath - explore possible paths further than path
func exploreFromPath(path string, allLegs map[string][]string, allowOneDoubleLow bool) int {
	points := strings.Split(path, "-")
	doubleVisits := calcDoubleLowVisits(points)
	end := points[len(points)-1]
	nrPaths := 0
	for _, new := range allLegs[end] {
		if new == "start" {
			continue
		}
		if new == "end" {
			// endPath := path + "-" + "end"
			// allPaths = append(allPaths, endPath)
			nrPaths++
			continue
		}
		if strings.ToLower(new) == new {
			if strings.Contains(path, "-"+new) {
				if !allowOneDoubleLow {
					continue
				} else if doubleVisits > 0 { // Already one double visit
					continue
				}
			}
		}
		newPath := path + "-" + new
		nrPaths += exploreFromPath(newPath, allLegs, allowOneDoubleLow)
	}
	return nrPaths
}

func calcDoubleLowVisits(points []string) int {
	doubleVisits := 0
	low := u.CreateSet()
	for _, p := range points {
		if strings.ToLower(p) == p {
			if low.Contains(p) {
				doubleVisits++
				break
			} else {
				low.Add(p)
			}
		}
	}
	return doubleVisits
}

// getLegs - get a map of all legs with src as key and dst as value
func getLegs(d []string) map[string][]string {
	legs := make(map[string][]string)

	for _, l := range d {
		parts := strings.Split(l, "-")
		if _, ok := legs[parts[0]]; !ok {
			legs[parts[0]] = make([]string, 0, 10)
		}
		legs[parts[0]] = append(legs[parts[0]], parts[1])
		if _, ok := legs[parts[1]]; !ok {
			legs[parts[1]] = make([]string, 0, 10)
		}
		legs[parts[1]] = append(legs[parts[1]], parts[0])
	}
	return legs
}
