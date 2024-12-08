package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coordinate struct {
	x, y int
}

type Antenna struct {
	frequency string
	location  Coordinate
}

type CityMap [][]string

type Antennas map[string][]Antenna

func main() {
	cityMap := readMap("input.txt")
	antennas := locateAntennas(cityMap)
	part1(cityMap, antennas)
	part2(cityMap, antennas)
}

func readMap(file string) CityMap {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := make(CityMap, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		input = append(input, row)
	}

	return input
}

func part1(cityMap CityMap, antennas Antennas) {
	antinodes := locateAntinodes(cityMap, antennas)
	fmt.Printf("There are %d unique antinodes\n", len(antinodes))
}

func part2(cityMap CityMap, antennas Antennas) {
	antinodes := locateResonantAntinodes(cityMap, antennas)
	fmt.Printf("There are %d unique resonant antinodes\n", len(antinodes))
}

func locateAntennas(cityMap CityMap) Antennas {
	results := make(Antennas, 0)
	for y, row := range cityMap {
		for x, col := range row {
			if col != "." {
				antenna := Antenna{col, Coordinate{x, y}}
				antennas := results[antenna.frequency]
				results[antenna.frequency] = append(antennas, antenna)
			}
		}
	}

	return results
}

func locateAntinodes(cityMap CityMap, antennas Antennas) map[Coordinate]bool {
	maxX, maxY := len(cityMap[0])-1, len(cityMap)-1
	antinodes := make(map[Coordinate]bool, 0)
	for _, freq := range antennas {
		for i, antenna := range freq {
			for j := 0; j < len(freq); j++ {
				if j != i {
					offset := getOffset(antenna.location, freq[j].location)
					antinode := Coordinate{
						antenna.location.x + offset.x,
						antenna.location.y + offset.y,
					}
					if inBounds(antinode, maxX, maxY) {
						_, exists := antinodes[antinode]
						if !exists {
							antinodes[antinode] = true
						}
					}
				}
			}
		}
	}

	return antinodes
}

func locateResonantAntinodes(cityMap CityMap, antennas Antennas) map[Coordinate]bool {
	maxX, maxY := len(cityMap[0])-1, len(cityMap)-1
	antinodes := make(map[Coordinate]bool, 0)
	for _, freq := range antennas {
		for i, antenna := range freq {
			for j := 0; j < len(freq); j++ {
				if j != i {
					offset := getOffset(antenna.location, freq[j].location)
					antinode := Coordinate{antenna.location.x, antenna.location.y}
					for inBounds(antinode, maxX, maxY) {
						_, exists := antinodes[antinode]
						if !exists {
							antinodes[antinode] = true
						}
						antinode = Coordinate{
							antinode.x + offset.x,
							antinode.y + offset.y,
						}
					}
				}
			}
		}
	}

	return antinodes
}

func getOffset(a, b Coordinate) Coordinate {
	return Coordinate{
		a.x - b.x,
		a.y - b.y,
	}
}

func inBounds(c Coordinate, maxX, maxY int) bool {
	return c.x >= 0 && c.x <= maxX && c.y >= 0 && c.y <= maxY
}
