package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Item struct {
	val1, val2 int
	capture    bool
}

type Input struct {
	muls []Item
}

func ParseInput(lines []string) (Input, error) {
	re := regexp.MustCompile(`(?:mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\))`)

	input := Input{}
	capture := true
	for _, line := range lines {
		for _, res := range re.FindAllStringSubmatch(line, -1) {
			switch {
			case res[0] == "don't()":
				capture = false
			case res[0] == "do()":
				capture = true
			case strings.HasPrefix(res[0], "mul("):
				v1, err := strconv.ParseInt(res[1], 10, 0)
				if err != nil {
					return Input{}, err
				}
				v2, err := strconv.ParseInt(res[2], 10, 0)
				if err != nil {
					return Input{}, err
				}

				input.muls = append(input.muls, Item{
					val1:    int(v1),
					val2:    int(v2),
					capture: capture,
				})
			}
		}
	}

	return input, nil
}

func PartOneSolution(input Input) (int, error) {
	sum := 0
	for _, item := range input.muls {
		sum += item.val1 * item.val2
	}
	return sum, nil
}

func PartTwoSolution(input Input) (int, error) {
	sum := 0
	for _, item := range input.muls {
		if item.capture {
			sum += item.val1 * item.val2
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
