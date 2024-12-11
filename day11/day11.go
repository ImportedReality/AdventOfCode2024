package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var stones = map[int]int{}
var blinks = 75

func main() {
	readInput("input.txt")
	start := time.Now()
	solve()
	duration := time.Since(start)
	fmt.Printf("Exectuion time: %v\n", duration)
}

func readInput(file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	line, err := bufio.NewReader(f).ReadString('\n')
	if err != nil {
		panic(err)
	}

	vals := strings.Split(line, " ")
	for _, val := range vals {
		num := strToInt(val)
		stones[num] += 1
	}
}

func strToInt(s string) int {
	num, err := strconv.Atoi(strings.Trim(s, "\n"))
	if err != nil {
		fmt.Println(s)
		panic(err)
	}
	return num
}

func solve() {
	for range blinks {
		blink()
	}
	fmt.Printf("After %d blinks there are %d stones\n", blinks, countStones())
}

func blink() {
	updatedStones := make(map[int]int)
	for val, count := range stones {
		if val == 0 {
			updatedStones[1] += count
		} else if digits(val)%2 == 0 {
			l, r := split(val)
			updatedStones[l] += count
			updatedStones[r] += count
		} else {
			updatedStones[val*2024] += count
		}
	}
	stones = updatedStones
}

func digits(n int) int {
	str := strconv.Itoa(n)
	return len(str)
}

func split(n int) (int, int) {
	str := strconv.Itoa(n)
	half := len(str) / 2
	l, r := str[:half], str[half:]
	return strToInt(l), strToInt(r)
}

func countStones() int {
	count := 0
	for _, v := range stones {
		count += v
	}
	return count
}
