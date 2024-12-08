package main

import (
	"bufio"
	"fmt"
	"log"
	"maps"
	"os"
	"strconv"
)

func main() {
	lines := readFile("input.txt")

	antennas, c := readLine(lines)

	mapAntennas := groupAntennas(antennas)
	mapBoundaries := <-c

	countOfAntinodes := 0

	uniqueAntinodes := make(map[string]Position)

	for sameAntennasKind := range maps.Values(mapAntennas) {
		antinodes := getAntinodesStep2(sameAntennasKind, mapBoundaries)
		for _, antinode := range antinodes {
			_, ok := uniqueAntinodes[antinode.String()]

			if !ok {
				uniqueAntinodes[antinode.String()] = antinode
			}
		}

		countOfAntinodes += len(antinodes)
	}

	fmt.Println("Number of antinodes", len(uniqueAntinodes))

}

func readFile(path string) <-chan []byte {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	out := make(chan []byte)

	go func() {
		defer file.Close()
		for scanner.Scan() {
			text := scanner.Bytes()
			out <- text
		}
		close(out)
	}()

	return out
}

type Antenna struct {
	kind     byte
	position Position
}

type Position struct {
	x, y int
}

func (p Position) String() string {
	return strconv.Itoa(p.x) + ":" + strconv.Itoa(p.y)
}

type InputBoundaries struct {
	lines, columns int
}

func readLine(lines <-chan []byte) (<-chan Antenna, <-chan InputBoundaries) {
	out := make(chan Antenna)
	mapBoundaries := make(chan InputBoundaries)

	lineIndex := 0

	go func() {
		columns := 0
		for line := range lines {
			if len(line) == 0 {
				break
			}

			columns = len(line)
			for i, c := range line {
				if c != '.' {
					out <- Antenna{kind: c, position: Position{x: i, y: lineIndex}}
				}
			}
			lineIndex++
		}

		close(out)

		mapBoundaries <- InputBoundaries{lineIndex - 1, columns - 1}
		close(mapBoundaries)
	}()

	return out, mapBoundaries
}

func groupAntennas(antennas <-chan Antenna) map[byte][]Antenna {
	out := make(map[byte][]Antenna)

	for antenna := range antennas {
		a, ok := out[antenna.kind]
		if !ok {
			arr := make([]Antenna, 0)
			arr = append(arr, antenna)
			out[antenna.kind] = arr
		} else {
			b := append(a, antenna)
			out[antenna.kind] = b
		}
	}

	return out
}

func getAntinodes(a []Antenna, b InputBoundaries) []Position {
	antinodes := make([]Position, 0)

	for i, antenna := range a {
		for j := i + 1; j < len(a); j++ {
			antenna2 := a[j]
			x1 := 2*antenna.position.x - antenna2.position.x
			y1 := 2*antenna.position.y - antenna2.position.y

			x2 := 2*antenna2.position.x - antenna.position.x
			y2 := 2*antenna2.position.y - antenna.position.y

			antinode1 := Position{x1, y1}
			antinode2 := Position{x2, y2}

			if isInBoundaries(antinode1, b) {
				antinodes = append(antinodes, antinode1)
			}

			if isInBoundaries(antinode2, b) {
				antinodes = append(antinodes, antinode2)
			}
		}

	}

	return antinodes
}

func getAntinodesStep2(a []Antenna, b InputBoundaries) []Position {
	antinodes := make([]Position, 0)

	for i, antenna := range a {
		for j := i + 1; j < len(a); j++ {
			vec := Position{antenna.position.x - a[j].position.x, antenna.position.y - a[j].position.y}

			k := 0

			for {
				antinode1 := Position{antenna.position.x - k*vec.x, antenna.position.y - k*vec.y}
				antinode2 := Position{antenna.position.x + k*vec.x, antenna.position.y + k*vec.y}

				one, two := false, false

				if isInBoundaries(antinode1, b) {
					antinodes = append(antinodes, antinode1)
					one = true
				}

				if isInBoundaries(antinode2, b) {
					antinodes = append(antinodes, antinode2)
					two = true
				}

				k++

				if !one && !two {
					break
				}
			}
		}
	}

	return antinodes
}

func isInBoundaries(position Position, boundaries InputBoundaries) bool {
	return position.x >= 0 && position.x <= boundaries.columns && position.y >= 0 && position.y <= boundaries.lines
}
