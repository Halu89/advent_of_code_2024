package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"

	"strings"
)

type OrderingRule struct {
	smaller int
	bigger  int
}

type Report struct {
	report  []ComparablePage
	isValid bool
}

type ComparablePage struct {
	value         int
	orderingRules map[int][]int
}

func main() {
	lines := readFile("input.txt")
	rules, reports := parseLines(lines)

	comparisonMap := makeComparisonMap(rules)

	index := 0

	sumOfMiddlePages := 0
	sumOfUnorderedReports := 0

	for report := range reports {
		parsedReport := processReport(&comparisonMap, report)

		log.Printf("Report %d (%v): %v", index, report, parsedReport.isValid)
		index++
		if parsedReport.isValid {
			middlePage := parsedReport.middlePage()
			log.Printf("Middle page: %d", middlePage)
			sumOfMiddlePages += parsedReport.middlePage()
		} else {
			parsedReport.sort()
			middlePage := parsedReport.middlePage()
			log.Printf("Middle page: %d", middlePage)
			sumOfUnorderedReports += middlePage
		}
	}

	log.Printf("Sum of middle pages: %d", sumOfMiddlePages)
	log.Printf("Sum of unordered reports: %d", sumOfUnorderedReports)
}

func (r *Report) middlePage() int {
	middleIndex := len(r.report) / 2
	return r.report[middleIndex].value
}

func (r *Report) sort() {
	slices.SortFunc(r.report, sortFunction)
}

func processReport(comparisonMap *map[int][]int, report []int) Report {
	pages := make([]ComparablePage, len(report))

	for i, page := range report {
		comparablePage := makeComparablePage(comparisonMap, page)
		pages[i] = comparablePage
	}

	return Report{pages, slices.IsSortedFunc(pages, sortFunction)}
}

// Compare pages
func sortFunction(a, b ComparablePage) int {
	if a.orderingRules[a.value] != nil {
		if slices.Contains(a.orderingRules[a.value], b.value) {

			if b.orderingRules[b.value] != nil &&
				slices.Contains(b.orderingRules[b.value], a.value) {
				panic("Inconsistent ordering rules")
			} else {
				return -1
			}
		}
	}

	if b.orderingRules[b.value] != nil {
		return 1
	}

	return 0
}

func parseLines(in <-chan string) (<-chan OrderingRule, <-chan []int) {
	rules := make(chan OrderingRule)
	reports := make(chan []int)

	isOrderingRule := true

	go func() {
		for line := range in {
			if line == "" {
				isOrderingRule = false
				close(rules)
				continue
			}

			if isOrderingRule {
				orderingRule := parseRule(line)
				rules <- orderingRule
			} else {
				report := parseReport(line)
				reports <- report
			}
		}
		close(reports)
	}()

	return rules, reports
}

func parseReport(line string) []int {
	reports := strings.Split(line, ",")

	intReports := make([]int, len(reports))
	for i, report := range reports {
		intReport, parseError := strconv.Atoi(report)
		if parseError != nil {
			log.Fatal(parseError)
		}
		intReports[i] = intReport
	}

	return intReports
}

func parseRule(line string) OrderingRule {
	pages := strings.Split(line, "|")

	if len(pages) != 2 {
		log.Fatal("Invalid rule")
	}

	smallerInt, parseError := strconv.Atoi(pages[0])
	if parseError != nil {
		log.Fatal(parseError)
	}

	biggerInt, parseError := strconv.Atoi(pages[1])
	if parseError != nil {
		log.Fatal(parseError)
	}

	return OrderingRule{smallerInt, biggerInt}
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

func makeComparisonMap(rules <-chan OrderingRule) map[int][]int {
	comparisonMap := make(map[int][]int)

	for rule := range rules {
		if _, ok := comparisonMap[rule.smaller]; ok {
			comparisonMap[rule.smaller] = append(comparisonMap[rule.smaller], rule.bigger)
		} else {
			comparisonMap[rule.smaller] = make([]int, 0)
			comparisonMap[rule.smaller] = append(comparisonMap[rule.smaller], rule.bigger)
		}
	}

	return comparisonMap
}

func makeComparablePage(rules *map[int][]int, page int) ComparablePage {
	return ComparablePage{page, *rules}
}
