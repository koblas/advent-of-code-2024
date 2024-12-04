package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Pos [2]int
type Input struct {
	maxX int
	maxY int
	data map[Pos]rune
}

var dirs = []Pos{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
	{-1, 1},
	{1, -1},
	{-1, -1},
	{1, 1},
}

const xmas = "XMAS"

func ParseInput(lines []string) (Input, error) {
	input := Input{
		maxY: len(lines),
		maxX: len(lines[0]),
		data: map[Pos]rune{},
	}

	for y, row := range lines {
		for x, char := range row {
			input.data[Pos{x, y}] = char
		}
	}

	return input, nil
}

func checkDir(input Input, x, y int, dir Pos, word string) bool {
	for offset, ch := range word {
		if input.data[Pos{x + offset*dir[0], y + offset*dir[1]}] != ch {
			return false
		}
	}

	return true
}

func checkOne(input Input, x, y int) int {
	count := 0
	for _, d := range dirs {
		if checkDir(input, x, y, d, "XMAS") {
			count += 1
		}
	}

	return count
}

func checkTwo(input Input, x, y int) bool {
	d00 := Pos{1, 1}
	d10 := Pos{-1, 1}

	if !checkDir(input, x, y, d00, "MAS") && !checkDir(input, x, y, d00, "SAM") {
		return false
	}

	return checkDir(input, x+2, y, d10, "MAS") || checkDir(input, x+2, y, d10, "SAM")
}

func PartOneSolution(input Input) (int, error) {
	count := 0

	for y := range input.maxY {
		for x := range input.maxX {
			count += checkOne(input, x, y)
		}
	}

	return count, nil
}

func PartTwoSolution(input Input) (int, error) {
	count := 0

	for y := range input.maxY {
		for x := range input.maxX {
			if checkTwo(input, x, y) {
				count += 1
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
