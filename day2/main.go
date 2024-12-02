package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	levels   []int
	dampener bool
}

func main() {
	lines := readFile("input.txt")
	reports := parseLines(lines)

	safeReports := 0
	for report := range reports {
		if report.isSafeWithDampen() {
			safeReports++
		}
	}

	log.Println("Safe reports: ", safeReports)
}

func NewReport(levels []int) Report {
	return Report{levels: levels}
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

func parseLines(in <-chan string) <-chan Report {
	out := make(chan Report)

	go func() {
		for line := range in {
			if line == "" {
				continue
			}
			report := parseLine(line)
			out <- report
		}
		close(out)
	}()

	return out
}

func parseLine(line string) Report {
	numbers := strings.Split(line, " ")

	levels := make([]int, 0)

	for _, number := range numbers {
		num, parseError := strconv.Atoi(number)
		if parseError != nil {
			log.Fatal(parseError)
		}
		levels = append(levels, num)
	}

	return Report{levels: levels, dampener: true}
}

func (r *Report) isSafe() bool {
	isIncrementing := r.levels[0] < r.levels[1]
	safeLevel := 3

	for i := 0; i < len(r.levels)-1; i++ {
		if isIncrementing {
			if r.levels[i] >= r.levels[i+1] {
				return false
			}
			if r.levels[i+1]-r.levels[i] > safeLevel {
				return false
			}
		} else {
			if r.levels[i] <= r.levels[i+1] {
				return false
			}
			if r.levels[i]-r.levels[i+1] > safeLevel {
				return false
			}
		}
	}

	return true
}

func (r *Report) isSafeWithDampen() bool {
	if r.isSafe() {
		return true
	}

	for i := 0; i < len(r.levels); i++ {
		newLevel := make([]int, len(r.levels))
		copy(newLevel, r.levels)
		dampenedReport := NewReport(newLevel)
		dampenedReport.dampen(i)
		if dampenedReport.isSafe() {
			return true
		}
	}

	return false
}

func (r *Report) dampen(index int) {
	r.dampener = false
	// Remove the element at index

	if index == len(r.levels)-1 {
		r.levels = r.levels[:index]
		return
	}

	if index == 0 {
		r.levels = r.levels[1:]
		return
	}

	r.levels = append(r.levels[:index], r.levels[index+1:]...)
}
