package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func readWinners(leftPart string) (map[int]int, int) {
	winners := make(map[int]int)
	cardDesc := strings.Split(strings.Trim(leftPart, " "), ":")
	var id int
	fmt.Sscanf(strings.Trim(cardDesc[0], " "), "Card %d", &id)
	wins := strings.Split(strings.Trim(cardDesc[1], " "), " ")
	for _, swin := range wins {
		if swin != "" { //there are multiple spaces between numbers
			var win int
			fmt.Sscanf(swin, "%d", &win)
			count := winners[win]
			winners[win] = count + 1
		}
	}
	return winners, id
}

func readGuesses(rightPart string) (guesses map[int]int) {
	guesses = make(map[int]int)
	sguesses := strings.Split(strings.Trim(rightPart, " "), " ")
	for _, sguess := range sguesses {
		if sguess != "" { //there are multiple spaces between numbers
			var guess int
			fmt.Sscanf(sguess, "%d", &guess)
			count := guesses[guess]
			guesses[guess] = count + 1
		}
	}
	return guesses
}

// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
func solve(input string) (worth int) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		winners, id := readWinners(parts[0])
		guesses := readGuesses(parts[1])
		sum := 0
		cardvalue := 0
		for guess := range guesses {
			if _, ok := winners[guess]; ok {
				sum++
			}
		}
		if sum > 0 {
			cardvalue = int(math.Pow(2, float64(sum-1)))
			worth += int(cardvalue)
		}
		fmt.Printf("Card %d has %d wins, is worth %d. Total is %d", id, sum, cardvalue, worth)
		fmt.Printf("  \tMultiple winners: ")
		for win, count := range winners {
			if count > 1 {
				fmt.Printf(" %d:%d", win, count)
			}
		}
		fmt.Printf("\tMultiple guesses: ")
		for guess, count := range guesses {
			if count > 1 {
				fmt.Printf(" %d:%d", guess, count)
			}
		}
		fmt.Println()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return worth
}

func main() {

	// const MaxUint = ^uint(0)
	// const MaxInt = int(MaxUint >> 1)

	// fmt.Printf("MaxUint: %d\n", MaxUint)
	// fmt.Printf("MaxInt: %d\n", MaxInt)

	input := "input.txt"

	fmt.Println("Task 1 - sum of digits                     \t =  ", solve(input))
	//fmt.Println("Task 2 - sum of digits and written numbers \t =  ", solve(input))
}
