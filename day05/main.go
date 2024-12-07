package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Rules map[[2]int]int

type Input struct {
	rules Rules
	books [][]int
}

func intSplit(str, sep string) ([]int, error) {
	var result []int

	for _, item := range strings.Split(str, sep) {
		v, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}

	return result, nil
}

func ParseInput(lines []string) (Input, error) {
	input := Input{
		rules: Rules{},
	}

	state := "|"
	for _, line := range lines {
		if line == "" {
			state = ","
			continue
		}
		parts, err := intSplit(line, state)
		if err != nil {
			return Input{}, err
		}
		switch state {
		case "|":
			input.rules[[2]int{parts[0], parts[1]}] = -1
			input.rules[[2]int{parts[1], parts[0]}] = 1
		case ",":
			input.books = append(input.books, parts)
		}
	}

	return input, nil
}

func mkSorter(rules Rules) func(a, b int) int {
	return func(a, b int) int {
		return rules[[2]int{a, b}]
	}
}

func PartOneSolution(input Input) (int, error) {
	sum := 0

	sorter := mkSorter(input.rules)
	for _, row := range input.books {
		if slices.IsSortedFunc(row, sorter) {
			sum += row[len(row)/2]
		}
	}

	return sum, nil
}

func PartTwoSolution(input Input) (int, error) {
	sum := 0

	sorter := mkSorter(input.rules)
	for _, row := range input.books {
		pages := row[:]
		if !slices.IsSortedFunc(pages, sorter) {
			slices.SortFunc(pages, sorter)

			sum += pages[len(pages)/2]
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
