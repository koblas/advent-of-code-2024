package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"time"
)

type Pos [2]int
type Grid map[Pos]rune

type Input struct {
	grid  Grid
	moves []rune
	start Pos
}

func (p Pos) move(dir rune) Pos {
	switch dir {
	case '^':
		return Pos{p[0], p[1] - 1}
	case 'v':
		return Pos{p[0], p[1] + 1}
	case '<':
		return Pos{p[0] - 1, p[1]}
	case '>':
		return Pos{p[0] + 1, p[1]}
	}

	return p
}

func (p Pos) eq(t Pos) bool {
	return p[0] == t[0] && p[1] == t[1]
}

func (grid Grid) print(start Pos) {
	var mx, my int
	for p := range grid {
		mx = max(mx, p[0])
		my = max(my, p[1])
	}

	for y := range my + 1 {
		for x := range mx + 1 {
			p := Pos{x, y}
			if start.eq(p) {
				fmt.Print("@")
			} else {
				fmt.Print(string(grid[p]))
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func ParseInput(lines []string) (Input, error) {
	input := Input{
		grid: Grid{},
	}

	inGrid := true
	for y, line := range lines {
		if line == "" {
			inGrid = false
		}
		for x, ch := range line {
			if inGrid {
				if ch == '@' {
					input.start = Pos{x, y}
					input.grid[Pos{x, y}] = '.'
				} else {
					input.grid[Pos{x, y}] = ch
				}
			} else {
				input.moves = append(input.moves, ch)
			}
		}
	}

	return input, nil
}

func move(grid Grid, d rune, pos Pos) Pos {
	np := pos.move(d)

	if grid[np] == '#' {
		return pos
	}
	if grid[np] == '.' {
		return np
	}
	tp := np
	for ; grid[tp] == 'O'; tp = tp.move(d) {
		// nothing
	}
	if grid[tp] == '#' {
		return pos
	}
	if grid[tp] == '.' {
		grid[tp] = 'O'
		grid[np] = '.'
	}
	return np
}

func canPush(grid Grid, pos Pos, dir rune) bool {
	var pl Pos
	var pr Pos

	if grid[pos] == '[' {
		pl = pos.move(dir)
		pr = pos.move(dir).move('>')
	} else {
		pl = pos.move(dir).move('<')
		pr = pos.move(dir)
	}

	if grid[pl] == '#' || grid[pr] == '#' {
		return false
	}

	if grid[pl] == '[' || grid[pl] == ']' {
		if !canPush(grid, pl, dir) {
			return false
		}
	}
	if grid[pr] == '[' || grid[pr] == ']' {
		if !canPush(grid, pr, dir) {
			return false
		}
	}

	return true
}

func push(grid Grid, pos Pos, dir rune) {
	var pl Pos
	var pr Pos

	if grid[pos] == '[' {
		pl = pos
		pr = pos.move('>')
	} else {
		pl = pos.move('<')
		pr = pos
	}
	nl := pl.move(dir)
	nr := pr.move(dir)

	if grid[nl] == '[' || grid[nl] == ']' {
		push(grid, nl, dir)
	}
	if grid[nr] == '[' || grid[nr] == ']' {
		push(grid, nr, dir)
	}

	grid[pl] = '.'
	grid[pr] = '.'
	grid[pl.move(dir)] = '['
	grid[pr.move(dir)] = ']'
}

func moveGrow(grid Grid, dir rune, pos Pos) Pos {
	np := pos.move(dir)

	if grid[np] == '#' {
		return pos
	}
	if grid[np] == '.' {
		return np
	}

	if dir == '<' || dir == '>' {
		tp := np
		for ; grid[tp] == '[' || grid[tp] == ']'; tp = tp.move(dir) {
			// nothing
		}
		if grid[tp] == '#' {
			return pos
		}
		if grid[tp] == '.' {
			ch := '.'
			for p := np; !p.eq(tp); p = p.move(dir) {
				nc := grid[p]
				grid[p] = ch
				ch = nc
			}
			grid[tp] = ch
		}
		return np
	}

	pushable := false
	if grid[np] == '[' || grid[np] == ']' {
		pushable = canPush(grid, np, dir)
	}
	if !pushable {
		return pos
	}

	push(grid, np, dir)

	return np
}

func grow(grid Grid) Grid {
	output := Grid{}

	for p, ch := range grid {
		switch ch {
		case '.':
			output[Pos{p[0] * 2, p[1]}] = '.'
			output[Pos{p[0]*2 + 1, p[1]}] = '.'
		case 'O':
			output[Pos{p[0] * 2, p[1]}] = '['
			output[Pos{p[0]*2 + 1, p[1]}] = ']'
		case '#':
			output[Pos{p[0] * 2, p[1]}] = '#'
			output[Pos{p[0]*2 + 1, p[1]}] = '#'
		}
	}

	return output
}

func PartOneSolution(input Input) (int, error) {
	sum := 0

	grid := maps.Clone(input.grid)

	// grid.print(input.start)

	pos := input.start
	for _, d := range input.moves {
		pos = move(grid, d, pos)
	}

	for p, ch := range grid {
		if ch == 'O' {
			sum += p[0] + p[1]*100
		}
	}

	// grid.print(pos)

	return sum, nil
}

func PartTwoSolution(input Input) (int, error) {
	sum := 0

	grid := grow(input.grid)

	pos := Pos{input.start[0] * 2, input.start[1]}
	// fmt.Println("MOVE ", string(input.moves))
	// grid.print(pos)
	for _, d := range input.moves {
		pos = moveGrow(grid, d, pos)
		// fmt.Println("AFTER ", string(d))
		// grid.print(pos)
	}

	for p, ch := range grid {
		if ch == '[' {
			sum += p[0] + p[1]*100
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
