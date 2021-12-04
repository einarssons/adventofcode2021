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
