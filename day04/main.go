package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func readWinners(leftPart string) (map[int]bool, int) {
	winners := make(map[int]bool)
	cardDesc := strings.Split(leftPart, ":")
	var id int
	fmt.Sscanf(cardDesc[0], "Card %d", &id)
	wins := strings.Split(cardDesc[1], " ")
	for _, swin := range wins {
		if swin != "" { //there are multiple spaces between numbers
			var win int
			fmt.Sscanf(swin, "%d", &win)
			winners[win] = true
		}
	}
	return winners, id
}

func readGuesses(rightPart string) (guesses map[int]bool) {
	guesses = make(map[int]bool)
	sguesses := strings.Split(rightPart, " ")
	for _, sguess := range sguesses {
		if sguess != "" { //there are multiple spaces between numbers
			var guess int
			fmt.Sscanf(sguess, "%d", &guess)
			guesses[guess] = true
		}
	}
	return guesses
}

type card struct {
	id      int
	winners map[int]bool
	guesses map[int]bool
	matches int
}

// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
func solve(input string) (worth, linescount int) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var cards []*card
	cardindex := make(map[int]*card)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		winners, id := readWinners(parts[0])
		guesses := readGuesses(parts[1])
		cardvalue := 0
		sumOfMatches := 0
		for guess := range guesses {
			if _, ok := winners[guess]; ok {
				sumOfMatches++
			}
		}
		if sumOfMatches > 0 { //task 1
			cardvalue = int(math.Pow(2, float64(sumOfMatches-1)))
			worth += int(cardvalue)
		}
		//prep task2
		card := &card{id, winners, guesses, sumOfMatches}
		cards = append(cards, card)
		cardindex[id] = card
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	// task2
	for i := 0; i < len(cards); i++ {
		for m := 1; m <= cards[i].matches; m++ {
			x := cardindex[cards[i].id+m]
			if x != nil {
				newcards := append(cards[:i+m], append([]*card{x}, cards[i+m:]...)...)
				cards = newcards
			}
		}
	}
	return worth, len(cards)
}

func main() {
	input := "input.txt"
	t1, t2 := solve(input)
	fmt.Println("Task 1 - summed wort of cards  \t =  ", t1)
	fmt.Println("Task 2 - total number of cards \t =  ", t2)
}
