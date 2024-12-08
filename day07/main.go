package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Eqn struct {
	expect int
	values []int
}

type Input struct {
	data []Eqn
}

func ParseInput(lines []string) (Input, error) {
	input := Input{}

	for _, line := range lines {
		parts := strings.Split(line, " ")

		v, err := strconv.Atoi(strings.Trim(parts[0], ":"))
		if err != nil {
			return Input{}, nil
		}

		eq := Eqn{
			expect: v,
		}
		for _, item := range parts[1:] {
			v, err := strconv.Atoi(item)
			if err != nil {
				return Input{}, nil
			}
			eq.values = append(eq.values, v)
		}

		input.data = append(input.data, eq)
	}

	return input, nil
}

func total(values []int, ops []int) int {
	sum := values[0]

	for idx := 1; idx < len(values); idx++ {
		switch ops[idx-1] {
		case 0:
			sum = sum + values[idx]
		case 1:
			sum = sum * values[idx]
		case 2:
			v := fmt.Sprintf("%d%d", sum, values[idx])
			sum, _ = strconv.Atoi(v)
		}
	}

	return sum
}

func walk(eq Eqn, ops []int, base int) bool {
	t := total(eq.values, ops)
	// fmt.Println("TRY ", eq.expect, eq.values, ops, t)
	if t == eq.expect {
		// fmt.Println("  GOOD ", ops)
		return true
	}

	nops := append([]int{}, ops...)
	nops[0] += 1
	for idx := 0; nops[idx] == base; idx += 1 {
		if idx == len(ops)-1 {
			return false
		}
		nops[idx] = 0
		nops[idx+1] += 1
	}

	return walk(eq, nops, base)
}

func canWork(eq Eqn, base int) bool {
	var ops []int

	for i := 0; i < len(eq.values)-1; i++ {
		ops = append(ops, 0)
	}

	return walk(eq, ops, base)
}

func PartOneSolution(input Input) (int, error) {
	sum := 0

	for _, eq := range input.data {
		if canWork(eq, 2) {
			sum += eq.expect
		}
	}

	return sum, nil
}

func PartTwoSolution(input Input) (int, error) {
	sum := 0

	for _, eq := range input.data {
		if canWork(eq, 3) {
			sum += eq.expect
		}
	}

	return sum, nil
}

func main() {
	fd, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	timeStart := time.Now()
	input, err := ParseInput(lines)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Build input (%.2fms)\n", float64(time.Since(timeStart).Microseconds())/1000)

	timeStart = time.Now()
	values, err := PartOneSolution(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1 (%.2fms): %v\n", float64(time.Since(timeStart).Microseconds())/1000, values)

	timeStart = time.Now()
	values, err = PartTwoSolution(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 2 (%.2fms): %v\n", float64(time.Since(timeStart).Microseconds())/1000, values)
}
