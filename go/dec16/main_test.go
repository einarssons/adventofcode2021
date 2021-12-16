package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseLiteral(t *testing.T) {
	input := "D2FE28"
	br := createBR(input)
	pVer := br.read(3)
	fmt.Println(pVer)
	pType := br.read(3)
	fmt.Println(pType)
	nr := br.parseNumber()
	require.Equal(t, 2021, nr)
}

func TestParseOperator(t *testing.T) {
	input := "38006F45291200"
	br := createBR(input)
	pVer := br.read(3)
	require.Equal(t, 1, pVer)
	pType := br.read(3)
	require.Equal(t, 6, pType)
	lType := br.read(1)
	require.Equal(t, 0, lType)
	spLength := br.read(15)
	require.Equal(t, 27, spLength)
	nr := br.parseLiteral()
	require.Equal(t, 10, nr)
	nr = br.parseLiteral()
	require.Equal(t, 20, nr)
}

func TestParseOperator2(t *testing.T) {
	input := "EE00D40C823060"
	br := createBR(input)
	pVer := br.read(3)
	require.Equal(t, 7, pVer)
	pType := br.read(3)
	require.Equal(t, 3, pType)
	lType := br.read(1)
	require.Equal(t, 1, lType)
	spNr := br.read(11)
	require.Equal(t, 3, spNr)
	nr := br.parseLiteral()
	require.Equal(t, 1, nr)
	nr = br.parseLiteral()
	require.Equal(t, 2, nr)
	nr = br.parseLiteral()
	require.Equal(t, 3, nr)
}

func TestParsing(t *testing.T) {
	testCases := []struct {
		hex  string
		vSum int
	}{
		{"8A004A801A8002F478", 16},
		{"620080001611562C8802118E34", 12},
		{"C0015000016115A2E0802F182340", 23},
		{"A0016C880162017C3686B18A3D4780", 31},
	}
	for _, tc := range testCases {
		gotSum, _ := parse(tc.hex)
		require.Equal(t, tc.vSum, gotSum)
	}
}

func TestCase2(t *testing.T) {
	testCases := []struct {
		hex    string
		result int
	}{
		{"C200B40A82", 3},
		{"04005AC33890", 54},
		{"880086C3E88112", 7},
		{"CE00C43D881120", 9},
		{"D8005AC2A8F0", 1},
		{"F600BC2D8F", 0},
		{"9C005AC2F8F0", 0},
		{"9C0141080250320F1802104A08", 1},
	}
	for _, tc := range testCases {
		_, gotValues := parse(tc.hex)
		require.Equal(t, tc.result, gotValues[0])
	}
}
