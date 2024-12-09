package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Disk []int
type File struct {
	id, length, startIdx int
}
type Files []File
type Spaces map[int]int

func main() {
	input := readInput("input.txt")

	start := time.Now()
	disk := parseDisk(input)
	part1(&disk)
	duration := time.Since(start)
	fmt.Printf("Part 1 execution time: %v\n", duration)

	start = time.Now()
	disk = parseDisk(input)
	part2(&disk)
	duration = time.Since(start)
	fmt.Printf("Part 2 execution time: %v\n", duration)
}

func readInput(file string) string {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	line, err := bufio.NewReader(f).ReadString('\n')
	if err != nil {
		panic(err)
	}

	return line
}

func parseDisk(input string) Disk {
	disk := make(Disk, 0)

	diskMap := make([]int, 0)
	vals := strings.Split(input, "")
	for _, val := range vals {
		if val != "\n" {
			diskMap = append(diskMap, strToInt(val))
		}

	}

	file := true
	currentId := 0
	for _, val := range diskMap {
		for i := 0; i < val; i++ {
			if file {
				disk = append(disk, currentId)
			} else {
				disk = append(disk, -1)
			}
		}
		if file {
			file = !file
			currentId++
		} else {
			file = !file
		}
	}

	return disk
}

func strToInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}

func part1(disk *Disk) {
	compactDisk(disk)
	checksum := calculateChecksum(disk)
	fmt.Printf("The checksum after compaction is %d\n", checksum)
}

func part2(disk *Disk) {
	defragDisk(disk)
	checksum := calculateChecksum(disk)
	fmt.Printf("The checksum after defragging is %d\n", checksum)
}

func compactDisk(disk *Disk) {
	for spaceAvailable(disk) {
		a := getLastFileBlockIndex(disk)
		b := getFirstEmptyBlockIndex(disk)
		(*disk)[a], (*disk)[b] = (*disk)[b], (*disk)[a]
	}
}

func spaceAvailable(disk *Disk) bool {
	space := false
	for _, v := range *disk {
		if v == -1 {
			space = true
		} else if space {
			return true
		}
	}

	return false
}

func getLastFileBlockIndex(disk *Disk) int {
	for i := len(*disk) - 1; i >= 0; i-- {
		if (*disk)[i] != -1 {
			return i
		}
	}
	panic("No file blocks found!")
}

func getFirstEmptyBlockIndex(disk *Disk) int {
	for i := 0; i < len(*disk); i++ {
		if (*disk)[i] == -1 {
			return i
		}
	}
	panic("No empty blocks found!")
}

func calculateChecksum(disk *Disk) int {
	sum := 0
	for i, id := range *disk {
		if id != -1 {
			sum += i * id
		}
	}

	return sum
}

func defragDisk(disk *Disk) {
	files, spaces := analyzeDisk(disk)

	for i := len(files) - 1; i > 0; i-- {
		file := files[i]
		space_idx := getFirstAvailableSpaceIdx(spaces, file.length)
		if space_idx == -1 || space_idx > file.startIdx {
			continue
		}
		moveFile(disk, file, space_idx)
		spaces = updateSpaces(spaces, space_idx, file.length)

	}

}

func analyzeDisk(disk *Disk) (Files, Spaces) {
	files := make(Files, 0)
	spaces := make(Spaces, 0)
	spaceIdx := -1

	for i, id := range *disk {
		if id == -1 {
			if spaceIdx == -1 {
				spaceIdx = i
			}
			spaces[spaceIdx]++
		} else {
			if spaceIdx != -1 {
				spaceIdx = -1
			}
			if len(files) < id+1 {
				files = append(files, File{id, 0, i})
			}
			files[id].length++
		}
	}

	return files, spaces
}

func updateSpaces(spaces Spaces, idx, file_size int) Spaces {
	if (spaces)[idx] == file_size {
		delete((spaces), idx)
	} else {
		size := (spaces)[idx] - file_size
		delete((spaces), idx)
		(spaces)[idx+file_size] = size
	}

	return spaces
}

func getFirstAvailableSpaceIdx(spaces Spaces, size int) int {
	minIdx := -1
	for idx, count := range spaces {
		if count >= size {
			if minIdx == -1 {
				minIdx = idx
			} else if idx < minIdx {
				minIdx = idx
			}
		}
	}

	return minIdx
}

func moveFile(disk *Disk, file File, idx int) {
	for i := 0; i < file.length; i++ {
		(*disk)[idx+i], (*disk)[file.startIdx+i] = (*disk)[file.startIdx+i], (*disk)[idx+i]
	}
}
