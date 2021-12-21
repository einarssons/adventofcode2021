package main

import (
	"fmt"
)

func main() {
	task1()
	task2()
}

func task1() {
	res := playGame([]int{2, 10})
	fmt.Printf("Task 1: Result = %d\n", res)
}

func task2() {
	win0, win1 := playGameDirac([]int{2, 10})
	if win0 > win1 {
		fmt.Printf("Result = %d\n", win0)
	} else {
		fmt.Printf("\nTask 2: Result = %d\n", win1)
	}
}

type Die struct {
	nrTurns int
}

func (d *Die) Roll() int {
	d.nrTurns++
	return (d.nrTurns-1)%100 + 1
}

type Player struct {
	pos   int
	score int
}

func (p *Player) Move(steps int) int {
	p.pos += steps
	nrTurns := (p.pos - 1) / 10
	p.pos = p.pos - nrTurns*10
	return p.pos
}

func playGame(start []int) int {
	die := Die{}
	players := make([]Player, len(start))
	for i := range players {
		players[i].pos = start[i]
	}
	winner := 0
	nrSteps := 0
game:
	for {
		for i := 0; i < 2; i++ {
			r := die.Roll() + die.Roll() + die.Roll()
			players[i].Move(r)
			players[i].score += players[i].pos
			nrSteps++
			//fmt.Printf("%d %d %d\n", nrSteps, i, players[i].score)
			if players[i].score >= 1000 {
				winner = i
				break game
			}
		}
	}
	loser := 1 - winner
	result := players[loser].score * die.nrTurns
	return result
}

type Universes struct {
	nextPlayer int
	univs      map[Universe]int
	wins0      int
	wins1      int
}

func (u *Universes) nrUniverses() int {
	tot := 0
	for _, v := range u.univs {
		tot += v
	}
	return tot
}

type Universe struct {
	pos0   int
	score0 int
	pos1   int
	score1 int
}

func createUniverses(start []int) *Universes {
	us := Universes{}
	us.univs = make(map[Universe]int)
	us.univs[Universe{pos0: start[0] - 1, pos1: start[1] - 1}] = 1
	return &us
}

func (u *Universes) Move3() {
	splits := []int{1, 3, 6, 7, 6, 3, 1}
	newUnivs := make(map[Universe]int)
	nrWins := 0
	if u.nextPlayer == 0 {
		for uv, oldNr := range u.univs {
			for i, split := range splits {
				finalPos := (uv.pos0 + i + 3) % 10 // zero-based
				nrNewUnivs := oldNr * split
				score := finalPos + 1
				newScore := uv.score0 + score
				if newScore >= 21 {
					nrWins += nrNewUnivs
				} else {
					newU := Universe{finalPos, newScore, uv.pos1, uv.score1}
					newUnivs[newU] += nrNewUnivs
				}
			}
		}
		u.wins0 += nrWins
	} else {
		for uv, oldNr := range u.univs {
			for i, split := range splits {
				finalPos := (uv.pos1 + i + 3) % 10 // zero-based
				nrNewUnivs := oldNr * split
				score := finalPos + 1
				newScore := uv.score1 + score
				if newScore >= 21 {
					nrWins += nrNewUnivs
				} else {
					newU := Universe{uv.pos0, uv.score0, finalPos, newScore}
					newUnivs[newU] += nrNewUnivs
				}
			}
		}
		u.wins1 += nrWins
	}
	u.univs = newUnivs
	u.nextPlayer = 1 - u.nextPlayer
}

func playGameDirac(start []int) (int, int) {
	un := createUniverses(start)
	nrSteps := 0
game:
	for {
		un.Move3()
		nrSteps++
		nrUniverses := un.nrUniverses()
		fmt.Printf("steps = %2d, wins1 = %14d, wins2 = %14d, universes = %13d\n", nrSteps, un.wins0, un.wins1, nrUniverses)
		if nrUniverses == 0 {
			break game
		}
	}
	return un.wins0, un.wins1
}
