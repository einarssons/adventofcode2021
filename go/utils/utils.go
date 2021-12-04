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
