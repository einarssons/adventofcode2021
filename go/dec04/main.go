package main

import (
	"fmt"

	"github.com/einarssons/adventofcode2021/go/utils"
)

func main() {
	lines := utils.ReadLinesFromFile("test.txt")
	fmt.Printf("Nr test lines: %d\n", len(lines))
	task1(lines)
	task2(lines)
	lines = utils.ReadLinesFromFile("data.txt")
	fmt.Printf("Nr data lines: %d\n", len(lines))
	task1(lines)
	task2(lines)
}

func task1(lines []string) {
	numbers, boards := parseData(lines)
	//fmt.Printf("numbers: %v\n", numbers)
	//fmt.Printf("boards: %v\n", boards)
outer:
	for i := 5; i <= len(numbers); i++ {
		seq := numbers[:i]
		for nr, board := range boards {
			if board.fullRow(seq) {
				sum := board.sumNonMatching(seq)
				fmt.Printf("First board nr %d at %d prod %d \n", nr, numbers[i-1], sum*numbers[i-1])
				break outer
			}
		}
	}

}

func task2(lines []string) {
	numbers, boards := parseData(lines)
	//fmt.Printf("numbers: %v\n", numbers)
	//fmt.Printf("boards: %v\n", boards)
	won := make([]bool, len(boards))
	nrWon := 0
outer:
	for i := 5; i <= len(numbers); i++ {
		seq := numbers[:i]
		for nr, board := range boards {
			if won[nr] {
				continue
			}
			if board.fullRow(seq) {
				won[nr] = true
				nrWon++
				if nrWon < len(boards) {
					continue
				}
				sum := board.sumNonMatching(seq)
				fmt.Printf("Last board nr %d at %d prod %d \n", nr, numbers[i-1], sum*numbers[i-1])
				break outer
			}
		}
	}
}

type Board []int

const (
	w = 5
)

func (b Board) sumNonMatching(numbers []int) int {
	sum := 0
	for _, boardNr := range b {
		if !utils.ContainsInt(boardNr, numbers) {
			sum += boardNr
		}
	}
	return sum
}

func (b Board) fullRow(numbers []int) bool {
	var all bool
	for r := 0; r < w; r++ {
		all = true
		for c := 0; c < w; c++ {
			if !utils.ContainsInt(b[r*w+c], numbers) {
				all = false
				break
			}
		}
		if all {
			return true
		}
	}
	for c := 0; c < w; c++ {
		all = true
		for r := 0; r < w; r++ {
			if !utils.ContainsInt(b[r*w+c], numbers) {
				all = false
				break
			}
		}
		if all {
			return true
		}
	}
	return false
}

func parseData(lines []string) (numbers []int, boards []Board) {
	var board Board
	for i, line := range lines {
		if i == 0 {
			numbers = utils.SplitToInts(line)
			continue
		}
		if line == "" {
			if len(board) > 0 {
				boards = append(boards, board)
				board = Board{}
			}
			continue
		}
		row := utils.SplitToInts(line)
		board = append(board, row...)
	}
	if len(board) != 0 {
		boards = append(boards, board)
	}
	return numbers, boards
}
