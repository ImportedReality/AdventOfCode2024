package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calibration struct {
	result   int
	operands []int
}

type Operator int

const (
	Add Operator = iota
	Mult
	Cat
)

const Ops = 3

func main() {
	calibrations := readInput("./input.txt")

	solve(calibrations)
}

func readInput(file string) []Calibration {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := []Calibration{}

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

		calibration := Calibration{result, operands}

		input = append(input, calibration)
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

func solve(calibrations []Calibration) {
	total := 0

	for _, c := range calibrations {
		if canBeSolved(c.result, c.operands) {
			total += c.result
		}
	}

	fmt.Printf("Total calibration result: %d\n", total)
}

func canBeSolved(result int, values []int) bool {
	subsets := pow(Ops, len(values)-1)
	permutations := generatePermutations(subsets)

	for _, p := range permutations {
		if testEquation(result, values, p) {
			return true
		}
	}

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

func generatePermutations(n int) (results [][]Operator) {
	results = [][]Operator{}

	for i := range n {
		str := strconv.FormatInt(int64(i), Ops)
		values := strings.Split(str, "")
		set := []Operator{}
		for _, v := range values {
			set = append(set, Operator(strToInt(v)))
		}
		results = append(results, set)
	}

	results = normalizeOperationsLength(results)

	return
}

func normalizeOperationsLength(operations [][]Operator) (results [][]Operator) {
	results = [][]Operator{}
	max_len := 1
	for _, v := range operations {
		if len(v) > max_len {
			max_len = len(v)
		}
	}
	for _, v := range operations {
		for len(v) < max_len {
			v = append([]Operator{Operator(0)}, v...)
		}
		results = append(results, v)
	}

	return
}

func testEquation(result int, values []int, operations []Operator) bool {
	a := values[0]
	b := -1
	for i := 1; i < len(values); i++ {
		b = values[i]
		op := operations[i-1]
		a = performOperation(a, b, op)
	}

	return a == result
}

func performOperation(a, b int, op Operator) int {
	switch op {
	case Add:
		return a + b
	case Mult:
		return a * b
	case Cat:
		return a*padding(b) + b
	default:
		panic("HOW")
	}
}

func padding(n int) int {
	p := 1
	for p <= n {
		p *= 10
	}
	return p
}
