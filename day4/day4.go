package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readInput("input.txt")

	part1(input)
	part2(input)
}

func readInput(file string) [][]string {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		lines = append(lines, split)
	}

	return lines
}

func part1(input [][]string) {
	count := 0
	for y, row := range input {
		for x, v := range row {
			if v == "X" {
				count += checkVert(input, x, y)
				count += checkHoriz(input, x, y)
				count += checkDiag(input, x, y)
			}
		}
	}
	fmt.Printf("XMAS' Found: %d\n", count)
}

func part2(input [][]string) {
	count := 0
	for y, row := range input {
		for x, v := range row {
			if v == "A" {
				if isValidXMas(input, x, y) {
					count++
				}
			}
		}
	}
	fmt.Printf("X-MAS' Found: %d\n", count)
}

func checkVert(input [][]string, x, y int) int {
	count := 0

	//Up
	if y >= 3 {
		lastChar := input[y][x]
		for i := 1; i <= 3; i++ {
			nextChar := input[y-i][x]
			if isNext(nextChar, lastChar) {
				lastChar = nextChar
			} else {
				break
			}
			if lastChar == "S" {
				count++
			}
		}
	}

	//Down
	if y <= len(input)-4 {
		lastChar := input[y][x]
		for i := 1; i <= 3; i++ {
			nextChar := input[y+i][x]
			if isNext(nextChar, lastChar) {
				lastChar = nextChar
			} else {
				break
			}
			if lastChar == "S" {
				count++
			}
		}
	}

	return count
}

func checkHoriz(input [][]string, x, y int) int {
	count := 0

	//Left
	if x >= 3 {
		lastChar := input[y][x]
		for i := 1; i <= 3; i++ {
			nextChar := input[y][x-i]
			if isNext(nextChar, lastChar) {
				lastChar = nextChar
			} else {
				break
			}
			if lastChar == "S" {
				count++
			}
		}
	}

	//Right
	if x <= len(input[y])-4 {
		lastChar := input[y][x]
		for i := 1; i <= 3; i++ {
			nextChar := input[y][x+i]
			if isNext(nextChar, lastChar) {
				lastChar = nextChar
			} else {
				break
			}
			if lastChar == "S" {
				count++
			}
		}
	}

	return count
}

func checkDiag(input [][]string, x, y int) int {
	count := 0

	//Up-Left
	if x >= 3 && y >= 3 {
		lastChar := input[y][x]
		for i := 1; i <= 3; i++ {
			nextChar := input[y-i][x-i]
			if isNext(nextChar, lastChar) {
				lastChar = nextChar
			} else {
				break
			}
			if lastChar == "S" {
				count++
			}
		}
	}

	//Up-Right
	if x <= len(input[y])-4 && y >= 3 {
		lastChar := input[y][x]
		for i := 1; i <= 3; i++ {
			nextChar := input[y-i][x+i]
			if isNext(nextChar, lastChar) {
				lastChar = nextChar
			} else {
				break
			}
			if lastChar == "S" {
				count++
			}
		}
	}

	//Down-Left
	if x >= 3 && y <= len(input)-4 {
		lastChar := input[y][x]
		for i := 1; i <= 3; i++ {
			nextChar := input[y+i][x-i]
			if isNext(nextChar, lastChar) {
				lastChar = nextChar
			} else {
				break
			}
			if lastChar == "S" {
				count++
			}
		}
	}

	//Down-Right
	if x <= len(input[y])-4 && y <= len(input)-4 {
		lastChar := input[y][x]
		for i := 1; i <= 3; i++ {
			nextChar := input[y+i][x+i]
			if isNext(nextChar, lastChar) {
				lastChar = nextChar
			} else {
				break
			}
			if lastChar == "S" {
				count++
			}
		}
	}

	return count
}

func isNext(nextChar string, lastChar string) bool {
	return lastChar == "X" && nextChar == "M" || lastChar == "M" && nextChar == "A" || lastChar == "A" && nextChar == "S"
}

func isValidXMas(input [][]string, x, y int) bool {
	valid := false

	if x > 0 && x < len(input[y])-1 && y > 0 && y < len(input)-1 {
		topLeft := input[y-1][x-1]
		topRight := input[y-1][x+1]
		btmLeft := input[y+1][x-1]
		btmRight := input[y+1][x+1]

		valid = (topLeft == "M" && btmRight == "S" || topLeft == "S" && btmRight == "M") &&
			(topRight == "M" && btmLeft == "S" || topRight == "S" && btmLeft == "M")
	}

	return valid
}
