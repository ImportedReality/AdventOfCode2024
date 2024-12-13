package main

import (
	"fmt"
	"regexp"

	"github.com/ImportedReality/aocutils"
)

var filename = "input.txt"

type Equation struct {
	ax, bx, ay, by, x, y int
}

var buttonRegexp = regexp.MustCompile("\\+(\\d+).*\\+(\\d+)")
var prizeRegexp = regexp.MustCompile("\\=(\\d+).*\\=(\\d+)")

func main() {
	input := aocutils.ReadLines(filename)
	equations := getEquations(input)
	part1(equations)
	part2(equations)
}

func getEquations(input []string) []Equation {
	equations := make([]Equation, 0)
	for i := 0; i < len(input); i += 4 {
		buttonA := input[i]
		buttonB := input[i+1]
		prize := input[i+2]

		ax, ay := parseValues(buttonA, buttonRegexp)
		bx, by := parseValues(buttonB, buttonRegexp)
		x, y := parseValues(prize, prizeRegexp)

		eq := Equation{ax, bx, ay, by, x, y}
		equations = append(equations, eq)
	}

	return equations
}

func parseValues(str string, reg *regexp.Regexp) (int, int) {
	x, y := -1, -1
	match := reg.FindAllStringSubmatch(str, -1)
	if len(match) > 1 || len(match[0]) > 3 {
		panic("Unexpected button values detected!")
	}
	x = aocutils.StrToInt(match[0][1])
	y = aocutils.StrToInt(match[0][2])

	return x, y
}

func part1(equations []Equation) {
	total := 0
	for _, eq := range equations {
		a, b := solveEq(eq)
		total += a * 3
		total += b
	}

	fmt.Printf("%d tokens\n", total)
}

func part2(equations []Equation) {
	total := 0
	for _, eq := range equations {
		eq.x = eq.x + 10000000000000
		eq.y = eq.y + 10000000000000
		a, b := solveEq(eq)
		total += a * 3
		total += b
	}

	fmt.Printf("%d tokens\n", total)
}

// Solve equation using Cramer's rule, return 0,0 if the result is not a whole
// number.
func solveEq(eq Equation) (int, int) {
	d := eq.ax*eq.by - eq.bx*eq.ay
	d1 := eq.x*eq.by - eq.y*eq.bx
	d2 := eq.y*eq.ax - eq.x*eq.ay

	if d1%d != 0 || d2%d != 0 {
		return 0, 0
	}

	return d1 / d, d2 / d
}
