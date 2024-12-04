package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	left := []int{}
	right := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")

		lnum, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		rnum, err := strconv.Atoi(nums[len(nums)-1])
		if err != nil {
			panic(err)
		}

		left = append(left, lnum)
		right = append(right, rnum)
	}

	part1(left, right)
	part2(left, right)
}

func part1(left []int, right []int) {
	sort.Ints(left[:])
	sort.Ints(right[:])
	dists := []int{}

	for i, v := range left {
		dist := abs(v - right[i])
		dists = append(dists, dist)
	}

	total_dist := reduce(dists)
	fmt.Printf("Total Distance: %d\n", total_dist)
}

func part2(left []int, right []int) {
	rfreq := make(map[int]int)
	sim := []int{}

	for _, v := range right {
		val := rfreq[v]
		val++
		rfreq[v] = val
	}

	for _, v := range left {
		val := rfreq[v] * v
		sim = append(sim, val)
	}

	sim_score := reduce(sim)
	fmt.Printf("Similarity Score: %d\n", sim_score)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func reduce(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
