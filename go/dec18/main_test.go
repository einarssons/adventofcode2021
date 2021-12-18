package main

import (
	"fmt"
	"testing"

	u "github.com/einarssons/adventofcode2021/go/utils"
	"github.com/stretchr/testify/require"
)

func TestParseSnailFish(t *testing.T) {
	lines := u.ReadLinesFromFile("test.txt")
	for i, l := range lines {
		chars := u.SplitToChars(l)
		sf := parse(chars)
		fmt.Printf("%d, %+v\n", i, sf)
	}
}

func TestSum2(t *testing.T) {
	lines := u.ReadLinesFromFile("test2.txt")
	sf := parse(u.SplitToChars(lines[0]))
	for _, l := range lines[1:] {
		sf = add(sf, parse(u.SplitToChars(l)), true)
		_, _ = sf.reduce(0)
		fmt.Printf("%s\n", sf)
	}
	m := sf.magnitude()
	require.Equal(t, 4140, m)
}

func TestSum0(t *testing.T) {
	lines := u.ReadLinesFromFile("test0.txt")
	sf := parse(u.SplitToChars(lines[0]))
	for _, l := range lines[1:] {
		sf = add(sf, parse(u.SplitToChars(l)), true)
		_, _ = sf.reduce(0)
		fmt.Printf("%s\n", sf)
	}
	m := sf.magnitude()
	fmt.Printf("magnitude =%d\n", m)
}

func TestSum1(t *testing.T) {
	lines := u.ReadLinesFromFile("test1.txt")
	sf := parse(u.SplitToChars(lines[0]))
	for _, l := range lines[1:] {
		sf = add(sf, parse(u.SplitToChars(l)), true)
		_, _ = sf.reduce(0)
		fmt.Printf("%s\n", sf)
	}
	m := sf.magnitude()
	fmt.Printf("magnitude =%d\n", m)
}

func TestSum(t *testing.T) {
	line1 := "[[[[4,3],4],4],[7,[[8,4],9]]]"
	line2 := "[1,1]"
	sf1 := parse(u.SplitToChars(line1))
	sf2 := parse(u.SplitToChars(line2))
	sf := add(sf1, sf2, true)
	_, _ = sf.reduce(0)
	fmt.Printf("%v\n", sf)
}

func TestMagnitude(t *testing.T) {
	tcs := []struct {
		str string
		mag int
	}{
		{"[[1,2],[[3,4],5]]", 143},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
	}

	for _, tc := range tcs {
		sf := parse(u.SplitToChars(tc.str))
		gotM := sf.magnitude()
		require.Equal(t, tc.mag, gotM)
	}

}
