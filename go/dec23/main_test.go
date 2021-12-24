package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var testData = `#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`

var testData2 = `#############
#...........#
###B#C#B#D###
  #D#C#B#A#
  #D#B#A#C#
  #A#D#C#A#
  #########`

func Test1WalkExample(t *testing.T) {
	lines := strings.Split(testData, "\n")
	apods := parse(lines)
	c := cave{apods, 0, 2}
	fmt.Printf("%s\n\n", c.String())
	c.moveApod(2, 3, 0)
	fmt.Printf("%s\n\n", c.String())
	c.moveApod(1, 6, 1)
	fmt.Printf("%s\n\n", c.String())
	c.moveApod(5, 5, 0)
	c.moveApod(2, 4, 2)
	c.moveApod(0, 4, 1)
	c.moveApod(3, 7, 0)
	c.moveApod(7, 9, 0)
	c.moveApod(3, 8, 2)
	c.moveApod(5, 8, 1)
	c.moveApod(7, 2, 1)
	fmt.Printf("%s\n\n", c.String())
	require.Equal(t, 12521, c.cost)
}

func Test1Solve(t *testing.T) {
	lines := strings.Split(testData, "\n")
	apods := parse(lines)
	c := cave{apods, 0, 2}
	lowestCost := play1(c, 0, 15000, 0)
	require.Equal(t, 12521, lowestCost)
}

func Test2Solve(t *testing.T) {
	lines := strings.Split(testData2, "\n")
	apods := parse(lines)
	c := cave{apods, 0, 4}
	fmt.Printf("%s\n", &c)
	lowestCost := play1(c, 0, 50000, 0)
	require.Equal(t, 44169, lowestCost)
}

var testDataPartial = `#############
#AA.......AD#
###.#B#C#.###
  #.#B#C#.#
  #D#B#C#D#
  #A#B#C#D#
  #########`

func Test2Partial(t *testing.T) {
	lines := strings.Split(testDataPartial, "\n")
	apods := parse(lines)
	c := cave{apods, 0, 4}
	fmt.Printf("%s\n", &c)
	play1(c, 0, 50000, 0)
}
