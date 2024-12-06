package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type Direction string

const (
	N Direction = "N"
	E           = "E"
	S           = "S"
	W           = "W"
)

type Position struct {
	x            int
	y            int
	positionType PositionType
}

type PositionType string

const (
	Empty    PositionType = "."
	Obstacle              = "#"
	Guard                 = "^"
	Walked                = "X"
)

func (d Direction) getDirection() (int, int) {
	switch d {
	case N:
		return -1, 0
	case E:
		return 0, 1
	case S:
		return 1, 0
	case W:
		return 0, -1
	}
	return 0, 0
}

func main() {
	grid := readInput()

	// Step 1
	startTime := time.Now()

	parsedGrid := parseGrid(grid)
	parsedGrid.walk()

	walkedPositions := make([]Position, 0)
	for _, row := range parsedGrid.positions {
		for _, position := range row {
			if position.positionType == Walked {
				walkedPositions = append(walkedPositions, position)
			}
		}
	}
	log.Println("Walked positions: ", len(walkedPositions))
	log.Println("Elapsed time: ", time.Since(startTime))

	// Step 2

	loopCounter := 0

	for _, position := range walkedPositions {
		newGrid := parseGrid(grid)
		newGrid.addObstacle(position.x, position.y)

		looping := newGrid.walk()

		if looping {
			loopCounter++
		}
	}

	log.Println("Loop counter: ", loopCounter)

	elapsed := time.Since(startTime)
	log.Printf("Execution time: %s", elapsed)
}

func (g *Grid) addObstacle(colIndex, rowIndex int) PositionType {
	originalPositionType := g.positions[rowIndex][colIndex].positionType

	g.positions[rowIndex][colIndex] = Position{x: colIndex, y: rowIndex, positionType: Obstacle}

	return originalPositionType
}

func (g *Grid) removeObstacle(colIndex, rowIndex int, originalPositionType PositionType) {
	g.positions[rowIndex][colIndex].positionType = originalPositionType
}

func readInput() [][]byte {
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

	return grid
}

func parseGrid(grid [][]byte) Grid {
	positions := make([][]Position, 0)
	guardPosition := GuardPosition{-1, -1, N}

	for rowIndex := 0; rowIndex < len(grid); rowIndex++ {
		row := make([]Position, 0)

		for colIndex := 0; colIndex < len(grid[rowIndex]); colIndex++ {
			positionType := getPositionType(grid[rowIndex][colIndex])

			if positionType == Guard {
				guardPosition = GuardPosition{x: colIndex, y: rowIndex, Direction: N}
				positionType = Empty
			}

			row = append(row, Position{x: colIndex, y: rowIndex, positionType: positionType})
		}
		positions = append(positions, row)
	}

	if guardPosition.x == -1 || guardPosition.y == -1 {
		log.Fatal("Guard position not found")
	}

	walkedPositions := make(map[string]struct{})

	return Grid{positions: positions, guardPosition: guardPosition, walkedPositions: walkedPositions}
}

type Grid struct {
	positions       [][]Position
	guardPosition   GuardPosition
	walkedPositions map[string]struct{}
}

func (g GuardPosition) String() string {
	return fmt.Sprintf("%d:%d %s", g.x, g.y, g.Direction)
}

type GuardPosition struct {
	x int
	y int
	Direction
}

func (g GuardPosition) Equal(other GuardPosition) bool {
	return g.x == other.x && g.y == other.y && g.Direction == other.Direction
}

func (g GuardPosition) Compare(other GuardPosition) int {
	if g.Equal(other) {
		return 0
	}

	return 1
}

func getPositionType(b byte) PositionType {
	switch b {
	case '#':
		return Obstacle
	case '^':
		return Guard
	case 'X':
		return Walked
	default:
		return Empty
	}
}

func (g *Grid) String() string {
	out := "\n"
	for _, row := range g.positions {
		for _, position := range row {
			out += " " + string(position.positionType)
		}
		out += "\n"
	}
	return out
}

func (g *Grid) walk() bool {
	g.positions[g.guardPosition.y][g.guardPosition.x].positionType = Walked
	g.walkedPositions[g.guardPosition.String()] = struct{}{}

	for {
		err := g.walkStep()

		if err != nil {
			if string(err.Error()) == "in a loop" {
				return true
			}
			break
		}
	}

	return false
}

func (g *Grid) walkStep() error {
	x, y := g.guardPosition.x, g.guardPosition.y
	direction := g.guardPosition.Direction

	newX, newY := x, y

	switch direction {
	case N:
		newY--
	case E:
		newX++
	case S:
		newY++
	case W:
		newX--
	}

	if newX < 0 || newX >= len(g.positions[x]) || newY < 0 || newY >= len(g.positions) {
		return errors.New("out of bounds")
	}

	if g.positions[newY][newX].positionType == Obstacle {
		g.guardPosition.Direction = turnRight(direction)

		if has(g.walkedPositions, g.guardPosition) {
			return errors.New("in a loop")
		} else {
			g.walkedPositions[g.guardPosition.String()] = struct{}{}
		}

		return nil
	}

	if g.positions[newY][newX].positionType == Empty || g.positions[newY][newX].positionType == Walked {
		g.guardPosition.x = newX
		g.guardPosition.y = newY

		g.positions[newY][newX].positionType = Walked

		if has(g.walkedPositions, g.guardPosition) {
			return errors.New("in a loop")
		} else {
			g.walkedPositions[g.guardPosition.String()] = struct{}{}
		}

		return nil
	}

	return errors.New("unknown location")
}

func has(positions map[string]struct{}, guardPosition GuardPosition) bool {
	_, ok := positions[guardPosition.String()]
	return ok
}

func turnRight(direction Direction) Direction {
	switch direction {
	case N:
		return E
	case E:
		return S
	case S:
		return W
	case W:
		return N
	}
	return N
}
