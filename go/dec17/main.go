package main

import (
	"fmt"
	"log"

	u "github.com/einarssons/adventofcode2021/go/utils"
	"github.com/oriser/regroup"
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

func task1(lines []string) {
	target := parseTarget(lines[0])
	i := 1
	var maxXVel int
	var minXVel int
	for {
		t := u.Triangle(i)
		if t >= target.X1 && minXVel == 0 {
			minXVel = i
		}
		if t > target.X2 {
			maxXVel = i - 1
			break
		}
		i++
	}
	maxHeight := 0
	for vx := minXVel; vx <= maxXVel; vx++ {
		for vy := 0; vy < 1000; vy++ {
			m := shoot(pair{vx, vy}, target)
			if m > maxHeight {
				maxHeight = m
				//fmt.Printf("%d  %v\n", m, pair{vx, vy})
			}

		}
	}
	fmt.Printf("maxHeight = %d\n", maxHeight)
}

func task2(lines []string) {
	target := parseTarget(lines[0])
	i := 1
	var maxXVel int
	var minXVel int
	nrHits := 0
	for {
		t := u.Triangle(i)
		if t >= target.X1 && minXVel == 0 {
			minXVel = i
			break
		}
		i++
	}
	maxXVel = target.X2
	maxHeight := 0
	for vx := minXVel; vx <= maxXVel; vx++ {
		for vy := target.Y1; vy < 1000; vy++ {
			m := shoot(pair{vx, vy}, target)
			if m >= 0 {
				nrHits++
			}
			if m > maxHeight {
				maxHeight = m
				//fmt.Printf("%d  %v\n", m, pair{vx, vy})
			}

		}
	}
	fmt.Printf("maxHeight = %d nrHits = %d\n", maxHeight, nrHits)
}

// shoot - Return maxHeight or -1 if missed
func shoot(vel pair, t Target) int {
	maxHeight := 0
	pos := pair{0, 0}
	for {
		pos.x += vel.x
		pos.y += vel.y
		if pos.y > maxHeight {
			maxHeight = pos.y
		}
		if t.hit(pos) {
			return maxHeight
		}
		if t.missed(pos, vel) {
			return -1 //missed
		}
		vel.x--
		if vel.x < 0 {
			vel.x = 0
		}
		vel.y--
	}
}

type Target struct {
	X1 int `regroup:"x1"`
	X2 int `regroup:"x2"`
	Y1 int `regroup:"y1"`
	Y2 int `regroup:"y2"`
}

type pair struct {
	x, y int
}

func (t Target) hit(p pair) bool {
	return t.X1 <= p.x && p.x <= t.X2 && t.Y1 <= p.y && p.y <= t.Y2
}

func (t Target) missed(pos, velocity pair) bool {
	if velocity.x == 0 && (pos.x < t.X1 || pos.x > t.X2) {
		return true
	}
	return velocity.y < 0 && pos.y < t.Y1
}

var rex = regroup.MustCompile(`target area: x=(?P<x1>\d+)\.\.(?P<x2>\d+),\s+y=(?P<y1>[-\d]+)\.\.(?P<y2>[-\d]+)`)

func parseTarget(line string) Target {
	t := Target{}
	if err := rex.MatchToTarget(line, &t); err != nil {
		log.Fatal(err)
	}
	return t
}
