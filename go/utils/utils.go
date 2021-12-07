package utils

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/oriser/regroup"
)

func ReadNumbersFromFile(path string) []int {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	var numbers []int
	s := bufio.NewScanner(f)
	for s.Scan() {
		txt := s.Text()
		nr, err := strconv.Atoi(txt)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, nr)
	}
	return numbers
}

func ReadLinesFromFile(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	var lines []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		trimmed := strings.Trim(line, " ")
		lines = append(lines, trimmed)
	}
	return lines
}

func SplitToInts(line string) []int {
	re := regexp.MustCompile("[0-9]+")
	parts := re.FindAllString(line, -1)
	var numbers []int
	for _, p := range parts {
		number, err := strconv.Atoi(p)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func SplitToChars(line string) []string {
	chars := make([]string, len(line))
	for i := 0; i < len(line); i++ {
		chars[i] = line[i : i+1]
	}
	return chars
}

func ContainsInt(x int, entries []int) bool {
	for _, n := range entries {
		if x == n {
			return true
		}
	}
	return false
}

func ContainsString(x string, entries []string) bool {
	for _, n := range entries {
		if x == n {
			return true
		}
	}
	return false
}

var rex = regroup.MustCompile(`(?P<verb>[a-zA-Z]+)\s+(?P<value>\d+)`)

type Command struct {
	Verb  string `regroup:"verb"`
	Value int    `regroup:"value"`
}

func ParseCommand(line string) Command {
	c := Command{}
	if err := rex.MatchToTarget(line, &c); err != nil {
		log.Fatal(err)
	}
	return c
}

func MinMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func MinMaxInts(numbers []int) (int, int) {
	min, max := numbers[0], numbers[0]
	for _, nr := range numbers {
		if nr > max {
			max = nr
		}
		if nr < min {
			min = nr
		}
	}
	return min, max
}

func Min(numbers []int) int {
	minNr := 1 << 40
	for _, nr := range numbers {
		if nr < minNr {
			minNr = nr
		}
	}
	return minNr
}

func Max(numbers []int) int {
	maxNr := -(1 << 40)
	for _, nr := range numbers {
		if nr > maxNr {
			maxNr = nr
		}
	}
	return maxNr
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Triangle(nr int) int {
	return nr * (nr + 1) / 2
}
