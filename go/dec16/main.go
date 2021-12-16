package main

import (
	"fmt"

	u "github.com/einarssons/adventofcode2021/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("data.txt")
	fmt.Printf("### TASK 1 MY DATA: %d lines ###\n", len(lines))
	task1(lines)
	fmt.Printf("### TASK 2 MY DATA: %d lines ###\n", len(lines))
	task2(lines)
}

func task1(lines []string) {
	sumVersions, _ := parse(lines[0])
	fmt.Printf("sum versions = %d\n", sumVersions)
}

func parse(d string) (sumVersions int, values []int) {
	br := createBR(d)
	sumVersions, values = br.parse(0, 1)
	if !br.onlyZerosLeft() {
		panic("BAD")
	}
	return sumVersions, values
}

func task2(lines []string) {
	_, values := parse(lines[0])
	if len(values) != 1 {
		panic("BAD")
	}
	fmt.Printf("result = %d\n", values[0])
}

type bitReader struct {
	data []int
	pos  int
}

func (b *bitReader) read(n int) int {
	nr := 0
	for i := b.pos; i < b.pos+n; i++ {
		nr = nr<<1 + b.data[i]
	}
	b.pos += n
	return nr
}

func createBR(hexString string) *bitReader {
	br := bitReader{
		data: readHex(hexString),
		pos:  0,
	}
	return &br
}

// parse - return sum of version
func (b *bitReader) parse(nrBits, nrPackets int) (sumVersions int, values []int) {
	nrPacketsRead := 0
	bitPos := b.pos
	var value int
	for {
		pVer := b.read(3)
		sumVersions += pVer
		pType := b.read(3)
		if pType == 4 {
			value = b.parseNumber()
		} else {
			lVer := b.read(1)
			var subValues []int
			var sumVer int
			if lVer == 0 {
				subNrBits := b.read(15)
				sumVer, subValues = b.parse(subNrBits, -1)
				sumVersions += sumVer
			} else {
				subNrPackets := b.read(11)
				sumVer, subValues = b.parse(-1, subNrPackets)
				sumVersions += sumVer
			}
			switch pType {
			case 0: // SUM
				value = 0
				for _, sp := range subValues {
					value += sp
				}
			case 1: // Prod
				value = 1
				for _, sp := range subValues {
					value *= sp
				}
			case 2: // Min
				value = 1 << 50
				for _, sp := range subValues {
					if sp < value {
						value = sp
					}
				}
			case 3: // Max
				value = 0
				for _, sp := range subValues {
					if sp > value {
						value = sp
					}
				}
			case 5: // >
				value = 0
				if subValues[0] > subValues[1] {
					value = 1
				}
			case 6: // <
				value = 0
				if subValues[0] < subValues[1] {
					value = 1
				}
			case 7: // =
				value = 0
				if subValues[0] == subValues[1] {
					value = 1
				}
			}
		}
		values = append(values, value)
		nrPacketsRead++
		nrBitsConsumed := b.pos - bitPos
		if nrBits > 0 && nrBitsConsumed == nrBits {
			break
		}
		if nrPackets > 0 && nrPacketsRead == nrPackets {
			break
		}
	}
	return sumVersions, values
}

func (b *bitReader) onlyZerosLeft() bool {
	for p := b.pos; p < len(b.data); p++ {
		if b.data[p] == 1 {
			return false
		}
	}
	return true
}

func (b *bitReader) parseNumber() int {
	nr := 0
	for {
		lastBit := b.read(1)
		bits := b.read(4)
		nr = nr<<4 + bits
		if lastBit == 0 {
			break
		}
	}
	return nr
}

func (b *bitReader) parseLiteral() int {
	pVer := b.read(3)
	fmt.Println(pVer)
	pType := b.read(3)
	fmt.Println(pType)
	return b.parseNumber()
}

func readHex(l string) []int {
	hexD := u.SplitToChars(l)
	bits := make([]int, 0, 4*len(hexD))
	for _, h := range hexD {
		switch h {
		case "0":
			bits = append(bits, []int{0, 0, 0, 0}...)
		case "1":
			bits = append(bits, []int{0, 0, 0, 1}...)
		case "2":
			bits = append(bits, []int{0, 0, 1, 0}...)
		case "3":
			bits = append(bits, []int{0, 0, 1, 1}...)
		case "4":
			bits = append(bits, []int{0, 1, 0, 0}...)
		case "5":
			bits = append(bits, []int{0, 1, 0, 1}...)
		case "6":
			bits = append(bits, []int{0, 1, 1, 0}...)
		case "7":
			bits = append(bits, []int{0, 1, 1, 1}...)
		case "8":
			bits = append(bits, []int{1, 0, 0, 0}...)
		case "9":
			bits = append(bits, []int{1, 0, 0, 1}...)
		case "A":
			bits = append(bits, []int{1, 0, 1, 0}...)
		case "B":
			bits = append(bits, []int{1, 0, 1, 1}...)
		case "C":
			bits = append(bits, []int{1, 1, 0, 0}...)
		case "D":
			bits = append(bits, []int{1, 1, 0, 1}...)
		case "E":
			bits = append(bits, []int{1, 1, 1, 0}...)
		case "F":
			bits = append(bits, []int{1, 1, 1, 1}...)
		}
	}
	return bits
}
