package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	lines := readLines(f)

	part1(lines)
	part2(lines)
}

func readLines(f *os.File) [][]int {
	scanner := bufio.NewScanner(f)
	lines := [][]int{}

	for scanner.Scan() {
		text := scanner.Text()
		vals := strings.Split(text, " ")
		line := []int{}

		for _, v := range vals {
			val, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			line = append(line, val)
		}

		lines = append(lines, line)
	}

	return lines
}

func part1(lines [][]int) {
	safeCount := 0
	for _, line := range lines {
		if lineSafe(line) {
			safeCount++
		}
	}

	fmt.Printf("Safe reports: %d\n", safeCount)
}

func part2(lines [][]int) {
	safeCount := 0
	for _, line := range lines {
		if lineSafeDampened(line) {
			safeCount++
		}
	}

	fmt.Printf("Safe reports with dampening: %d\n", safeCount)
}

func lineSafeDampened(line []int) bool {
	if lineSafe(line) {
		return true
	}

	for i := range line {
		subLine := remove(line, i)
		if lineSafe(subLine) {
			return true
		}
	}

	return false
}

func lineSafe(line []int) bool {
	// bool representing whether or not to expect increases or decreases
	inc := false
	for i, v := range line {
		if i == len(line)-1 {
			break
		}
		if i == 0 {
			inc = v < line[i+1]
		}

		if inc && v > line[i+1] || !inc && v < line[i+1] {
			return false
		}

		diff := abs(v - line[i+1])

		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func remove(slice []int, s int) []int {
	result := make([]int, 0, len(slice)-1)
	for i, v := range slice {
		if i == s {
			continue
		}

		result = append(result, v)
	}

	return result
}
