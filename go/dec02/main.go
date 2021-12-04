package main

import (
	"fmt"

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
		cmd := utils.ParseCommand(l)
		switch cmd.Verb {
		case "forward":
			pos += cmd.Value
			depth += aim * cmd.Value
		case "up":
			aim -= cmd.Value
			if aim < 0 {
				aim = 0
			}
		case "down":
			aim += cmd.Value
		}
	}
	fmt.Printf("Task2: %d\n", depth*pos)
}

func task1(lines []string) {
	pos := 0
	depth := 0
	for _, l := range lines {
		cmd := utils.ParseCommand(l)
		switch cmd.Verb {
		case "forward":
			pos += cmd.Value
		case "up":
			depth -= cmd.Value
			if depth < 0 {
				depth = 0
			}
		case "down":
			depth += cmd.Value
		}
	}
	fmt.Printf("Task1: %d\n", depth*pos)
}
