package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Input struct {
	levels [][]int
}

func ParseInput(lines []string) (Input, error) {
	input := Input{}
	for _, line := range lines {
		var level []int
		for _, value := range strings.Split(line, " ") {
			v, err := strconv.Atoi(value)
			if err != nil {
				return Input{}, err
			}
			level = append(level, v)
		}
		input.levels = append(input.levels, level)
	}

	return input, nil
}

func isSafe(level []int) bool {
	var psign int
	for idx := 0; idx < len(level)-1; idx++ {
		delta := level[idx] - level[idx+1]
		sign := 1
		if delta < 0 {
			sign = -1
		}
		if max(delta, -delta) > 3 || delta == 0 {
			return false
		}
		if idx != 0 && psign != sign {
			return false
		}
		psign = sign
	}

	return true
}

func PartOneSolution(input Input) (int, error) {
	count := 0

	for _, level := range input.levels {
		if isSafe(level) {
			count += 1
		}
	}

	return count, nil
}

func without(value []int, idx int) []int {
	var output []int
	output = append(output, value[:idx]...)
	return append(output, value[idx+1:]...)

}

func PartTwoSolution(input Input) (int, error) {
	count := 0

	// 560 too high
	// 291 to low
	for _, level := range input.levels {
		if isSafe(level) {
			count += 1
		} else {
			for idx := range level {
				attempt := without(level, idx)
				if isSafe(attempt) {
					count += 1
					break
				}
			}
		}
	}

	return count, nil
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
