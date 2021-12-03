package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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
		if trimmed == "" {
			continue
		}
		lines = append(lines, trimmed)
	}
	return lines
}
