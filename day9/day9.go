package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Disk []int
type File struct {
	id, length, startIdx int
}
type Files []File
type Spaces map[int]int

func main() {
	input := readInput("input.txt")
	disk := parseDisk(input)
	part1(disk)
	disk = parseDisk(input)
	part2(disk)
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

func part1(disk Disk) {
	compacted := compactDisk(disk)
	checksum := calculateChecksum(&compacted)
	fmt.Printf("The checksum after compaction is %d\n", checksum)
}

func part2(disk Disk) {
	defraged := defragDisk(disk)
	checksum := calculateChecksum(&defraged)
	fmt.Printf("The checksum after defragging is %d\n", checksum)
}

func compactDisk(disk Disk) Disk {
	for spaceAvailable(&disk) {
		a := getLastFileBlockIndex(&disk)
		b := getFirstEmptyBlockIndex(&disk)
		disk[a], disk[b] = disk[b], disk[a]
	}

	return disk
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

func defragDisk(disk Disk) Disk {
	files := analyzeFiles(&disk)
	spaces := analyzeSpaces(&disk)

	for i := len(files) - 1; i > 0; i-- {
		file := files[i]
		space_idx := getFirstAvailableSpaceIdx(&spaces, file.length)
		if space_idx == -1 || space_idx > file.startIdx {
			continue
		}
		moveFile(disk, file, space_idx)
		spaces = analyzeSpaces(&disk)

	}

	return disk
}

func analyzeFiles(disk *Disk) Files {
	files := make(Files, 0)

	counts := make(map[int]int, 0)
	for _, v := range *disk {
		if v != -1 {
			counts[v]++
		}
	}

	for id, count := range counts {
		idx := getFileStartIdx(disk, id)
		files = append(files, File{id, count, idx})
	}

	return sortFiles(files)
}

func sortFiles(files Files) Files {
	for range files {
		swapped := false
		for j := 0; j < len(files)-1; j++ {
			if files[j].id > files[j+1].id {
				files[j], files[j+1] = files[j+1], files[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return files
}

func getFileStartIdx(disk *Disk, id int) int {
	for i, v := range *disk {
		if v == id {
			return i
		}
	}

	return -1
}

func analyzeSpaces(disk *Disk) Spaces {
	spaces := make(Spaces, 0)
	startIdx := -1
	count := 0

	for i, v := range *disk {
		if v == -1 {
			if startIdx == -1 {
				startIdx = i
			}
			count++
		} else {
			if startIdx == -1 {
				continue
			}
			spaces[startIdx] = count
			startIdx = -1
			count = 0
		}
	}

	return spaces
}

func getFirstAvailableSpaceIdx(spaces *Spaces, size int) int {
	minIdx := -1
	for idx, count := range *spaces {
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

func moveFile(disk Disk, file File, idx int) {
	for i := 0; i < file.length; i++ {
		if file.id == 7 {
		}
		disk[idx+i], disk[file.startIdx+i] = disk[file.startIdx+i], disk[idx+i]
	}
}
