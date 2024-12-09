package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const InputPath = "./input.txt"

func main() {
	data, err := os.ReadFile(InputPath)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")

	partOne(strings.Split(lines[0], ""))
	partTwo(strings.Split(lines[0], ""))

}

func partOne(line []string) {
	solution := int64(0)
	diskMap := make([]string, 0)
	for i := range line {
		if i%2 == 0 {
			id := i / 2
			fileSpace, _ := strconv.ParseInt(line[i], 10, 32)
			for j := 0; j < int(fileSpace); j++ {
				diskMap = append(diskMap, fmt.Sprintf("%d", id))
			}
		} else {
			freeSpace, _ := strconv.ParseInt(line[i], 10, 32)
			for j := 0; j < int(freeSpace); j++ {
				diskMap = append(diskMap, ".")
			}
		}
	}
	for leftPointer, rightPointer := 0, len(diskMap)-1; leftPointer <= rightPointer; {
		if diskMap[leftPointer] == "." && diskMap[rightPointer] != "." {
			diskMap[leftPointer], diskMap[rightPointer] = diskMap[rightPointer], diskMap[leftPointer]
			leftPointer++
			rightPointer--
		} else {
			if diskMap[leftPointer] != "." {
				leftPointer++
			} else if diskMap[rightPointer] != "." {
				rightPointer--
			} else if diskMap[leftPointer] == "." && diskMap[rightPointer] == "." {
				rightPointer--
			}
		}
	}
	for i := range diskMap {
		if diskMap[i] != "." {
			x, _ := strconv.ParseInt(diskMap[i], 10, 32)
			solution += int64(i) * x
		}
	}
	fmt.Printf("Part One: %d\n", solution)
}

func partTwo(line []string) {
	solution := int64(0)
	diskMap := make([]string, 0)
	files := make([][]int64, 0)
	spaces := make([][]int64, 0)
	for i := range line {
		if i%2 == 0 {
			id := i / 2
			fileSpace, _ := strconv.ParseInt(line[i], 10, 32)
			files = append(files, []int64{int64(len(diskMap)), int64(len(diskMap)-1) + fileSpace})
			for j := 0; j < int(fileSpace); j++ {
				diskMap = append(diskMap, fmt.Sprintf("%d", id))
			}
		} else {
			space, _ := strconv.ParseInt(line[i], 10, 32)
			spaces = append(spaces, []int64{int64(len(diskMap)), int64(len(diskMap)-1) + space})
			for j := 0; j < int(space); j++ {
				diskMap = append(diskMap, ".")
			}
		}
	}
	for rightPointer := len(files) - 1; rightPointer >= 0; rightPointer-- {
		for leftPointer := 0; leftPointer < len(spaces); leftPointer++ {
			space := spaces[leftPointer]
			if space[1] < files[rightPointer][1] && space[1]-space[0] >= files[rightPointer][1]-files[rightPointer][0] {
				for k := space[0]; k <= space[0]+files[rightPointer][1]-files[rightPointer][0]; k++ {
					diskMap[k] = fmt.Sprintf("%d", rightPointer)
				}
				for k := files[rightPointer][0]; k <= files[rightPointer][1]; k++ {
					diskMap[k] = "."
				}
				space[0] = space[0] + files[rightPointer][1] - files[rightPointer][0] + 1
				break
			}
		}

	}
	for i := range diskMap {
		if diskMap[i] != "." {
			x, _ := strconv.ParseInt(diskMap[i], 10, 32)
			solution += int64(i) * x
		}
	}
	fmt.Printf("Part Two: %d\n", solution)
}
