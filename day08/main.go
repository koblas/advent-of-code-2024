package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"time"
)

type Pos [2]int

type Input struct {
	nodes  map[rune][]Pos
	bounds Pos
}

func ParseInput(lines []string) (Input, error) {
	input := Input{
		nodes:  map[rune][]Pos{},
		bounds: Pos{len(lines[0]), len(lines)},
	}

	for y, line := range lines {
		for x, ch := range line {
			if ch != '.' {
				input.nodes[ch] = append(input.nodes[ch], Pos{x, y})
			}
		}
	}

	return input, nil
}

func addOne(p Pos, bounds Pos, output []Pos) ([]Pos, bool) {
	if p[0] < 0 || p[0] >= bounds[0] || p[1] < 0 || p[1] >= bounds[1] {
		return output, false
	}

	return append(output, p), true
}

func calc(na Pos, nb Pos, bounds Pos) []Pos {
	dx := na[0] - nb[0]
	dy := na[1] - nb[1]

	p0 := Pos{na[0] + dx, na[1] + dy}
	p1 := Pos{nb[0] - dx, nb[1] - dy}

	// fmt.Println(" TRY NA:", na, "NB:", nb, "D:", Pos{dx, dy}, "P0:", p0, "P1:", p1)

	output, _ := addOne(p0, bounds, []Pos{})
	output, _ = addOne(p1, bounds, output)

	return output
}

func calcTwo(na Pos, nb Pos, bounds Pos) []Pos {
	dx := na[0] - nb[0]
	dy := na[1] - nb[1]

	var output []Pos
	var ok bool
	p := na
	for {
		p = Pos{p[0] + dx, p[1] + dy}

		output, ok = addOne(p, bounds, output)
		if !ok {
			break
		}
	}

	p = nb
	for {
		p = Pos{p[0] - dx, p[1] - dy}

		output, ok = addOne(p, bounds, output)
		if !ok {
			break
		}
	}

	return output
}

func dump(input Input, output map[Pos]bool) {
	for y := range input.bounds[1] {
		for x := range input.bounds[0] {
			p := Pos{x, y}
			if output[p] {
				fmt.Print("#")
				continue
			}

			found := '.'
			for ch, row := range input.nodes {
				if slices.Contains(row, p) {
					found = ch
					break
				}
			}

			fmt.Print(string(found))
		}
		fmt.Println("")
	}
}

func place(input Input, partTwo bool) map[Pos]bool {
	output := map[Pos]bool{}

	for _, places := range input.nodes {
		for idx, na := range places {
			if partTwo {
				output[na] = true
			}
			for _, nb := range places[idx+1:] {
				if partTwo {
					output[nb] = true
				}
				if partTwo {
					for _, n := range calcTwo(na, nb, input.bounds) {
						output[n] = true
					}
				} else {
					for _, n := range calc(na, nb, input.bounds) {
						output[n] = true
					}
				}

			}
		}
	}

	// dump(input, output)

	return output
}

func PartOneSolution(input Input) (int, error) {
	vals := place(input, false)

	return len(vals), nil
}

func PartTwoSolution(input Input) (int, error) {
	vals := place(input, true)

	// 947 too high

	return len(vals), nil
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
