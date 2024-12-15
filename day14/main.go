package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Pos [2]int
type Robot struct {
	pos      Pos
	velocity Pos
}

type Input struct {
	robots []Robot
}

func (p Pos) move(d Pos) Pos {
	return Pos{p[0] + d[0], p[1] + d[1]}
}

var lineRE = regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

func ParseInput(lines []string) (Input, error) {
	input := Input{}

	for _, line := range lines {
		parts := lineRE.FindStringSubmatch(line)

		if parts == nil {
			return input, errors.New("bad input")
		}

		px, _ := strconv.Atoi(parts[1])
		py, _ := strconv.Atoi(parts[2])
		vx, _ := strconv.Atoi(parts[3])
		vy, _ := strconv.Atoi(parts[4])

		input.robots = append(input.robots, Robot{
			pos:      Pos{px, py},
			velocity: Pos{vx, vy},
		})
	}

	return input, nil
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func iterate(pos Pos, vel Pos, grid Pos, count int) Pos {
	px := (pos[0] + vel[0]*count)
	py := (pos[1] + vel[1]*count)

	px = mod(px, grid[0])
	py = mod(py, grid[1])

	return Pos{px, py}
}

func PartOneSolution(input Input) (int, error) {
	sum := 0

	grid := Pos{101, 103}
	// grid := Pos{11, 7} // test grid

	midX := grid[0] / 2
	midY := grid[1] / 2

	counts := [4]int{}
	for _, robot := range input.robots {
		pos := iterate(robot.pos, robot.velocity, grid, 100)

		switch {
		case pos[0] < midX && pos[1] < midY:
			counts[0] += 1
		case pos[0] > midX && pos[1] < midY:
			counts[1] += 1
		case pos[0] < midX && pos[1] > midY:
			counts[2] += 1
		case pos[0] > midX && pos[1] > midY:
			counts[3] += 1
		}
	}

	sum = counts[0] * counts[1] * counts[2] * counts[3]

	// too low: 115463700
	return sum, nil
}

func PartTwoSolution(input Input) (int, error) {
	sum := 0

	grid := Pos{101, 103}

	for ; sum < 100_000; sum += 1 {
		seen := map[Pos]bool{}
		for _, robot := range input.robots {
			pos := iterate(robot.pos, robot.velocity, grid, sum)

			seen[pos] = true
		}

		adj := 0
		// var doprint = false
		for y := range grid[1] {
			for x := range grid[0] {
				pos := Pos{x, y}
				if seen[pos] && seen[pos.move(Pos{1, 0})] {
					adj += 1
				} else {
					adj = 0
				}
				if adj > 10 {
					return sum, nil
				}
			}
		}

		// if doprint {
		// 	for y := range grid[1] {
		// 		for x := range grid[0] {
		// 			if seen[Pos{x, y}] {
		// 				fmt.Print("X")
		// 			} else {
		// 				fmt.Print(".")
		// 			}
		// 		}
		// 		fmt.Println("")
		// 	}
		// 	fmt.Println("")
		// }
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
