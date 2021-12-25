package main

import (
	"fmt"
	"strings"
	"testing"

	u "github.com/einarssons/adventofcode2021/go/utils"
	"github.com/stretchr/testify/require"
)

var testData = `v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`

func TestTask1(t *testing.T) {
	lines := strings.Split(testData, "\n")
	floor := CreateCucumberGridFromLines(lines)
	fmt.Printf("%s\n", &floor)
	nrSteps := 0
	for {
		nrMoves := floor.Moov()
		nrSteps++
		fmt.Printf("steps: %d nrMoves %d\n", nrSteps, nrMoves)
		fmt.Printf("%s\n", &floor)
		if nrMoves == 0 {
			break
		}
	}
	require.Equal(t, 58, nrSteps)
}

func TestTask2(t *testing.T) {
	lines := u.ReadLinesFromFile("test.txt")
	task2(lines)
}
