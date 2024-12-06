package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := read_input("input.txt")

	rules := make_rules_map(input)
	page_lists := make_page_lists(input)

	part1(rules, page_lists)
	part2(rules, page_lists)
}

func read_input(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	return input
}

func make_rules_map(input []string) map[int][]int {
	rules := make(map[int][]int)

	for _, v := range input {
		if strings.Contains(v, "|") {
			vals := strings.Split(v, "|")
			before, err := strconv.Atoi(vals[0])
			if err != nil {
				panic(err)
			}
			after, err := strconv.Atoi(vals[1])
			if err != nil {
				panic(err)
			}

			rule := rules[after]

			if !contains(rule, before) {
				rule = append(rule, before)
				rules[after] = rule
			}
		}
	}

	return rules
}

func contains(rule []int, x int) bool {
	for _, v := range rule {
		if v == x {
			return true
		}
	}

	return false
}

func make_page_lists(input []string) [][]int {
	page_lists := [][]int{}

	for _, v := range input {
		if strings.Contains(v, ",") {
			list := []int{}
			vals := strings.Split(v, ",")

			for _, val := range vals {
				num, err := strconv.Atoi(val)
				if err != nil {
					panic(err)
				}
				list = append(list, num)
			}

			page_lists = append(page_lists, list)
		}
	}

	return page_lists
}

func part1(rules map[int][]int, page_lists [][]int) {
	sum := 0

	for _, list := range page_lists {
		if valid_list(rules, list) {
			sum += list[len(list)/2]
		}
	}

	fmt.Printf("Sum of correctly-ordered middle page numbers: %d\n", sum)
}

func part2(rules map[int][]int, page_lists [][]int) {
	sum := 0

	for _, list := range page_lists {
		if !valid_list(rules, list) {
			fixed := fix_list(rules, list)
			sum += list[len(fixed)/2]
		}
	}

	fmt.Printf("Sum of fixed middle page numbers: %d\n", sum)
}

func valid_list(rules map[int][]int, list []int) bool {
	is_valid := true

	for i, page := range list {
		if len(rules[page]) > 0 {
			page_rules := rules[page]
			for j := i + 1; j < len(list); j++ {
				if !valid_num(page_rules, list[j]) {
					is_valid = false
				}
			}
		}
	}

	return is_valid
}

func valid_num(rules []int, num int) bool {
	for _, v := range rules {
		if v == num {
			return false
		}
	}

	return true
}

func fix_list(rules map[int][]int, list []int) []int {
	fixed := list

	for !valid_list(rules, fixed) {
		for i := range fixed {
			idx := invalid_idx(rules, fixed, i)

			if idx == -1 {
			} else {
				fixed[i], fixed[idx] = fixed[idx], fixed[i]
			}
		}
	}

	return fixed
}

// Returns an int representing the index of the first broken rule,
// or -1 if no rules are broken
func invalid_idx(rules map[int][]int, list []int, i int) int {
	val := list[i]

	if i == len(list)-1 {
		return -1
	}

	for j := i + 1; i < len(list); i++ {
		for _, v := range rules[val] {
			if v == list[j] {
				return j
			}
		}
	}

	return -1
}
