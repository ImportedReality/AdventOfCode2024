package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type coord struct {
	x, y int
	dir  Direction
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

var nextDir = map[Direction]Direction{
	Up:    Right,
	Right: Down,
	Down:  Left,
	Left:  Up,
}

var dirOffsets = map[Direction]coord{
	Up:    {0, -1, Up},
	Right: {1, 0, Right},
	Down:  {0, 1, Down},
	Left:  {-1, 0, Left},
}

func main() {
	lab_map := readInput("input.txt")

	path := part1(dupclicateMap(lab_map))

	part2(lab_map, path)
}

func readInput(file string) [][]string {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lab_map := [][]string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		lab_map = append(lab_map, row)
	}

	return lab_map
}

func dupclicateMap(lab_map [][]string) [][]string {
	duplicate := make([][]string, len(lab_map))

	for i := range lab_map {
		duplicate[i] = append([]string{}, lab_map[i]...)
	}

	return duplicate
}

func printMap(lab_map [][]string) {
	for _, row := range lab_map {
		for _, col := range row {
			print(col)
		}
		print("\n")
	}
}

func part1(lab_map [][]string) (path []coord) {
	path = getPath(lab_map)
	fmt.Printf("Guard visited %d spaces.\n", countGuardHist(lab_map))

	return
}

func part2(lab_map [][]string, path []coord) {
	obstacle_locations := getObstacleLocations(lab_map, path)

	fmt.Printf("There are %d positions where an obstacle will create a loop\n", obstacle_locations)
}

func getGuardLoc(lab_map [][]string) coord {
	for y, row := range lab_map {
		for x, col := range row {
			if col == "^" || col == ">" || col == "<" || col == "V" {
				return coord{x, y, Up}
			}
		}
	}
	return coord{-1, -1, Down}
}

func getPath(lab_map [][]string) (path []coord) {
	guard_loc := getGuardLoc(lab_map)
	path = []coord{guard_loc}
	// dir := Up

	for !isLeavingMap(lab_map, guard_loc) {
		if isBlocked(lab_map, guard_loc) {
			guard_loc.dir = nextDir[guard_loc.dir]
		}
		new_loc := getNewCoords(guard_loc)
		path = append(path, new_loc)
		lab_map[guard_loc.y][guard_loc.x] = "X"
		guard_loc = new_loc
		lab_map[guard_loc.y][guard_loc.x] = "G"
	}
	lab_map[guard_loc.y][guard_loc.x] = "X"

	return
}

func isLeavingMap(lab_map [][]string, guard_loc coord) bool {
	return guard_loc.dir == Up && guard_loc.y == 0 ||
		guard_loc.dir == Right && guard_loc.x == len(lab_map[guard_loc.y])-1 ||
		guard_loc.dir == Down && guard_loc.y == len(lab_map)-1 ||
		guard_loc.dir == Left && guard_loc.x == 0
}

func isBlocked(lab_map [][]string, guard_loc coord) bool {
	new_loc := getNewCoords(guard_loc)
	return lab_map[new_loc.y][new_loc.x] == "#" || lab_map[new_loc.y][new_loc.x] == "O"
}

func getNewCoords(loc coord) coord {
	offset := dirOffsets[loc.dir]

	return coord{loc.x + offset.x, loc.y + offset.y, loc.dir}
}

func countGuardHist(lab_map [][]string) (count int) {
	count = 0
	for _, row := range lab_map {
		for _, col := range row {
			if col == "X" {
				count++
			}
		}
	}
	return
}

// Try putting an obstacle in every location along the guards path,
//
//	then check if it creates a loop.
func getObstacleLocations(lab_map [][]string, path []coord) int {
	starting_loc := getGuardLoc(lab_map)
	valid_obstacles := []coord{}

	for i, loc := range path {
		if i == 0 || loc == starting_loc || obstacleExists(valid_obstacles, loc) {
			continue
		}
		tmp_map := dupclicateMap(lab_map)
		tmp_map[loc.y][loc.x] = "O"
		guard_loc := starting_loc

		if hasLoop(tmp_map, guard_loc) {
			valid_obstacles = append(valid_obstacles, loc)
		}
	}

	return len(valid_obstacles)
}

func hasLoop(lab_map [][]string, guard_loc coord) bool {
	visited := []coord{guard_loc}

	for !isLeavingMap(lab_map, guard_loc) {
		for isBlocked(lab_map, guard_loc) {
			guard_loc.dir = nextDir[guard_loc.dir]
		}
		new_loc := getNewCoords(guard_loc)
		if alreadyVisited(visited, new_loc) {
			return true
		}
		lab_map[guard_loc.y][guard_loc.x] = "X"
		guard_loc = new_loc
		lab_map[guard_loc.y][guard_loc.x] = "G"
		visited = append(visited, guard_loc)
	}

	return false
}

func obstacleExists(obstacles []coord, obstacle coord) bool {
	for _, v := range obstacles {
		if v.x == obstacle.x && v.y == obstacle.y {
			return true
		}
	}

	return false
}

func alreadyVisited(visited []coord, loc coord) bool {
	for _, v := range visited {
		if v == loc {
			return true
		}
	}
	return false
}
