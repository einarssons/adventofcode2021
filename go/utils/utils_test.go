package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitToInts(t *testing.T) {
	testCases := []struct {
		name            string
		line            string
		expectedNumbers []int
	}{
		{"commma-list", "1,2,3", []int{1, 2, 3}},
		{"empty-list", "", nil},
		{"space-list", "1 2   4", []int{1, 2, 4}},
	}

	for _, tc := range testCases {
		gotNumbers := SplitToInts(tc.line)
		require.Equal(t, tc.expectedNumbers, gotNumbers, tc.name)
	}
}

func TestParseCommand(t *testing.T) {
	testCases := []struct {
		name            string
		line            string
		expectedCommand Command
	}{
		{"up2", "UP 2", Command{"UP", 2}},
		{"forward3", "FORWARD   3", Command{"FORWARD", 3}},
	}

	for _, tc := range testCases {
		gotCommand := ParseCommand(tc.line)
		require.Equal(t, tc.expectedCommand, gotCommand, tc.name)
	}
}

func TestSplitToChars(t *testing.T) {
	testCases := []struct {
		name     string
		line     string
		expected []string
	}{
		{"two", "ab", []string{"a", "b"}},
		{"zero", "", []string{}},
	}

	for _, tc := range testCases {
		got := SplitToChars(tc.line)
		require.Equal(t, tc.expected, got, tc.name)
	}
}
