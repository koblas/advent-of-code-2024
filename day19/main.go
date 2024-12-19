package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Input struct {
	patterns []string

	towels []string
}

func ParseInput(lines []string) (Input, error) {
	input := Input{}

	partTwo := false
	for _, line := range lines {
		if line == "" {
			partTwo = true
			continue
		}
		if partTwo {
			input.towels = append(input.towels, line)
		} else {
			for _, v := range strings.Split(line, ",") {
				input.patterns = append(input.patterns, strings.TrimSpace(v))
			}
		}
	}

	return input, nil
}

func countPossible(cache map[string]int, towel string, patterns []string) int {
	if towel == "" {
		return 1
	}
	if value, found := cache[towel]; found {
		return value
	}
	cnt := 0
	for _, p := range patterns {
		if sub, found := strings.CutPrefix(towel, p); found {
			cnt += countPossible(cache, sub, patterns)
		}
	}
	cache[towel] = cnt

	return cnt
}

func PartOneSolution(input Input) (int, error) {
	sum := 0

	cache := map[string]int{}
	for _, towel := range input.towels {
		if countPossible(cache, towel, input.patterns) != 0 {
			sum += 1
		}
	}

	return sum, nil
}

func PartTwoSolution(input Input) (int, error) {
	sum := 0

	cache := map[string]int{}
	for _, towel := range input.towels {
		sum += countPossible(cache, towel, input.patterns)
	}

	return sum, nil
}

func main() {
	infile := "input.txt"
	if len(os.Args) == 2 {
		infile = os.Args[1]
	}
	fd, err := os.Open(infile)
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
	values2, err := PartTwoSolution(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 2 (%.2fms): %v\n", float64(time.Since(timeStart).Microseconds())/1000, values2)
}
