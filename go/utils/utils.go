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

// GCDuint64 - greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		b, a = a%b, b
	}
	return a
}

func ReverseStrings(s []string) {
	l := len(s)
	nrSwaps := l / 2
	for i := 0; i < nrSwaps; i++ {
		s[i], s[l-1-i] = s[l-1-i], s[i]
	}
}

// Stack of strings
type StackStrings struct {
	elems []string
	nr    int
}

func (s *StackStrings) Push(elem string) {
	s.elems = append(s.elems, elem)
	s.nr++
}

// Pop - get element if available as signaled by ok
func (s *StackStrings) Pop() (elem string, ok bool) {
	if s.nr == 0 {
		return "", false
	}
	elem = s.elems[s.nr-1]
	s.nr--
	s.elems = s.elems[:s.nr]
	return elem, true
}

func (s *StackStrings) IsEmpty() bool {
	return s.nr == 0
}

func (s *StackStrings) Depth() int {
	return s.nr
}

// Stack of ints
type StackInts struct {
	elems []int
	nr    int
}

func (s *StackInts) Push(elem int) {
	s.elems = append(s.elems, elem)
	s.nr++
}

// Pop - get element if available as signaled by ok
func (s *StackInts) Pop() (elem int, ok bool) {
	if s.nr == 0 {
		return 0, false
	}
	elem = s.elems[s.nr-1]
	s.nr--
	s.elems = s.elems[:s.nr]
	return elem, true
}

func (s *StackInts) IsEmpty() bool {
	return s.nr == 0
}

func (s *StackInts) Depth() int {
	return s.nr
}

type DigitGrid struct {
	Grid   [][]int
	Width  int
	Height int
}

func CreateDigitGridFromLines(lines []string) DigitGrid {
	g := DigitGrid{}
	for i, line := range lines {
		if i == 0 {
			g.Width = len(line)
		}
		if len(line) != g.Width {
			panic("non-rectangular grid")
		}
		row := make([]int, 0, g.Width)
		digits := SplitToChars(line)
		for _, digit := range digits {
			nr, err := strconv.Atoi(digit)
			if err != nil {
				panic(err)
			}
			row = append(row, nr)
		}
		g.Grid = append(g.Grid, row)
		g.Height++
	}
	return g
}

// InBounds - is (i, j) in grid
func (g DigitGrid) InBounds(i, j int) bool {
	return 0 <= i && i < g.Height && 0 <= j && j < g.Width
}
