package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	res := playGame([]int{4, 8})
	require.Equal(t, 739785, res)
}

func TestTask2(t *testing.T) {
	win0, win1 := playGameDirac([]int{4, 8})
	require.Equal(t, 444356092776315, win0)
	require.Equal(t, 341960390180808, win1)
}
