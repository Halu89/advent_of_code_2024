package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Read the input file
	sumOfDistances := calculateDistance("input.txt")

	similarities := step2("input.txt")

	log.Println("Sum of distances: ", sumOfDistances)

	log.Println("Sum of similarities: ", similarities)
}

func calculateDistance(path string) int {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	list1 := make([]int, 0)
	list2 := make([]int, 0)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			continue
		}

		int1, int2 := parseLine(text)

		list1 = append(list1, int1)
		list2 = append(list2, int2)
	}

	// Sort the lists
	sort.Ints(list1)
	sort.Ints(list2)

	sumOfDistances := 0
	for i := 0; i < len(list1); i++ {
		sumOfDistances += int(math.Abs(float64(list2[i] - list1[i])))
	}
	return sumOfDistances
}

func step2(path string) int {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	list1 := make([]int, 0)
	list2 := make([]int, 0)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			continue
		}

		int1, int2 := parseLine(text)

		list1 = append(list1, int1)
		list2 = append(list2, int2)
	}

	return calculateSimilarities(list1, list2)
}

func calculateSimilarities(list1 []int, list2 []int) int {
	occurrences := countOccurrences(list2)

	similarities := 0
	for i := 0; i < len(list1); i++ {
		if occurrences[list1[i]] > 0 {
			similarities += list1[i] * occurrences[list1[i]]
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

func readFile(path string) chan string {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	out := make(chan string)

	go func() {
		for scanner.Scan() {
			text := scanner.Text()
			out <- text
		}
		close(out)
	}()

	return out
}
