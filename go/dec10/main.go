package main

import (
	"fmt"
	"sort"

	u "github.com/einarssons/adventofcode2021/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("test.txt")
	fmt.Printf("Nr test lines: %d\n", len(lines))
	task1(lines)
	task2(lines)
	lines = u.ReadLinesFromFile("data.txt")
	fmt.Printf("Nr data lines: %d\n", len(lines))
	task1(lines)
	task2(lines)
}

func task1(d []string) {
	points := 0

	for _, l := range d {
		c, ok := check(u.SplitToChars(l))
		if !ok {
			//fmt.Printf("%d %s\n", i, c)
			switch c {
			case ")":
				points += 3
			case "]":
				points += 57
			case "}":
				points += 1197
			case ">":
				points += 25137
			}
		}
	}
	fmt.Printf("nr points %d\n", points)

}

func task2(d []string) {
	var scores []int
	for _, l := range d {
		st := getUncomplete(u.SplitToChars(l))
		if st == nil {
			continue
		}
		r := st.reverse()
		sc := 0
		for _, c := range r {
			sc *= 5
			switch c {
			case ")":
				sc += 1
			case "]":
				sc += 2
			case "}":
				sc += 3
			case ">":
				sc += 4
			}
		}
		scores = append(scores, sc)
	}
	sort.Ints(scores)
	mp := (len(scores) - 1) / 2
	mid := scores[mp]
	fmt.Printf("mid score %d\n", mid)
}

type StackStrings struct {
	s  []string
	nr int
}

func (s *StackStrings) Push(c string) {
	s.s = append(s.s, c)
	s.nr++
}

func (s *StackStrings) Pop() string {
	if s.nr == 0 {
		return ""
	}
	r := s.s[s.nr-1]
	s.nr--
	s.s = s.s[:s.nr]
	return r
}

func (s *StackStrings) IsEmpty() bool {
	return s.nr == 0
}

func (s *StackStrings) Depth() int {
	return s.nr
}

func (s *StackStrings) reverse() []string {
	r := make([]string, 0, len(s.s))
	for i := len(s.s) - 1; i >= 0; i-- {
		var m string
		switch s.s[i] {
		case "(":
			m = ")"
		case "{":
			m = "}"
		case "[":
			m = "]"
		case "<":
			m = ">"
		}
		r = append(r, m)
	}
	return r
}

func check(s []string) (string, bool) {
	st := StackStrings{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case "{", "[", "<", "(":
			st.Push(s[i])
		case "}":
			c := st.Pop()
			if c == "" || c != "{" {
				return s[i], false
			}
		case ")":
			c := st.Pop()
			if c == "" || c != "(" {
				return s[i], false
			}
		case ">":
			c := st.Pop()
			if c == "" || c != "<" {
				return s[i], false
			}
		case "]":
			c := st.Pop()
			if c == "" || c != "[" {
				return s[i], false
			}
		}
	}
	return "", true
}

func getUncomplete(s []string) *StackStrings {
	st := StackStrings{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case "{", "[", "<", "(":
			st.Push(s[i])
		case "}":
			c := st.Pop()
			if c == "" || c != "{" {
				return nil
			}
		case ")":
			c := st.Pop()
			if c == "" || c != "(" {
				return nil
			}
		case ">":
			c := st.Pop()
			if c == "" || c != "<" {
				return nil
			}
		case "]":
			c := st.Pop()
			if c == "" || c != "[" {
				return nil
			}
		}
	}
	return &st
}
