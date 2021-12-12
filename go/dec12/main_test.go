package main

import (
	"testing"

	u "github.com/einarssons/adventofcode2021/go/utils"
	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	lines := u.ReadLinesFromFile("test.txt")
	expectedNrPaths := 226
	allowOneDoubleLow := false
	gotNrPaths := task(lines, allowOneDoubleLow)
	require.Equal(t, expectedNrPaths, gotNrPaths)
}

func TestTask2(t *testing.T) {
	lines := u.ReadLinesFromFile("test.txt")
	expectedNrPaths := 3509
	allowOneDoubleLow := true
	gotNrPaths := task(lines, allowOneDoubleLow)
	require.Equal(t, expectedNrPaths, gotNrPaths)
}
