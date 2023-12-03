package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func min(nums ...int) int {
	m := nums[0]
	for _, n := range nums {
		if n < m {
			m = n
		}
	}
	return m
}

func max(nums ...int) int {
	m := nums[0]
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}

type field []string

func readData(input string) field {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var fi field
	for scanner.Scan() {
		line := scanner.Text()
		fi = append(fi, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return fi
}

type number struct {
	value          int
	xstart, ystart int
	xend, yend     int
	isPartNumber   bool
}

func (n *number) setValue(value string) {
	fmt.Sscanf(value, "%d", &n.value)
}

func NewNumber(value string, xstart, ystart, xend, yend int) *number {
	n := &number{xstart: xstart, ystart: ystart, xend: xend, yend: yend}
	n.setValue(value)
	return n
}

func (n number) String() string {
	return fmt.Sprintf("%d", n.value)
}

func isPart(value byte) bool {
	return value != '.' && !isDigit(value)
}

func isGear(value byte) bool {
	return value == '*'
}

func isDigit(value byte) bool {
	return value >= '0' && value <= '9'
}

func (f field) getNumbers() []*number {
	var numbers []*number
	for y, line := range f {
		for x := 0; x < len(line); x++ {

			if isDigit(line[x]) {
				value := string(line[x])
				xstart := x
				xend := x
				ystart := y
				yend := y
				for xx := x + 1; xx < len(line); xx++ {
					if isDigit(line[xx]) {
						value += string(line[xx])
						xend = xx
						yend = y
					} else {
						break
					}
				}
				x = xend
				numbers = append(numbers, NewNumber(value, xstart, ystart, xend, yend))
			}
		}
	}
	return numbers
}

type coord struct {
	x, y int
}

func (c coord) String() string {
	return fmt.Sprintf("[%d,%d]", c.x, c.y)
}

func (f field) getANumbersNeighbours(num *number) []*coord {
	var neighbours []*coord
	for y := max(0, num.ystart-1); y <= min(len(f)-1, num.yend+1); y++ {
		for x := max(0, num.xstart-1); x <= min(len(f[y])-1, num.xend+1); x++ {
			if x >= num.xstart && x <= num.xend && y >= num.ystart && y <= num.yend {
				x = num.xend
			} else {
				neighbours = append(neighbours, &coord{x: x, y: y})
			}
		}
	}
	return neighbours
}

func solve2(input string) (sumOfGearPowers int) {
	field := readData(input)
	nums := field.getNumbers()
	geared := make(map[string][]*number)
	for _, num := range nums {
		adjacents := field.getANumbersNeighbours(num)
		for _, adj := range adjacents {
			if isGear(field[adj.y][adj.x]) {
				geared[adj.String()] = append(geared[adj.String()], num)
				break
			}
		}
	}
	for _, nums := range geared {
		if len(nums) > 1 {
			power := 1
			for _, n := range nums {
				power *= n.value
			}
			sumOfGearPowers += power
		}

	}

	return sumOfGearPowers
}

func solve1(input string) (sumOfPartNumbers int) {
	field := readData(input)
	nums := field.getNumbers()
	for _, num := range nums {
		adjacents := field.getANumbersNeighbours(num)
		for _, adj := range adjacents {
			if isPart(field[adj.y][adj.x]) {
				sumOfPartNumbers += num.value
				num.isPartNumber = true
				break
			}
		}
	}
	return sumOfPartNumbers
}

func main() {
	input := "input.txt"

	t1 := solve1(input)
	t2 := solve2(input)
	fmt.Println("Task 1 - sum of part numbers \t =  ", t1)
	fmt.Println("Task 2 - sum of gear powers  \t =  ", t2)
}
