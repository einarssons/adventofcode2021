package main

import (
	"fmt"
	"strings"

	u "github.com/einarssons/adventofcode2021/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("data.txt")
	task(lines)
}

func task(lines []string) {
	beacons, scanners := findBeaconsAndScanners(lines)
	fmt.Printf("nr beacons = %d\n", len(beacons))
	maxDist := findMaxDist(scanners)
	fmt.Printf("max distance = %d\n", maxDist)
}

func findBeaconsAndScanners(lines []string) (SetCoords, []scanner) {
	scannedData := scanData(lines)
	scanners := make([]scanner, len(scannedData))
	scanners[0].calibrated = true
	allBeacons := CreateSetCoords()
	for c := range scannedData[0] {
		allBeacons.Add(c)
	}
	nrCalibrated := 1
	for {
		if nrCalibrated == len(scanners) {
			break
		}
		for sc_nr := 1; sc_nr < len(scanners); sc_nr++ {
			if scanners[sc_nr].calibrated {
				continue
			}
		loopRefs:
			for sc_ref := 0; sc_ref < len(scanners); sc_ref++ {
				if sc_ref == sc_nr {
					continue
				}
				if !scanners[sc_ref].calibrated {
					continue
				}
				goodRotNr := -1
				var goodOrigin coord
			compareRef:
				for refBeacon := range scannedData[sc_ref] {
					var overlap SetCoords
					for rotNr := 0; rotNr < 24; rotNr++ {
						rotatedData := rotateData(scannedData[sc_nr], rotNr)
						for rc := range rotatedData {
							possibleOrigin := coordDiff(refBeacon, rc)
							displacedData := addOrigin(rotatedData, possibleOrigin)
							overlap = SetCoordsIntersection(scannedData[sc_ref], displacedData)
							if len(overlap) >= 12 {
								goodOrigin = possibleOrigin
								goodRotNr = rotNr
								scannedData[sc_nr] = displacedData
								allBeacons.Extend(displacedData)
								break compareRef
							}
						}
					}
				}
				fmt.Printf("Nr beacons found: %4d\n", len(allBeacons))
				if goodRotNr >= 0 {
					scanners[sc_nr].pos = goodOrigin
					scanners[sc_nr].rot = goodRotNr
					scanners[sc_nr].calibrated = true
					nrCalibrated++
					break loopRefs
				}
			}
		}
	}
	return allBeacons, scanners
}

func scanData(lines []string) []SetCoords {
	var scans []SetCoords
	scan := CreateSetCoords()
	for _, l := range lines {
		if strings.HasPrefix(l, "---") {
			if len(scan) > 0 {
				scans = append(scans, scan)
			}
			scan = CreateSetCoords()
			continue
		}
		if l == "" {
			continue
		}
		c := u.SplitToInts(l)
		scan.Add(coord{c[0], c[1], c[2]})
	}
	scans = append(scans, scan)
	return scans
}

type scanner struct {
	calibrated bool
	pos        coord
	rot        int
}

func rotateData(d SetCoords, rot int) SetCoords {
	r := CreateSetCoords()
	for c := range d {
		r.Add(c.rotate(rot))
	}
	return r
}

func addOrigin(d SetCoords, orig coord) SetCoords {
	r := CreateSetCoords()
	for c := range d {
		r.Add(coord{c.x + orig.x, c.y + orig.y, c.z + orig.z})
	}
	return r
}
