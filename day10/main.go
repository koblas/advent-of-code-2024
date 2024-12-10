package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"
)

type Pos [2]int

type Input struct {
	data   map[Pos]int
	bounds Pos
	starts []Pos
}

var dirs = [4]Pos{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func (p Pos) move(d Pos) Pos {
	return Pos{p[0] + d[0], p[1] + d[1]}
}

func ParseInput(lines []string) (Input, error) {
	input := Input{
		data:   map[Pos]int{},
		bounds: Pos{len(lines[0]), len(lines)},
	}

	for y, line := range lines {
		for x, ch := range line {
			v, err := strconv.Atoi(string(ch))
			if err != nil {
				return Input{}, err
			}
			input.data[Pos{x, y}] = v
			if v == 0 {
				input.starts = append(input.starts, Pos{x, y})
			}
		}
	}

	return input, nil
}

func walk(input Input, pos Pos, depth int, visited *[]Pos) int {
	if depth == 9 {
		if visited != nil {
			if slices.Contains(*visited, pos) {
				return 0
			}
			*visited = append(*visited, pos)
		}
		return 1
	}
	depth += 1

	count := 0

	for _, d := range dirs {
		np := pos.move(d)
		v, found := input.data[np]
		if !found || v != depth {
			continue
		}

		count += walk(input, np, depth, visited)
	}

	return count
}

func walkOne(input Input) int {
	count := 0
	for _, start := range input.starts {
		visited := []Pos{}
		score := walk(input, start, 0, &visited)
		count += score
	}
	return count
}

func walkTwo(input Input) int {
	count := 0
	for _, start := range input.starts {
		score := walk(input, start, 0, nil)
		count += score
	}
	return count
}

func PartOneSolution(input Input) (int, error) {
	result := walkOne(input)

	return result, nil
}

func PartTwoSolution(input Input) (int, error) {
	result := walkTwo(input)

	return result, nil
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
