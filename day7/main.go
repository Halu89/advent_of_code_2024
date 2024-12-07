package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readFile("input.txt")
	equations := processLines(lines)

	operators := []Operator{Plus, Multiply}
	validEquations := processValidEquations(equations, operators)

	result := sumValidEquations(validEquations)
	fmt.Println(result)
}

type Equation struct {
	target   int
	operands []int
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

func parseEquation(line string) Equation {
	operands := make([]int, 0)

	a := strings.Split(line, ": ")
	target, err := strconv.Atoi(a[0])
	if err != nil {
		log.Fatal(err)
	}

	b := strings.Split(a[1], " ")
	for _, operand := range b {
		op, err := strconv.Atoi(operand)
		if err != nil {
			log.Fatal(err)
		}
		operands = append(operands, op)
	}

	return Equation{target: target, operands: operands}
}

func processLines(lines <-chan string) <-chan Equation {
	out := make(chan Equation)

	go func() {
		for line := range lines {
			out <- parseEquation(line)
		}
		close(out)
	}()

	return out
}

type Operator byte

const (
	Plus     Operator = '+'
	Multiply Operator = '*'
)

func enumerate(size int, operators []Operator) [][]Operator {
	if size == 1 {
		return [][]Operator{{Plus}, {Multiply}}
	}

	var result [][]Operator

	for _, operator := range operators {
		for _, list := range enumerate(size-1, operators) {
			result = append(result, append([]Operator{operator}, list...))
		}
	}

	return result
}

func (o *Operator) doOperate(a, b int) int {
	switch *o {
	case Plus:
		return a + b
	case Multiply:
		return a * b
	}

	panic("Unknown operator")
}

func isEquationValid(equation Equation, operators []Operator) bool {
	for _, operatorList := range enumerate(len(equation.operands)-1, operators) {
		result := equation.operands[0]

		for i, operator := range operatorList {
			result = operator.doOperate(result, equation.operands[i+1])
		}

		if result == equation.target {
			return true
		}
	}

	return false
}

func processValidEquations(equation <-chan Equation, operators []Operator) <-chan Equation {
	out := make(chan Equation)

	go func() {
		for eq := range equation {
			if isEquationValid(eq, operators) {
				out <- eq
			}
		}
		close(out)
	}()

	return out
}

func sumValidEquations(equations <-chan Equation) int {
	sum := 0
	for eq := range equations {
		sum += eq.target
	}

	return sum
}
