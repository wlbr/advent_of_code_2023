package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

func readInput(input string) (left []int, right []int) {
	f, ferr := os.Open(input)
	if ferr != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, ferr)
	}
	defer f.Close()

	var l, r int = 0, 0
	var err error
	for err == nil {
		if _, err = fmt.Fscanf(f, "%d %d\n", &l, &r); err == nil {
			left = append(left, l)
			right = append(right, r)
			l, r = 0, 0
		}
	}

	return left, right
}

func abs(x int) (a int) {
	if x < 0 {
		a = -x
	} else {
		a = x
	}

	return a
}

func solve1(input string) (sum int) {
	left, right := readInput(input)
	sort.Ints(left)
	sort.Ints(right)

	for i, l := range left {
		r := right[i]
		sum += abs(l - r)
	}

	return sum
}

func solve2(input string) (sum int) {
	left, right := readInput(input)
	sort.Ints(left)
	sort.Ints(right)

	var frequencies map[int]int = make(map[int]int)

	for _, r := range right {
		frequencies[r]++
	}

	for _, l := range left {
		sum += l * frequencies[l]
	}
	return sum
}

func main() {
	input := "input.txt"

	fmt.Println("Task 1 - Total distance    \t =  ", solve1(input))
	fmt.Println("Task 2 - Similarity score \t =  ", solve2(input))
}
