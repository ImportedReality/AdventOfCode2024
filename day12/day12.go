package main

import (
	"fmt"
	"maps"
	"slices"

	"github.com/ImportedReality/aocutils"
	mapset "github.com/deckarep/golang-set/v2"
)

const filename = "input.txt"

type Plant struct {
	plantType  string
	loc        aocutils.Coordinate
	neighbours map[aocutils.Direction]Plant
}
type Region []Plant

var grid aocutils.Grid[string]
var garden aocutils.Grid[Plant]
var regions []Region

func main() {
	grid = aocutils.ReadGrid(filename, "")
	readGarden()
	findRegions()
	total := 0
	for _, region := range regions {
		total += calculateRegionPrice(region)
	}
	fmt.Printf("Total price: %d\n", total)
}

func readGarden() {
	garden = make(aocutils.Grid[Plant], 0)
	for y, row := range grid {
		gardenRow := make([]Plant, 0)
		for x, col := range row {
			plant := Plant{
				col,
				aocutils.Coordinate{X: x, Y: y},
				make(map[aocutils.Direction]Plant),
			}
			gardenRow = append(gardenRow, plant)
		}
		garden = append(garden, gardenRow)
	}
	for _, row := range garden {
		for _, plant := range row {
			findNeighbors(plant)
		}
	}
}

func findNeighbors(plant Plant) {
	for dir := aocutils.N; dir <= aocutils.W; dir += 2 {
		offset := aocutils.Offsets[dir]
		x, y := plant.loc.X+offset.X, plant.loc.Y+offset.Y
		toCheck := aocutils.Coordinate{X: x, Y: y}
		if aocutils.InBounds(garden, toCheck) && garden[y][x].plantType == plant.plantType {
			plant.neighbours[dir] = garden[y][x]
		}
	}
}

func printPlant(plant Plant) {
	fmt.Printf("%s (%d, %d) %d\n", plant.plantType, plant.loc.X, plant.loc.Y, len(plant.neighbours))
}

func findRegions() []Region {
	regions = make([]Region, 0)
	for _, row := range garden {
		for _, plant := range row {
			if !plantInExistingRegion(plant, regions) {
				coords := fillRegion(plant)
				region := regionFromCoords(coords)
				regions = append(regions, region)
			}
		}
	}

	return regions
}

func fillRegion(plant Plant) mapset.Set[aocutils.Coordinate] {
	visited := mapset.NewSet[aocutils.Coordinate]()
	queue := aocutils.Stack[Plant]{}
	p := plant
	queue = queue.Push(p)

	for len(queue) > 0 {
		queue, p = queue.Shift()
		if !visited.Contains(p.loc) {
			queue = queue.Push(slices.Collect(maps.Values(p.neighbours))...)
		}
		visited.Add(p.loc)
	}

	return visited
}

func plantInExistingRegion(plant Plant, regions []Region) bool {
	for _, region := range regions {
		for _, p := range region {
			if plant.loc == p.loc {
				return true
			}
		}
	}
	return false
}

func regionFromCoords(coords mapset.Set[aocutils.Coordinate]) Region {
	region := make(Region, 0)

	for c := range coords.Iter() {
		plant := plantFromCoord(c)
		region = append(region, plant)
	}

	return region
}

func plantFromCoord(coord aocutils.Coordinate) Plant {
	return garden[coord.Y][coord.X]
}

func calculateRegionPrice(region Region) int {
	area := len(region)
	perimeter := 0
	for _, plant := range region {
		perimeter += 4 - len(plant.neighbours)
	}
	return area * perimeter
}
