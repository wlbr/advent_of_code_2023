package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var numbers []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func getLineSum(line string, includeWords bool) (num int) {
	index := len(line)
	for pos, c := range line {
		if n, err := strconv.Atoi(string(c)); err == nil {
			num = 10 * n
			index = pos
			break
		}
	}
	if includeWords {
		for i, n := range numbers {
			pos := strings.Index(line, n)
			if pos >= 0 && pos <= index {
				num = 10 * (i + 1)
				index = pos
			}
		}
	}

	index = -1
	n := -1
	for pos := len(line) - 1; pos >= 0; pos-- {
		if t, err := strconv.Atoi(string(line[pos])); err == nil {
			n = t
			index = pos
			break
		}
	}
	if includeWords {
		for i, sn := range numbers {
			pos := strings.LastIndex(line, sn)
			if pos >= 0 && pos >= index {
				n = i + 1
				index = pos
			}
		}
	}
	num = num + n
	return num
}

func solve(input string, includeWords bool) (sum int) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		sum = sum + getLineSum(line, includeWords)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return sum
}

func main() {
	input := "input.txt"

	fmt.Println("Task 1 - sum of digits                     \t =  ", solve(input, false))
	fmt.Println("Task 2 - sum of digits and written numbers \t =  ", solve(input, true))
}
