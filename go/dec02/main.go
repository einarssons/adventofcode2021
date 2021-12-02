package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/einarssons/adventofcode2021/go/utils"
)

func main() {
	lines := utils.ReadLinesFromFile("data.txt")
	task1(lines)
	task2(lines)
}

func task2(lines []string) {
	pos := 0
	depth := 0
	aim := 0
	for _, l := range lines {
		cmd := getCommand(l)
		switch cmd.cmd {
		case "forward":
			pos += cmd.steps
			depth += aim * cmd.steps
		case "up":
			aim -= cmd.steps
			if aim < 0 {
				aim = 0
			}
		case "down":
			aim += cmd.steps
		}
	}
	fmt.Printf("Task2: %d\n", depth*pos)
}

func task1(lines []string) {
	pos := 0
	depth := 0
	for _, l := range lines {
		cmd := getCommand(l)
		switch cmd.cmd {
		case "forward":
			pos += cmd.steps
		case "up":
			depth -= cmd.steps
			if depth < 0 {
				depth = 0
			}
		case "down":
			depth += cmd.steps
		}
	}
	fmt.Printf("Task1: %d\n", depth*pos)
}

type command struct {
	cmd   string
	steps int
}

func getCommand(line string) command {
	parts := strings.Split(line, " ")
	cmd := command{parts[0], 0}
	nr, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	cmd.steps = nr
	return cmd
}
