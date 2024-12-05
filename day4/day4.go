package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := read_input("input.txt")

	part1(input)
	part2(input)
}

func read_input(file string) [][]string {
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
				count += check_vert(input, x, y)
				count += check_horiz(input, x, y)
				count += check_diag(input, x, y)
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
				if is_valid_x_mas(input, x, y) {
					count++
				}
			}
		}
	}
	fmt.Printf("X-MAS' Found: %d\n", count)
}

func check_vert(input [][]string, x, y int) int {
	count := 0

	//Up
	if y >= 3 {
		last_char := input[y][x]
		for i := 1; i <= 3; i++ {
			next_char := input[y-i][x]
			if is_next(next_char, last_char) {
				last_char = next_char
			} else {
				break
			}
			if last_char == "S" {
				count++
			}
		}
	}

	//Down
	if y <= len(input)-4 {
		last_char := input[y][x]
		for i := 1; i <= 3; i++ {
			next_char := input[y+i][x]
			if is_next(next_char, last_char) {
				last_char = next_char
			} else {
				break
			}
			if last_char == "S" {
				count++
			}
		}
	}

	return count
}

func check_horiz(input [][]string, x, y int) int {
	count := 0

	//Left
	if x >= 3 {
		last_char := input[y][x]
		for i := 1; i <= 3; i++ {
			next_char := input[y][x-i]
			if is_next(next_char, last_char) {
				last_char = next_char
			} else {
				break
			}
			if last_char == "S" {
				count++
			}
		}
	}

	//Right
	if x <= len(input[y])-4 {
		last_char := input[y][x]
		for i := 1; i <= 3; i++ {
			next_char := input[y][x+i]
			if is_next(next_char, last_char) {
				last_char = next_char
			} else {
				break
			}
			if last_char == "S" {
				count++
			}
		}
	}

	return count
}

func check_diag(input [][]string, x, y int) int {
	count := 0

	//Up-Left
	if x >= 3 && y >= 3 {
		last_char := input[y][x]
		for i := 1; i <= 3; i++ {
			next_char := input[y-i][x-i]
			if is_next(next_char, last_char) {
				last_char = next_char
			} else {
				break
			}
			if last_char == "S" {
				count++
			}
		}
	}

	//Up-Right
	if x <= len(input[y])-4 && y >= 3 {
		last_char := input[y][x]
		for i := 1; i <= 3; i++ {
			next_char := input[y-i][x+i]
			if is_next(next_char, last_char) {
				last_char = next_char
			} else {
				break
			}
			if last_char == "S" {
				count++
			}
		}
	}

	//Down-Left
	if x >= 3 && y <= len(input)-4 {
		last_char := input[y][x]
		for i := 1; i <= 3; i++ {
			next_char := input[y+i][x-i]
			if is_next(next_char, last_char) {
				last_char = next_char
			} else {
				break
			}
			if last_char == "S" {
				count++
			}
		}
	}

	//Down-Right
	if x <= len(input[y])-4 && y <= len(input)-4 {
		last_char := input[y][x]
		for i := 1; i <= 3; i++ {
			next_char := input[y+i][x+i]
			if is_next(next_char, last_char) {
				last_char = next_char
			} else {
				break
			}
			if last_char == "S" {
				count++
			}
		}
	}

	return count
}

func is_next(next_char string, last_char string) bool {
	return last_char == "X" && next_char == "M" || last_char == "M" && next_char == "A" || last_char == "A" && next_char == "S"
}

func is_valid_x_mas(input [][]string, x, y int) bool {
	valid := false

	if x > 0 && x < len(input[y])-1 && y > 0 && y < len(input)-1 {
		top_left := input[y-1][x-1]
		top_right := input[y-1][x+1]
		btm_left := input[y+1][x-1]
		btm_right := input[y+1][x+1]

		valid = (top_left == "M" && btm_right == "S" || top_left == "S" && btm_right == "M") &&
			(top_right == "M" && btm_left == "S" || top_right == "S" && btm_left == "M")
	}

	return valid
}
