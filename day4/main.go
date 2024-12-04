package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

type Direction int

const (
	SE Direction = iota
	S
	SW
	W
	NW
	N
	NE
	E
)

func (d Direction) String() string {
	return [...]string{"SE", "S", "SW", "W", "NW", "N", "NE", "E"}[d]
}

type Grid struct {
	tiles [][]byte
}

func main() {
	// open file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	grid := make([][]byte, 0)
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	foundWords := walkGridStep2(grid)
	//for _, word := range foundWords {
	//	//log.Printf("Found XMAS at %d, %d in direction %s", word.startLocation[0], word.startLocation[1], word.direction)
	//}

	//log.Printf("Found %d words", len(foundWords))
	log.Printf("Found %d words", foundWords)
}

type FoundWord struct {
	startLocation []int
	direction     Direction
}

func walkGrid(grid [][]byte) []FoundWord {
	foundWords := make([]FoundWord, 0)

	for i, row := range grid {
		for j, cell := range row {
			if cell == 'X' {
				foundWords = append(foundWords, checkAllDirections(grid, i, j)...)
			}
		}
	}
	return foundWords
}

func checkAllDirections(grid [][]byte, x, y int) []FoundWord {
	foundWords := make([]FoundWord, 0)
	startLocation := []int{x, y}

	for _, direction := range []Direction{S, SE, SW, N, NE, NW, W, E} {
		if checkDirection(grid, x, y, direction, 'M') {
			x, y, err := getLocation(grid, x, y, direction)
			if err != nil {
				continue
			}

			if checkDirection(grid, x, y, direction, 'A') {
				x, y, err := getLocation(grid, x, y, direction)
				if err != nil {
					continue
				}

				if checkDirection(grid, x, y, direction, 'S') {
					_, _, err := getLocation(grid, x, y, direction)
					if err != nil {
						continue
					}
					foundWords = append(foundWords, FoundWord{startLocation: startLocation, direction: direction})
				}
			}
		}
	}
	return foundWords
}

func checkDirection(grid [][]byte, x int, y int, direction Direction, target byte) bool {

	i, j, err := getLocation(grid, x, y, direction)

	if err != nil {
		return false
	}

	if grid[i][j] == target {
		return true
	}

	return false
}

func getLocation(grid [][]byte, x, y int, direction Direction) (int, int, error) {
	newX := -1
	newY := -1
	switch direction {
	case SE:
		newX = x + 1
		newY = y + 1
	case SW:
		newX = x + 1
		newY = y - 1
	case N:
		newX = x - 1
		newY = y
	case NE:
		newX = x - 1
		newY = y + 1
	case NW:
		newX = x - 1
		newY = y - 1
	case W:
		newX = x
		newY = y - 1
	case E:
		newX = x
		newY = y + 1
	case S:
		newX = x + 1
		newY = y
	}

	if newX < 0 || newX >= len(grid[0]) || newY < 0 || newY >= len(grid) {
		return -1, -1, errors.New("out of bounds")
	}

	return newX, newY, nil
}

func walkGridStep2(grid [][]byte) int {
	foundWords := 0

	for i, row := range grid {
		for j, cell := range row {
			if cell == 'A' {
				//foundWords = append(foundWords, checkAllDirections(grid, i, j)...)
				foundWords += checkAllDirectionsStep2(grid, i, j)
			}
		}
	}
	return foundWords
}

type Corner struct {
	x, y int
	Direction
	value byte
}

func checkAllDirectionsStep2(grid [][]byte, x, y int) int {
	corners := make([]Corner, 0)

	for _, direction := range []Direction{SE, SW, NE, NW} {
		x2, y2, err := getLocation(grid, x, y, direction)
		if err != nil {
			continue
		}
		corners = append(corners, Corner{x: x2, y: y2, Direction: direction, value: grid[x2][y2]})
	}
	fmt.Println(corners)

	if len(corners) < 4 {
		return 0
	}

	// Matching corners
	if (corners[0].value == 'M' && corners[3].value == 'S') || (corners[0].value == 'S' && corners[3].value == 'M') {
		if (corners[1].value == 'M' && corners[2].value == 'S') || (corners[1].value == 'S' && corners[2].value == 'M') {
			return 1
		}
	}
	return 0
}
