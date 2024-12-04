package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

type Input struct {
	left  []int
	right []int
}

func ParseInput(lines []string) (Input, error) {
	re := regexp.MustCompile(`\s+`)

	input := Input{}
	for _, line := range lines {
		parts := re.Split(line, -1)

		val, _ := strconv.Atoi(parts[0])
		input.left = append(input.left, val)
		val, _ = strconv.Atoi(parts[1])
		input.right = append(input.right, val)
	}

	return input, nil
}

func PartOneSolution(input Input) (int, error) {
	sort.Ints(input.left)
	sort.Ints(input.right)

	dist := 0
	for idx, l := range input.left {
		val := l - input.right[idx]
		dist += max(val, -val)
	}

	return dist, nil
}

func PartTwoSolution(input Input) (int, error) {
	counts := map[int]int{}

	for _, r := range input.right {
		counts[r] += 1
	}

	sum := 0
	for _, l := range input.left {
		sum += l * counts[l]
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
