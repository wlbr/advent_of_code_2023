package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type set struct {
	red   int
	green int
	blue  int
}

type game struct {
	id       int
	sets     []*set
	possible bool
	minred   int
	mingreen int
	minblue  int
}

type match struct {
	games    []*game
	maxred   int
	maxgreen int
	maxblue  int
}

func (m *match) String() (s string) {
	for _, g := range m.games {
		s += fmt.Sprintf("Game %d: ", g.id)
		for _, set := range g.sets {
			s += fmt.Sprintf("%d red, %d green, %d blue;", set.red, set.green, set.blue)
		}
		s += fmt.Sprintf(" %t\n", g.possible)
	}
	return s
}

func (g *game) setPossible(maxred, maxgreen, maxblue int) {
	for _, s := range g.sets {
		if s.red > maxred || s.green > maxgreen || s.blue > maxblue {
			g.possible = false
			break
		}
	}
}

func (g *game) setMax() {
	for _, s := range g.sets {
		if s.red > g.minred {
			g.minred = s.red
		}
		if s.green > g.mingreen {
			g.mingreen = s.green
		}
		if s.blue > g.minblue {
			g.minblue = s.blue
		}
	}
}

func (g *game) getPowerOfGame() (power int) {
	return g.minred * g.mingreen * g.minblue
}

func (m *match) setGamesPossible() {
	for _, g := range m.games {
		g.setPossible(m.maxred, m.maxgreen, m.maxblue)
	}
}

func (m *match) getScore() (sumOfPossibleGamesscore, sumOfPowers int) {
	for _, g := range m.games {
		if g.possible {
			sumOfPossibleGamesscore += g.id
		}
		sumOfPowers += g.getPowerOfGame()
	}
	return sumOfPossibleGamesscore, sumOfPowers
}

func (m *match) readData(input string) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		//Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		g := &game{possible: true}
		colonsplit := strings.Split(line, ":")
		fmt.Sscanf(colonsplit[0], "Game %d", &g.id)
		semicolonsplit := strings.Split(colonsplit[1], ";")
		for _, s := range semicolonsplit {
			commasplit := strings.Split(s, ",")
			set := &set{}
			for _, c := range commasplit {
				var color string
				var count int
				fmt.Sscanf(c, "%d %s", &count, &color)
				switch color {
				case "red":
					set.red = count
				case "green":
					set.green = count
				case "blue":
					set.blue = count
				}
			}
			g.sets = append(g.sets, set)
		}
		g.setPossible(m.maxred, m.maxgreen, m.maxblue)
		g.setMax()
		m.games = append(m.games, g)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func solve(input string, maxred, maxgreen, maxblue int) (sum, power int) {
	m := &match{maxred: maxred, maxgreen: maxgreen, maxblue: maxblue}
	m.readData(input)
	sum, power = m.getScore()
	return sum, power
}

func main() {
	input := "input.txt"

	t1, t2 := solve(input, 12, 13, 14)
	fmt.Println("Task 1 - sum of IDs of possible games \t =  ", t1)
	fmt.Println("Task 2 - sum of powers of cubes       \t =  ", t2)
}
