package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	target := Target{20, 30, -10, -5}

	testCases := []struct {
		vel       pair
		maxHeight int
	}{
		{vel: pair{7, 2}, maxHeight: 3},
		{vel: pair{6, 3}, maxHeight: 6},
		{vel: pair{17, -4}, maxHeight: -1},
		{vel: pair{6, 9}, maxHeight: 45},
	}

	for _, tc := range testCases {
		maxH := shoot(tc.vel, target)
		require.Equal(t, tc.maxHeight, maxH)
	}
}
