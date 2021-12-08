package main

import (
	"fmt"
	"sort"
	"strings"

	u "github.com/einarssons/adventofcode2021/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("test.txt")
	fmt.Printf("--- TASK 1 TEST: %d lines ---\n", len(lines))
	task1(lines)
	fmt.Printf("--- TASK 2 TEST: %d lines ---\n", len(lines))
	task2(lines)
	lines = u.ReadLinesFromFile("data.txt")
	fmt.Printf("### TASK 1 DATA: %d lines ###\n", len(lines))
	task1(lines)
	fmt.Printf("### TASK 2 DATA: %d lines ###\n", len(lines))
	task2(lines)
}

func task1(d []string) {
	n := 0
	for _, line := range d {
		_, output := splitLine(line)

		for _, o := range output {
			switch len(o) {
			case 2, 3, 4, 7:
				n++
			}
		}
	}
	fmt.Printf("Nr = %d\n", n)
}

func makeDigitMap() map[string]int {
	m := make(map[string]int)
	for i := 0; i <= 9; i++ {
		switch i {
		case 0:
			m["abcefg"] = 0
		case 1:
			m["cf"] = 1
		case 2:
			m["acdeg"] = 2
		case 3:
			m["acdfg"] = 3
		case 4:
			m["bcdf"] = 4
		case 5:
			m["abdfg"] = 5
		case 6:
			m["abdefg"] = 6
		case 7:
			m["acf"] = 7
		case 8:
			m["abcdefg"] = 8
		case 9:
			m["abcdfg"] = 9
		}
	}
	return m
}

const SEGS = "abcdefg"

func task2(d []string) {
	dm := makeDigitMap()
	sum := 0
	for _, line := range d {
		input, output := splitLine(line)
		//fmt.Println(input)

		var cf string
		var acf string
		var bcdf string
		freq := make(map[string]int)
		for _, c := range u.SplitToChars(SEGS) {
			freq[c] = 0
		}
		for _, ip := range input {
			chars := u.SplitToChars(ip)
			for _, c := range chars {
				freq[c]++
			}
			switch len(ip) {
			case 2:
				cf = ip
			case 3:
				acf = ip
			case 4:
				bcdf = ip
			}
		}
		charMap := make(map[string]string)
		cfm := segMap(cf)
		acfm := segMap(acf)
		// Key which is in acf but no cf must be a
		for k := range acfm {
			_, ok := cfm[k]
			if !ok {
				charMap[k] = "a"
			}
		}
		for k, v := range freq {
			switch v {
			case 4:
				charMap[k] = "e"
			case 9:
				charMap[k] = "f"
			case 6:
				charMap[k] = "b"
			case 8:
				if _, ok := charMap[k]; !ok {
					charMap[k] = "c"
				}
			case 7:
				// Can be d or g
				if !strings.Contains(bcdf, k) {
					charMap[k] = "g"
				} else {
					charMap[k] = "d"
				}
			}
		}
		//fmt.Printf("%v\n", charMap)
		nr := 0
		for _, o := range output {
			digit := digitFromCharMap(o, charMap, dm)
			nr = 10*nr + digit
		}
		sum += nr
	}
	fmt.Printf("sum = %d\n", sum)
}

func digitFromCharMap(input string, chm map[string]string, dm map[string]int) int {
	inps := u.SplitToChars(input)
	var ss []string
	for _, inp := range inps {
		ss = append(ss, chm[inp])
		sort.Strings(ss)
	}
	comb := strings.Join(ss, "")
	return dm[comb]
}

func segMap(s string) map[string]bool {
	cs := u.SplitToChars(s)
	m := make(map[string]bool)
	for _, c := range cs {
		m[c] = true
	}
	return m
}

func splitLine(line string) ([]string, []string) {
	parts := strings.Split(line, "|")
	input := strings.Fields(parts[0])
	output := strings.Fields(parts[1])
	return input, output
}
