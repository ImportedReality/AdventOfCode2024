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

func main() {
	lab_map := read_input("input.txt")

	path := part1(dupclicate_map(lab_map))

	part2(lab_map, path)
}

func read_input(file string) [][]string {
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

func dupclicate_map(lab_map [][]string) [][]string {
	duplicate := make([][]string, len(lab_map))

	for i := range lab_map {
		duplicate[i] = append([]string{}, lab_map[i]...)
	}

	return duplicate
}

func print_map(lab_map [][]string) {
	for _, row := range lab_map {
		for _, col := range row {
			print(col)
		}
		print("\n")
	}
}

func part1(lab_map [][]string) (path []coord) {
	path = get_path(lab_map)
	fmt.Printf("Guard visited %d spaces.\n", count_guard_hist(lab_map))

	return
}

func part2(lab_map [][]string, path []coord) {
	obstacle_locations := get_obstacle_locations(lab_map, path)

	fmt.Printf("There are %d positions where an obstacle will create a loop\n", obstacle_locations)
}

func get_guard_loc(lab_map [][]string) coord {
	for y, row := range lab_map {
		for x, col := range row {
			if col == "^" || col == ">" || col == "<" || col == "V" {
				return coord{x, y, Up}
			}
		}
	}
	return coord{-1, -1, Down}
}

func get_path(lab_map [][]string) (path []coord) {
	guard_loc := get_guard_loc(lab_map)
	path = []coord{guard_loc}
	// dir := Up

	for !is_leaving_map(lab_map, guard_loc) {
		if is_blocked(lab_map, guard_loc) {
			guard_loc.dir = get_new_dir(guard_loc.dir)
			// dir = get_new_dir(dir)
		}
		new_loc := get_new_coords(guard_loc)
		path = append(path, new_loc)
		lab_map[guard_loc.y][guard_loc.x] = "X"
		guard_loc = new_loc
		lab_map[guard_loc.y][guard_loc.x] = "G"
	}
	lab_map[guard_loc.y][guard_loc.x] = "X"

	return
}

func is_leaving_map(lab_map [][]string, guard_loc coord) bool {
	return guard_loc.dir == Up && guard_loc.y == 0 ||
		guard_loc.dir == Right && guard_loc.x == len(lab_map[guard_loc.y])-1 ||
		guard_loc.dir == Down && guard_loc.y == len(lab_map)-1 ||
		guard_loc.dir == Left && guard_loc.x == 0
}

func is_blocked(lab_map [][]string, guard_loc coord) bool {
	new_loc := get_new_coords(guard_loc)
	return lab_map[new_loc.y][new_loc.x] == "#" || lab_map[new_loc.y][new_loc.x] == "O"
}

func get_new_coords(loc coord) coord {
	switch loc.dir {
	case Up:
		return coord{loc.x, loc.y - 1, loc.dir}
	case Right:
		return coord{loc.x + 1, loc.y, loc.dir}
	case Down:
		return coord{loc.x, loc.y + 1, loc.dir}
	case Left:
		return coord{loc.x - 1, loc.y, loc.dir}
	default:
		panic("Undefined direction")
	}
}

func get_new_dir(dir Direction) Direction {
	switch dir {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	default:
		panic("Undefined direction")
	}
}

func count_guard_hist(lab_map [][]string) (count int) {
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
func get_obstacle_locations(lab_map [][]string, path []coord) int {
	starting_loc := get_guard_loc(lab_map)
	valid_obstacles := []coord{}

	for i, loc := range path {
		if i == 0 || loc == starting_loc || obstacle_exists(valid_obstacles, loc) {
			continue
		}
		tmp_map := dupclicate_map(lab_map)
		tmp_map[loc.y][loc.x] = "O"
		guard_loc := path[i-1]

		if has_loop(tmp_map, guard_loc) {
			valid_obstacles = append(valid_obstacles, loc)
		}
	}

	return len(valid_obstacles)
}

func has_loop(lab_map [][]string, guard_loc coord) bool {
	visited := []coord{guard_loc}

	for !is_leaving_map(lab_map, guard_loc) {
		for is_blocked(lab_map, guard_loc) {
			guard_loc.dir = get_new_dir(guard_loc.dir)
		}
		new_loc := get_new_coords(guard_loc)
		if already_visited(visited, new_loc) {

			return true
		}
		lab_map[guard_loc.y][guard_loc.x] = "X"
		guard_loc = new_loc
		lab_map[guard_loc.y][guard_loc.x] = "G"
		visited = append(visited, guard_loc)
	}

	return false
}

func obstacle_exists(obstacles []coord, obstacle coord) bool {
	for _, v := range obstacles {
		if v == obstacle {
			return true
		}
	}

	return false
}

func already_visited(visited []coord, loc coord) bool {
	for _, v := range visited {
		if v == loc {
			return true
		}
	}
	return false
}
