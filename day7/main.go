package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()

	lines := readFile("input.txt")
	equations := processLines(lines)

	operators := []Operator{Plus, Multiply, Concatenate}
	validEquations := processValidEquations(equations, operators)

	result := sumValidEquations(validEquations)
	fmt.Println(result)

	fmt.Println("Execution time: ", time.Since(startTime))
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
	Plus        Operator = '+'
	Multiply    Operator = '*'
	Concatenate Operator = '|'
)

func enumerate(size int, operators []Operator) [][]Operator {
	if size == 1 {
		var result [][]Operator
		for _, operator := range operators {
			result = append(result, []Operator{operator})
		}
		return result
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
	case Concatenate:
		return concatenate(a, b)
	}

	panic("Unknown operator")
}

func concatenate(a, b int) int {
	result, err := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func hash(size int, operators []Operator) string {
	return fmt.Sprintf("%d %v", size, operators)
}

func isEquationValid(equation Equation, operators []Operator, memo *map[string][][]Operator) bool {
	operatorEnumerations, ok := (*memo)[hash(len(equation.operands)-1, operators)]
	if !ok {
		operatorEnum := enumerate(len(equation.operands)-1, operators)
		(*memo)[hash(len(equation.operands)-1, operators)] = operatorEnum
	}

	for _, operatorList := range operatorEnumerations {
		result := equation.operands[0]

		for i, operator := range operatorList {
			result = operator.doOperate(result, equation.operands[i+1])
			if result > equation.target {
				break
			}
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
			if isEquationValid(eq, operators, &map[string][][]Operator{}) {
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
