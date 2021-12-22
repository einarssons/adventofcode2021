package main

import (
	"fmt"
	"strings"
	"testing"

	u "github.com/einarssons/adventofcode2021/go/utils"
	"github.com/stretchr/testify/require"
)

const testData = `on x=10..12,y=10..12,z=10..12
on x=11..13,y=11..13,z=11..13
off x=9..11,y=9..11,z=9..11
on x=10..10,y=10..10,z=10..10`

func TestSmall(t *testing.T) {
	lines := strings.Split(testData, "\n")
	c := createCuboid(21, 22, 23, 1, 2, 3)
	for _, line := range lines {
		step := parseLine(line)
		fmt.Printf("%v\n", step)
		c.step(step)
	}
	require.Equal(t, 39, c.nrOn())

}

func TestTask1(t *testing.T) {
	lines := u.ReadLinesFromFile("test.txt")
	c := createCuboid(101, 101, 101, -50, -50, -50)
	for _, line := range lines {
		step := parseLine(line)
		fmt.Printf("%v\n", step)
		c.step(step)
	}
	require.Equal(t, 590784, c.nrOn())
	//task1(lines)
}

func TestTask2Small(t *testing.T) {
	lines := strings.Split(testData, "\n")
	totVolume := task2(lines)
	require.Equal(t, 39, totVolume)

}

func TestTask2(t *testing.T) {
	lines := u.ReadLinesFromFile("test2.txt")
	totVolume := task2(lines)
	require.Equal(t, 2758514936282235, totVolume)
}
