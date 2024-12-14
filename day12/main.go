package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Pos [2]int

type Input struct {
	data   map[Pos]rune
	bounds Pos
}

var dirs = []Pos{
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
		data:   map[Pos]rune{},
		bounds: Pos{len(lines[0]), len(lines)},
	}

	for y, line := range lines {
		for x, ch := range line {
			input.data[Pos{x, y}] = ch
		}
	}

	return input, nil
}

func innerWalk(input Input, follow rune, p Pos, visited map[Pos]bool) int {
	perimiter := 0
	visited[p] = true
	for _, d := range dirs {
		np := p.move(d)
		if visited[np] {
			continue
		}
		if ch, found := input.data[np]; !found || ch != follow {
			perimiter += 1
			continue
		}
		perimiter += innerWalk(input, follow, np, visited)
	}

	return perimiter
}

func walk(input Input, pos Pos) (int, map[Pos]bool) {
	visited := map[Pos]bool{
		pos: true,
	}

	perimiter := innerWalk(input, input.data[pos], pos, visited)

	cost := len(visited) * perimiter

	return cost, visited
}

func walkTwo(input Input, pos Pos) (int, map[Pos]bool) {
	visited := map[Pos]bool{
		pos: true,
	}

	_ = innerWalk(input, input.data[pos], pos, visited)

	// moves := []Pos{
	// 	{0, 0},
	// 	{1, 0},
	// 	{0, 1},
	// 	{1, 1},
	// }

	// btoi := func(vals ...bool) int {
	// 	sum := 0
	// 	for _, v := range vals {
	// 		if v {
	// 			sum += 1
	// 		}
	// 	}
	// 	return sum
	// }

	const debug = false

	/*
		vertexes := map[Pos]bool{}
		for p := range visited {
			for _, d := range moves {
				vertexes[p.move(d)] = true
			}
		}

		sides := 0
		for v := range vertexes {
			a00 := visited[v.move(Pos{0, 0})]
			a01 := visited[v.move(Pos{0, 1})]
			a10 := visited[v.move(Pos{1, 0})]
			a11 := visited[v.move(Pos{1, 1})]

			cnt := btoi(a00, a01, a10, a11)
			fmt.Println("  ", string(input.data[pos]), cnt)

			switch cnt {
			case 3, 1:
				sides += 1
			case 2:
				if (a00 && a11) || (a10 && a01) {
					sides += 2
				}
			}
		}
	*/
	sides := 0

	for p := range visited {
		n := visited[p.move(Pos{0, -1})]
		s := visited[p.move(Pos{0, 1})]
		w := visited[p.move(Pos{-1, 0})]
		e := visited[p.move(Pos{1, 0})]
		ne := visited[p.move(Pos{1, -1})]
		nw := visited[p.move(Pos{-1, -1})]
		se := visited[p.move(Pos{1, 1})]
		sw := visited[p.move(Pos{-1, 1})]

		if !n && !e || n && e && !ne {
			sides += 1
		}
		if !n && !w || n && w && !nw {
			sides += 1
		}
		if !s && !e || s && e && !se {
			sides += 1
		}
		if !s && !w || s && w && !sw {
			sides += 1
		}
	}

	/*
		for p := range visited {
			for _, m := range moves {
				v := p.move(m)
				np := v.move(Pos{-1, -1})
				a00 := visited[np.move(Pos{0, 0})]
				a01 := visited[np.move(Pos{0, 1})]
				a10 := visited[np.move(Pos{1, 0})]
				a11 := visited[np.move(Pos{1, 1})]

				cnt := btoi(a00, a01, a10, a11)
				if debug {
					fmt.Println("POS = ", p, m, cnt)
				}
				switch cnt {
				case 4, 0:
					break
				case 1, 3:
					vertexes[v] = true
				case 2:
					if (a00 && a11) || (a10 && a01) {
						vertexes[v] = true
					}
				}
				if debug {
					fmt.Println("  ", a00, np.move(Pos{0, 0}), a10, np.move(Pos{1, 0}))
					fmt.Println("  ", a01, np.move(Pos{0, 1}), a11, np.move(Pos{1, 1}))
					fmt.Println("  ", vertexes[v])
				}

			}
		}
	*/

	// if debug {
	// 	output := [][]rune{}
	// 	for range input.bounds[1]*2 + 1 {
	// 		row := []rune{}
	// 		for range input.bounds[0]*2 + 1 {
	// 			row = append(row, ' ')
	// 		}

	// 		output = append(output, row)
	// 	}

	// 	for p := range visited {
	// 		output[p[1]*2+1][p[0]*2+1] = input.data[pos]
	// 	}

	// 	for p, v := range vertexes {
	// 		if v {
	// 			output[p[1]*2][p[0]*2] = '+'
	// 		}
	// 	}

	// 	for _, row := range output {
	// 		fmt.Println(string(row))
	// 	}
	// }

	cost := len(visited) * sides

	// fmt.Println(string(input.data[pos]), "AREA", len(visited), "SIDES", sides, "VTX", len(vertexes), "COST", cost, "DATA", vertexes)

	return cost, visited
}

func PartOneSolution(input Input) (int, error) {
	sum := 0
	visited := map[Pos]bool{}

	for x := range input.bounds[0] {
		for y := range input.bounds[1] {
			p := Pos{x, y}
			if visited[p] {
				continue
			}

			cost, visit := walk(input, p)

			for pos := range visit {
				visited[pos] = true
			}

			sum += cost
		}
	}

	return sum, nil
}

func PartTwoSolution(input Input) (int, error) {
	sum := 0
	visited := map[Pos]bool{}

	for x := range input.bounds[0] {
		for y := range input.bounds[1] {
			p := Pos{x, y}
			if visited[p] {
				continue
			}

			cost, visit := walkTwo(input, p)

			for pos := range visit {
				visited[pos] = true
			}

			sum += cost
		}
	}

	// Too low: 869144
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
