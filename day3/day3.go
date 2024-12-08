package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	instructions := readBasicInstructions(f)

	sum := 0
	for _, v := range instructions {
		sum += (v[0] * v[1])
	}

	fmt.Printf("Sum with basic instructions: %d\n", sum)
}

func part2() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	instructions := readAdvancedInstructions(f)

	sum := 0
	for _, v := range instructions {
		sum += (v[0] * v[1])
	}

	fmt.Printf("Sum with advanced instructions: %d\n", sum)
}

func readBasicInstructions(f *os.File) [][]int {
	instructions := [][]int{}
	scanner := bufio.NewScanner(f)
	r := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")

	for scanner.Scan() {
		line := scanner.Text()
		matches := r.FindAllStringSubmatch(line, -1)

		for _, v := range matches {
			a, err := strconv.Atoi(v[1])
			if err != nil {
				panic(err)
			}
			b, err := strconv.Atoi(v[2])
			if err != nil {
				panic(err)
			}
			vals := []int{
				a, b,
			}
			instructions = append(instructions, vals)
		}
	}
	return instructions
}

func readAdvancedInstructions(f *os.File) [][]int {
	instructions := [][]int{}
	scanner := bufio.NewScanner(f)
	r := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)|don\\'t\\(\\)|do\\(\\)")
	enabled := true

	for scanner.Scan() {
		line := scanner.Text()
		matches := r.FindAllStringSubmatch(line, -1)

		for _, v := range matches {
			switch v[0] {
			case "don't()":
				enabled = false
				break
			case "do()":
				enabled = true
				break
			default:
				if !enabled {
					continue
				}
				a, err := strconv.Atoi(v[1])
				if err != nil {
					panic(err)
				}
				b, err := strconv.Atoi(v[2])
				if err != nil {
					panic(err)
				}
				vals := []int{
					a, b,
				}
				instructions = append(instructions, vals)
			}
		}
	}
	return instructions
}
