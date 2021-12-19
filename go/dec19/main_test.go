package main

import (
	"fmt"
	"testing"

	u "github.com/einarssons/adventofcode2021/go/utils"
	"github.com/stretchr/testify/require"
)

func TestOverlapOneAndTwoWithKnownOrigin(t *testing.T) {
	lines := u.ReadLinesFromFile("test.txt")
	data := scanData(lines)
	scan0Data := data[0]
	scanner1Pos := coord{68, -1246, -43}
	goodRot := -1
	var overlap SetCoords
	for rotNr := 0; rotNr < 24; rotNr++ {
		rotatedData := rotateData(data[1], rotNr)
		displacedData := addOrigin(rotatedData, scanner1Pos)
		overlap = SetCoordsIntersection(scan0Data, displacedData)
		if len(overlap) >= 12 {
			goodRot = rotNr
			break
		}
	}
	require.Greater(t, goodRot, -1)
	for elem := range overlap {
		fmt.Printf("%v\n", elem)
	}
}

func TestOverlapOneAndTwoWithUnknownOrigin(t *testing.T) {
	lines := u.ReadLinesFromFile("test.txt")
	data := scanData(lines)
	scan0Data := data[0]
	scanner1Pos := coord{68, -1246, -43}
	goodRot := -1
	var goodOrigin coord
	for sc1Beacon := range scan0Data {
		var overlap SetCoords
		for rotNr := 0; rotNr < 24; rotNr++ {
			rotatedData := rotateData(data[1], rotNr)
			for rc := range rotatedData {
				possibleOrigin := coordDiff(sc1Beacon, rc)
				displacedData := addOrigin(rotatedData, possibleOrigin)
				overlap = SetCoordsIntersection(scan0Data, displacedData)
				if len(overlap) >= 12 {
					goodOrigin = possibleOrigin
					goodRot = rotNr
					break
				}
			}
		}
	}
	require.Greater(t, goodRot, -1)
	require.Equal(t, scanner1Pos, goodOrigin)
}

func TestAll(t *testing.T) {
	lines := u.ReadLinesFromFile("test.txt")
	beacons, _ := findBeaconsAndScanners(lines)
	require.Equal(t, 79, len(beacons))
}

func TestMaxDistance(t *testing.T) {
	lines := u.ReadLinesFromFile("test.txt")
	_, scanners := findBeaconsAndScanners(lines)
	maxDist := findMaxDist(scanners)
	require.Equal(t, 3621, maxDist)

}
