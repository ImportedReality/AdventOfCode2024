package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ImportedReality/aocutils"
)

var filename = "input.txt"

const gridX, gridY = 101, 103
const maxX, maxY = 100, 102

const part1Seconds = 100
const part2Seconds = 10000

type Coord aocutils.Coordinate

var robotRegexp = regexp.MustCompile("p=(-?\\d+,-?\\d+)\\sv=(-?\\d+,-?\\d+)")

type Robot struct {
	position, velocity Coord
}

func main() {
	input := aocutils.ReadLines(filename)
	robots := parseRobots(input)
	part1(robots)
	robots = parseRobots(input)
	part2(robots)
}

func parseRobots(input []string) []Robot {
	robots := make([]Robot, 0)

	for _, line := range input {
		matches := robotRegexp.FindAllStringSubmatch(line, -1)
		pos := strings.Split(matches[0][1], ",")
		posX, posY := aocutils.StrToInt(pos[0]), aocutils.StrToInt(pos[1])
		vel := strings.Split(matches[0][2], ",")
		velX, velY := aocutils.StrToInt(vel[0]), aocutils.StrToInt(vel[1])

		robot := Robot{
			Coord{posX, posY}, Coord{velX, velY},
		}
		robots = append(robots, robot)
	}

	return robots
}

func part1(robots []Robot) {
	for range part1Seconds {
		for i := range robots {
			robot := robots[i]
			robot.position = getNewRobotPosition(robot)
			robots[i] = robot
		}
	}
	a, b, c, d := countRobotsInQuadrants(robots)
	saftetyFactor := a * b * c * d
	fmt.Printf("After %d seconds the safety factor is: %d\n", part1Seconds, saftetyFactor)
}

func part2(robots []Robot) {
	lsf, t := 0, 0
	for i := range part2Seconds {
		for j := range robots {
			robot := robots[j]
			robot.position = getNewRobotPosition(robot)
			robots[j] = robot
		}
		a, b, c, d := countRobotsInQuadrants(robots)
		safetyFactor := a * b * c * d
		if safetyFactor != 0 && (safetyFactor < lsf || lsf == 0) {
			lsf, t = safetyFactor, i+1
			fmt.Printf("----- POTENTIAL PICTURE @ %d seconds -----", t)
			printBathroom(robots)
		}
	}

	fmt.Printf("After %d seconds, the lowest observed safety factor was %d at %d seconds.\n", part2Seconds, lsf, t)
	fmt.Println("This is the likliest time at which the robots have made a picture of a Christmas tree.")
}

func getNewRobotPosition(robot Robot) Coord {
	x, y := robot.position.X, robot.position.Y
	vX, vY := robot.velocity.X, robot.velocity.Y
	newX, newY := (x+vX)%gridX, (y+vY)%gridY
	if newX < 0 {
		abs := aocutils.Abs(newX)
		newX = maxX - (abs - 1)
	}
	if newY < 0 {
		abs := aocutils.Abs(newY)
		newY = maxY - (abs - 1)
	}
	return Coord{newX, newY}
}

func countRobotsInQuadrants(robots []Robot) (int, int, int, int) {
	a, b, c, d := 0, 0, 0, 0
	for _, robot := range robots {
		x, y := robot.position.X, robot.position.Y
		midX, midY := (gridX / 2), (gridY / 2)
		if x < midX && y < midY {
			a++
		}
		if x > midX && y < midY {
			b++
		}
		if x < midX && y > midY {
			c++
		}
		if x > midX && y > midY {
			d++
		}

	}
	return a, b, c, d
}

func printBathroom(robots []Robot) {
	bathroom := [gridY][gridX]int{}
	for _, robot := range robots {
		x, y := robot.position.X, robot.position.Y
		num := bathroom[y][x] + 1
		bathroom[y][x] = num
	}
	for y := range maxY {
		for x := range maxX {
			if bathroom[y][x] != 0 {
				fmt.Print(bathroom[y][x])
			} else {
				fmt.Print(".")
			}
		}
		println()
	}
}
