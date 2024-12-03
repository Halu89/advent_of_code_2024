package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"sync/atomic"
)

func main() {
	lines := readFile("input.txt")
	operands := parseLines(lines)

	result := atomic.Uint64{}
	for operand := range operands {
		for _, o := range operand {
			fmt.Printf("%d * %d\n", o.a, o.b)
			result.Add(uint64(o.a * o.b))
		}
	}

	fmt.Println("Result: ", result)

}

func readFile(path string) <-chan string {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	out := make(chan string)
	go func() {
		defer file.Close()
		for scanner.Scan() {
			text := scanner.Text()
			out <- text
		}
		close(out)
	}()

	return out
}

type Operand struct {
	a int
	b int
}

func (o Operand) String() string {
	return fmt.Sprintf("%d * %d", o.a, o.b)
}

func parseLines(in <-chan string) <-chan []Operand {
	out := make(chan []Operand)
	collectionEnabled := true

	go func() {
		for line := range in {
			operands, collect := parseLine(line, collectionEnabled)
			collectionEnabled = collect
			if operands != nil {
				out <- operands
			}
		}
		close(out)
	}()

	return out
}

func parseLine(line string, collectionEnabled bool) ([]Operand, bool) {
	regex := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)|do\\(\\)|don't\\(\\)")
	numberRegex := regexp.MustCompile("\\d{1,3}")

	matches := regex.FindAllString(line, -1)

	enableCollection := collectionEnabled

	if len(matches) == 0 {
		return nil, true
	}

	operands := make([]Operand, 0)

	for _, match := range matches {
		if match == "do()" {
			enableCollection = true
			continue
		}
		if match == "don't()" {
			enableCollection = false
			continue
		}

		numbers := numberRegex.FindAllString(match, -1)

		if len(numbers) != 2 {
			log.Fatal("Expected 2 numbers in the match")
		}

		nums := make([]int, len(numbers))

		for i, number := range numbers {
			num, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
			nums[i] = num
		}
		if enableCollection {
			operands = append(operands, Operand{a: nums[0], b: nums[1]})
		}
	}

	return operands, enableCollection
}
