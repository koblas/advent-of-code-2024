package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"time"
)

type Pos [2]int

func (p Pos) move(d Pos) Pos {
	return Pos{p[0] + d[0], p[1] + d[1]}
}

var Dirs = []Pos{
	{0, -1}, // UP
	{1, 0},  // RIGHT
	{0, 1},  // DOWN
	{-1, 0}, // LEFT
}

type Map map[Pos]map[int]bool

type Input struct {
	board Map
	start Pos
}

func ParseInput(lines []string) (Input, error) {
	input := Input{
		board: Map{},
	}

	for y, line := range lines {
		for x, ch := range line {
			p := Pos{x, y}
			input.board[p] = map[int]bool{}
			switch ch {
			case '#':
				input.board[p][-1] = true
			case '^':
				input.start = p
			case '.':
				// nothing
			}
		}
	}

	return input, nil
}

func copyMap(input Input) Map {
	output := Map{}

	for k, v := range input.board {
		output[k] = maps.Clone(v)
	}

	return output
}

func dump(input Map) {
	pos := Pos{}

	for {
		pos[0] = 0
		if _, found := input[pos]; !found {
			break
		}
		for {
			vals, found := input[pos]
			if !found {
				break
			}
			pos[0] += 1

			evens := vals[0] || vals[2]
			odds := vals[1] || vals[3]

			switch {
			case len(vals) == 0:
				fmt.Print(".")
			case vals[-1]:
				fmt.Print("#")
			case vals[-2]:
				fmt.Print("O")
			case evens && odds:
				fmt.Print("+")
			case evens:
				fmt.Print("|")
			case odds:
				fmt.Print("-")
			}
		}
		fmt.Println("")
		pos[1] += 1
	}

	fmt.Println("")
}

func walk(input Map, pos Pos) bool {
	var dir = 0

	for {
		if _, found := input[pos][dir]; found {
			// in a loop
			return true
		}
		input[pos][dir] = true

		// dump(input)

		next := pos.move(Dirs[dir])

		if vals, found := input[next]; !found {
			return false
		} else if vals[-1] || vals[-2] {
			dir = (dir + 1) % 4
		} else {
			pos = next
		}
	}
}

func findCanidates(board Map, start Pos) []Pos {
	output := []Pos{}

	for pos, vals := range board {
		if pos[0] == start[0] && pos[1] == start[1] {
			continue
		}

		if len(vals) == 0 {
			continue
		}
		if vals[-1] || vals[-2] {
			continue
		}
		output = append(output, pos)
	}

	return output
}

func PartOneSolution(input Input) (int, error) {
	scratch := copyMap(input)
	walk(scratch, input.start)
	canidates := findCanidates(scratch, Pos{-1, -1})
	return len(canidates), nil
}

func PartTwoSolution(input Input) (int, error) {
	sum := 0

	scratch := copyMap(input)
	walk(scratch, input.start)

	canidates := findCanidates(scratch, input.start)

	for _, possible := range canidates {
		scratch := copyMap(input)
		scratch[possible][-2] = true
		if walk(scratch, input.start) {
			sum += 1
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
