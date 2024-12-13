package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Machine struct {
	A, B, Prize Vec
}

type Vec struct {
	x, y int64
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	machines := make([]Machine, 0)

	for scanner.Scan() {
		machine, ok := readMachine(scanner, true)

		if ok {
			machines = append(machines, machine)
		}
	}

	totalTokenPrice := 0
	for _, machine := range machines {
		totalTokenPrice += int(calculatePrice(machine))
	}

	fmt.Println("Result: ", totalTokenPrice)
}

func readMachine(scanner *bufio.Scanner, part2 bool) (Machine, bool) {
	machine := Machine{}

	digitRegexp := regexp.MustCompile(`\d+`)

	for i := 0; i < 3; i++ {
		text := scanner.Text()
		if i == 0 {
			digits := digitRegexp.FindAllString(text, -1)
			machine.A = Vec{x: mustConvert(digits[0]), y: mustConvert(digits[1])}
		}
		if i == 1 {
			digits := digitRegexp.FindAllString(text, -1)
			machine.B = Vec{x: mustConvert(digits[0]), y: mustConvert(digits[1])}
		} else {
			digits := digitRegexp.FindAllString(text, -1)
			machine.Prize = Vec{x: mustConvert(digits[0]), y: mustConvert(digits[1])}

			if part2 {
				machine.Prize.x += 10000000000000
				machine.Prize.y += 10000000000000
			}
		}

		scanner.Scan()
	}

	return machine, true
}

func hasSolution(machine Machine) (Vec, bool) {
	var a, b int64

	det := machine.A.x*machine.B.y - machine.A.y*machine.B.x

	aNum := machine.Prize.x*machine.B.y - machine.Prize.y*machine.B.x
	bNum := machine.A.x*machine.Prize.y - machine.A.y*machine.Prize.x

	if a = aNum / det; a < 0 || a*det != aNum {
		return Vec{}, false
	}
	if b = bNum / det; b < 0 || b*det != bNum {
		return Vec{}, false
	}
	return Vec{a, b}, true
}

func calculatePrice(machine Machine) int64 {
	prize, ok := hasSolution(machine)

	if !ok {
		return 0
	}
	return prize.x*3 + prize.y
}

func (m *Machine) String() string {
	return fmt.Sprintf("A: %v, B: %v, Prize: %v", m.A, m.B, m.Prize)
}

func (v *Vec) String() string {
	return fmt.Sprintf("(%d, %d)", v.x, v.y)
}

func mustConvert(s string) int64 {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return int64(n)
}
