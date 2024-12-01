package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	firstCol, secondCol := parseInput("input.txt")

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		sumOfDistances := calculateDistance(firstCol, secondCol)
		fmt.Println("Sum of distances: ", sumOfDistances)
	}()

	go func() {
		defer wg.Done()
		similarities := calculateSimilarities(firstCol, secondCol)
		fmt.Println("Similarities: ", similarities)
	}()

	wg.Wait()
}

func calculateSimilarities(a []int, b []int) int {
	occurrences := countOccurrences(b)

	similarities := 0
	for i := 0; i < len(a); i++ {
		if occurrences[a[i]] > 0 {
			similarities += a[i] * occurrences[a[i]]
		}
	}

	return similarities
}

func countOccurrences(list []int) map[int]int {
	occurrences := make(map[int]int)
	for _, value := range list {
		if _, ok := occurrences[value]; ok {
			occurrences[value]++
		} else {
			occurrences[value] = 1
		}
	}
	return occurrences
}

func parseLine(line string) (int, int) {
	numbers := strings.Split(line, "   ")
	int1, parseError := strconv.Atoi(numbers[0])
	if parseError != nil {
		log.Fatal(parseError)
	}

	int2, parseError := strconv.Atoi(numbers[1])
	if parseError != nil {
		log.Fatal(parseError)
	}

	return int1, int2
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

type Vector struct {
	x int
	y int
}

func readLines(in <-chan string) <-chan Vector {
	out := make(chan Vector)

	go func() {
		for line := range in {
			int1, int2 := parseLine(line)
			out <- Vector{x: int1, y: int2}
		}
		close(out)
	}()

	return out
}

func calculateDistance(a []int, b []int) int {
	sumOfDistances := 0
	for i := 0; i < len(a); i++ {
		sumOfDistances += int(math.Abs(float64(b[i] - a[i])))
	}
	return sumOfDistances
}

func parseInput(path string) ([]int, []int) {
	in := readFile(path)
	vectors := readLines(in)

	firstCol := make([]int, 0)
	secondCol := make([]int, 0)

	for vector := range vectors {
		firstCol = append(firstCol, vector.x)
		secondCol = append(secondCol, vector.y)
	}

	sort.Ints(firstCol)
	sort.Ints(secondCol)

	return firstCol, secondCol
}
