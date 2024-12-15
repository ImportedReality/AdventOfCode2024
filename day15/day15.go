package main

import (
	"fmt"
	"strings"

	"github.com/ImportedReality/aocutils"
)

const filename = "input.test3.txt"

type Robot struct {
	loc   aocutils.Coordinate
	moves []aocutils.Direction
}

var Map aocutils.Grid[string]
var WideMap aocutils.Grid[string]
var robot Robot
var moveDirs = map[string]aocutils.Direction{
	"^": aocutils.N,
	">": aocutils.E,
	"v": aocutils.S,
	"<": aocutils.W,
}
var expandedTiles = map[string][]string{
	"#": {"#", "#"},
	"O": {"[", "]"},
	".": {".", "."},
	"@": {"@", "."},
}

func main() {
	input := aocutils.ReadLines(filename)
	readMap(input)
	readMoves(input)
	part1()
	readWideMap(input)
	part2()
}

func readMap(input []string) {
	Map = make(aocutils.Grid[string], 0)
	for y, line := range input {
		if len(line) == 0 {
			break
		}
		botIdx := strings.Index(line, "@")
		if botIdx >= 0 {
			robot = Robot{
				aocutils.Coordinate{X: botIdx, Y: y},
				[]aocutils.Direction{},
			}
		}
		row := strings.Split(line, "")
		Map = append(Map, row)
	}
}

func readWideMap(input []string) {
	WideMap = make(aocutils.Grid[string], 0)
	for y, line := range input {
		if len(line) == 0 {
			break
		}
		tiles := strings.Split(line, "")
		row := make([]string, 0)
		for x, tile := range tiles {
			if tile == "@" {
				robot = Robot{
					aocutils.Coordinate{X: x * 2, Y: y},
					[]aocutils.Direction{},
				}
			}
			row = append(row, expandedTiles[tile]...)
		}
		WideMap = append(WideMap, row)
	}
}

func readMoves(input []string) {
	moves := make([]aocutils.Direction, 0)
	for _, line := range input {
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		row := strings.Split(line, "")
		for _, inst := range row {
			dir := moveDirs[inst]
			moves = append(moves, dir)
		}
	}

	robot.moves = moves
}

func part1() {
	for _, move := range robot.moves {
		moveRobot(move)
	}
	GPSSum := calculateGPSSum()
	fmt.Printf("The sum of all boxes' GPS coordinates is %d\n", GPSSum)
}

func part2() {
	println("Initial State:")
	aocutils.PrintGrid(WideMap, "")
	for _, move := range robot.moves {
		moveRobotPt2(move)
	}
}

func moveRobot(move aocutils.Direction) {
	loc := robot.loc
	offset := aocutils.Offsets[move]
	newLoc := aocutils.Coordinate{X: loc.X + offset.X, Y: loc.Y + offset.Y}
	target := Map[newLoc.Y][newLoc.X]
	switch target {
	case "#":
		break
	case ".":
		robot.loc = newLoc
		Map[newLoc.Y][newLoc.X] = "@"
		Map[loc.Y][loc.X] = "."
		break
	case "O":
		// Only move robot if moving boxes was successful
		if moveBoxes(loc, offset) {
			robot.loc = newLoc
			Map[newLoc.Y][newLoc.X] = "@"
			Map[loc.Y][loc.X] = "."
		}
		break
	}
}

// Find furthest box
// Check if that box is blocked
// If not, move all boxes one space over
// update robot location
func moveBoxes(robotLoc, offset aocutils.Coordinate) bool {
	finalLoc := aocutils.Coordinate{X: robotLoc.X + offset.X, Y: robotLoc.Y + offset.Y}
	for Map[finalLoc.Y][finalLoc.X] != "." {
		if isBlocked(finalLoc, offset) {
			return false
		}
		finalLoc = aocutils.Coordinate{X: finalLoc.X + offset.X, Y: finalLoc.Y + offset.Y}
	}
	prevLoc := finalLoc
	currentLoc := aocutils.Coordinate{X: prevLoc.X - offset.X, Y: prevLoc.Y - offset.Y}
	for currentLoc != robotLoc {
		Map[prevLoc.Y][prevLoc.X] = "O"
		Map[currentLoc.Y][currentLoc.X] = "."
		prevLoc = currentLoc
		currentLoc = aocutils.Coordinate{X: prevLoc.X - offset.X, Y: prevLoc.Y - offset.Y}
	}
	return true
}

func isBlocked(loc, offset aocutils.Coordinate) bool {
	target := Map[loc.Y+offset.Y][loc.X+offset.X]
	return target == "#"
}

func calculateGPSSum() int {
	sum := 0
	for y, row := range Map {
		for x, col := range row {
			if col == "O" {
				sum += 100*y + x
			}
		}
	}
	return sum
}

func moveRobotPt2(move aocutils.Direction) {
	loc := robot.loc
	offset := aocutils.Offsets[move]
	newLoc := aocutils.Coordinate{X: loc.X + offset.X, Y: loc.Y + offset.Y}
	target := Map[newLoc.Y][newLoc.X]
	switch target {
	case "#":
		break
	case ".":
		robot.loc = newLoc
		Map[newLoc.Y][newLoc.X] = "@"
		Map[loc.Y][loc.X] = "."
		break
	case "[":
	case "]":
		if moveWideBoxes(loc, offset) {
			robot.loc = newLoc
			Map[newLoc.Y][newLoc.X] = "@"
			Map[loc.Y][loc.X] = "."
		}
		break
	}
}

func moveWideBoxes(robotLoc, offset aocutils.Coordinate) bool {
	finalLoc := aocutils.Coordinate{X: robotLoc.X + offset.X, Y: robotLoc.Y + offset.Y}
	for Map[finalLoc.Y][finalLoc.X] != "." {
		if isBlocked(finalLoc, offset) {
			return false
		}
		finalLoc = aocutils.Coordinate{X: finalLoc.X + offset.X, Y: finalLoc.Y + offset.Y}
	}

	return false
}
