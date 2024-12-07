package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Equations map[int][]int

var Operators = []string{"+", "*"}

func main() {
	equations := readInput("./input.test.txt")

	part1(equations)
}

func readInput(file string) Equations {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := make(Equations)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		equation := strings.Split(line, ":")
		nums := strings.Split(strings.TrimSpace(equation[1]), " ")

		result := strToInt(equation[0])
		operands := []int{}
		for _, v := range nums {
			operands = append(operands, strToInt(v))
		}

		input[result] = operands
	}

	return input
}

func strToInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return val
}

func part1(equations Equations) {
	total := 0

	for result, values := range equations {
		if canBeSolved(result, values) {
			total += result
		}
	}

	fmt.Printf("Total calibration result: %d\n", total)
}

func canBeSolved(result int, values []int) bool {
	subsets := pow(2, len(values)-1)
	fmt.Printf("%d possible combinations for %d = %d\n", subsets, result, values)

	return false
}

func pow(n, m int) int {
	if m == 0 {
		return 1
	}
	if m == 1 {
		return n
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func generatePermutations(n int, results [][]string) (results [][]string) {
    if n < 2 {
        return 
    }
    for i := 0

    return
}
