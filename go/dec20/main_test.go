package main

import (
	"fmt"
	"testing"

	u "github.com/einarssons/adventofcode2021/go/utils"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	lines := u.ReadLinesFromFile("test.txt")
	iea, img := readData(lines)
	p := pattern(2, 2, img, ".")
	require.Equal(t, "...#...#.", p)
	n := pattern2nr(p)
	require.Equal(t, 34, n)
	newP := iea[n : n+1]
	require.Equal(t, "#", newP)
}

func TestExpand(t *testing.T) {
	lines := u.ReadLinesFromFile("test.txt")
	_, img := readData(lines)
	newImg := expand(img, ".")
	printImg(newImg)
}

func TestCount(t *testing.T) {
	lines := u.ReadLinesFromFile("test.txt")
	iea, img := readData(lines)
	printImg(img)
	fmt.Println()

	for i := 0; i < 2; i++ {
		img = expand(img, ".")
		printImg(img)
		fmt.Println()
		img = enhance(img, iea, ".")
		printImg(img)
		fmt.Println()
	}
	nrLit := countLit(img)
	require.Equal(t, 35, nrLit)
}

func TestCount50(t *testing.T) {
	lines := u.ReadLinesFromFile("test.txt")
	iea, img := readData(lines)

	for i := 0; i < 50; i++ {
		img = expand(img, ".")
		img = enhance(img, iea, ".")
	}
	nrLit := countLit(img)
	require.Equal(t, 3351, nrLit)
}
