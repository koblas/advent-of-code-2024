package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"
)

type Pos [2]int
type Grid map[Pos]string

var dirs = []Pos{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

type QueueItem struct {
	position Pos
	score    int
}

type PQueue struct {
	q []QueueItem
}

type Input struct {
	points []Pos
}

func (p Pos) move(dir Pos) Pos {
	return Pos{p[0] + dir[0], p[1] + dir[1]}
}

func (p Pos) eq(t Pos) bool {
	return p[0] == t[0] && p[1] == t[1]
}

func ParseInput(lines []string) (Input, error) {
	input := Input{}

	for _, line := range lines {
		p := strings.Split(line, ",")
		x, err := strconv.Atoi(p[0])
		if err != nil {
			return Input{}, err
		}
		y, err := strconv.Atoi(p[1])
		if err != nil {
			return Input{}, err
		}
		input.points = append(input.points, Pos{x, y})
	}

	return input, nil
}

func (queue *PQueue) append(dp QueueItem) {
	pos := 0
	for ; pos < len(queue.q); pos += 1 {
		if queue.q[pos].score >= dp.score {
			break
		}
	}

	nq := append([]QueueItem{}, queue.q[0:pos]...)
	nq = append(nq, dp)
	nq = append(nq, queue.q[pos:]...)

	queue.q = nq
}

func (queue *PQueue) pop() QueueItem {
	item := queue.q[0]
	queue.q = queue.q[1:]

	return item
}

func (queue *PQueue) isEmpty() bool {
	return len(queue.q) == 0
}

func draw(grid Grid, visited []Pos) {
	maxX := 0
	maxY := 0
	for p := range grid {
		maxX = max(maxX, p[0])
		maxY = max(maxY, p[1])
	}
	maxY += 1
	maxX += 1

	for y := range maxY {
		for x := range maxX {
			p := Pos{x, y}
			if grid[p] == "." && slices.Contains(visited, p) {
				fmt.Print("O")
			} else if grid[p] == "#" {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func walk(grid Grid, end Pos) int {
	pqueue := PQueue{}
	distances := map[Pos]int{}

	insert := func(pos Pos, score int) {
		distances[pos] = score
		pqueue.append(QueueItem{position: pos, score: score})
	}

	insert(Pos{0, 0}, 0)

	for !pqueue.isEmpty() {
		item := pqueue.pop()

		if item.position.eq(end) {
			return item.score
		}

		// fmt.Println("HERE ", item.position)

		for _, dir := range dirs {
			np := item.position.move(dir)
			// fmt.Println("TRY", np)
			score := item.score + 1
			if v, found := grid[np]; !found || v == "#" {
				// fmt.Println("   NOPE", np, grid[np])
				continue
			}
			if dist, found := distances[np]; found && dist <= score {
				// fmt.Println("   SCORE", np, score)
				continue
			}
			// fmt.Println("   OK", np)

			insert(np, score)
		}
	}

	// draw(grid, slices.Collect(maps.Keys(distances)))

	return -1
}

func makeGrid(size int, steps int, input Input) Grid {
	grid := Grid{}
	for x := range size + 1 {
		for y := range size + 1 {
			grid[Pos{x, y}] = "."
		}
	}

	for idx := range steps {
		grid[input.points[idx]] = "#"
	}

	return grid
}

func PartOneSolution(input Input) (int, error) {
	size := 70
	steps := 1024
	// test size
	if testing.Testing() {
		size = 6
		steps = 12
	}

	grid := makeGrid(size, steps, input)

	// draw(grid, []Pos{})

	sum := walk(grid, Pos{size, size})

	return sum, nil
}

func PartTwoSolution(input Input) (string, error) {
	size := 70
	steps := 1024
	// test size
	if testing.Testing() {
		size = 6
		steps = 12
	}

	end := Pos{size, size}

	l := steps
	r := len(input.points)
	for l != r {
		var grid Grid

		mid := (l + r + 1) / 2
		grid = makeGrid(size, mid, input)

		if walk(grid, end) < 0 {
			r = mid - 1
		} else {
			l = mid
		}
	}

	p := input.points[l]
	return fmt.Sprintf("%d,%d", p[0], p[1]), nil
}

func main() {
	infile := "input.txt"
	if len(os.Args) == 2 {
		infile = os.Args[1]
	}
	fd, err := os.Open(infile)
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
	values2, err := PartTwoSolution(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 2 (%.2fms): %v\n", float64(time.Since(timeStart).Microseconds())/1000, values2)
}
