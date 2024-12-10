package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type TopoMap [][]int
type Node struct {
	x, y, elevation int
	branches        []*Node
}

var topoMap TopoMap

func (n *Node) addBranch(node *Node) {
	n.branches = append(n.branches, node)
}

func (n *Node) hasBranches() bool {
	return len(n.branches) > 0
}

func (n *Node) getBranches() []*Node {
	return n.branches
}

func main() {
	topoMap = readInput("input.txt")
	start := time.Now()
	solve()
	duration := time.Since(start)
	fmt.Printf("Execution time: %v\n", duration)
}

func readInput(file string) TopoMap {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	topoMap := make(TopoMap, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Split(line, "")
		row := make([]int, 0)
		for _, val := range vals {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			row = append(row, num)
		}
		topoMap = append(topoMap, row)
	}

	return topoMap
}

func solve() {
	trailheads := findTrailheads()
	var score, rating int
	for _, trailhead := range trailheads {
		peaks, trails := findPeaks(&trailhead, &[]*Node{}, &[]*Node{})
		score += len(*peaks)
		rating += len(*trails)
	}

	fmt.Printf("There are %d trailheads with a total score of %d and total rating of %d\n", len(trailheads), score, rating)
}

func findTrailheads() []Node {
	trailheads := make([]Node, 0)
	for y, row := range topoMap {
		for x, col := range row {
			if col == 0 {
				trailheads = append(trailheads,
					Node{
						x, y, col, []*Node{},
					})
			}
		}
	}

	return trailheads
}

func findPeaks(node *Node, peaks *[]*Node, trails *[]*Node) (*[]*Node, *[]*Node) {
	if node.elevation == 9 {
		if !peakVisited(node, peaks) {
			*peaks = append(*peaks, node)
		}
		*trails = append(*trails, node)
		return peaks, trails
	}
	findValidMoves(node)
	if node.hasBranches() {
		for _, branch := range node.getBranches() {
			findPeaks(branch, peaks, trails)
		}
	}
	return peaks, trails
}

func peakVisited(node *Node, peaks *[]*Node) bool {
	for _, peak := range *peaks {
		if node.x == peak.x && node.y == peak.y {
			return true
		}
	}

	return false
}

func findValidMoves(node *Node) {
	x, y := -1, -1
	if node.x > 0 {
		x, y = node.x-1, node.y
		elevation := topoMap[y][x]
		if elevation == node.elevation+1 {
			node.addBranch(&Node{x, y, elevation, make([]*Node, 0)})
		}
	}
	if node.x < len(topoMap[0])-1 {
		x, y = node.x+1, node.y
		elevation := topoMap[y][x]
		if elevation == node.elevation+1 {
			node.addBranch(&Node{x, y, elevation, make([]*Node, 0)})
		}
	}
	if node.y > 0 {
		x, y = node.x, node.y-1
		elevation := topoMap[y][x]
		if elevation == node.elevation+1 {
			node.addBranch(&Node{x, y, elevation, make([]*Node, 0)})
		}
	}
	if node.y < len(topoMap)-1 {
		x, y = node.x, node.y+1
		elevation := topoMap[y][x]
		if elevation == node.elevation+1 {
			node.addBranch(&Node{x, y, elevation, make([]*Node, 0)})
		}
	}
}
