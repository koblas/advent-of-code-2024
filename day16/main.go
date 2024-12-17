package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"time"
)

type Dir string

type Pos [2]int
type Grid map[Pos]string

type DirPos struct {
	dir Dir
	pos Pos
}

type QueueItem struct {
	dirpos  DirPos
	score   int
	visited []Pos
}

type PQueue []QueueItem

type Input struct {
	grid  Grid
	start DirPos
	end   Pos
}

func (dir Dir) turn(cwCount int) Dir {
	for range max(cwCount, -cwCount) {
		switch dir {
		case "^":
			if cwCount > 0 {
				dir = ">"
			} else {
				dir = "<"
			}
		case "v":
			if cwCount > 0 {
				dir = "<"
			} else {
				dir = ">"
			}
		case ">":
			if cwCount > 0 {
				dir = "v"
			} else {
				dir = "^"
			}
		case "<":
			if cwCount > 0 {
				dir = "^"
			} else {
				dir = "v"
			}
		}
	}

	return dir
}

func (p Pos) move(dir Dir) Pos {
	switch dir {
	case "^":
		return Pos{p[0], p[1] - 1}
	case "v":
		return Pos{p[0], p[1] + 1}
	case "<":
		return Pos{p[0] - 1, p[1]}
	case ">":
		return Pos{p[0] + 1, p[1]}
	}

	return p
}

func (p DirPos) move() DirPos {
	return DirPos{pos: p.pos.move(p.dir), dir: p.dir}
}

func (p DirPos) turn(cwCount int) DirPos {
	return DirPos{pos: p.pos, dir: p.dir.turn(cwCount)}
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

	for y, line := range lines {
		for x, ch := range line {
			switch ch {
			case 'S':
				input.start = DirPos{pos: Pos{x, y}, dir: ">"}
				input.grid[Pos{x, y}] = "."
			case 'E':
				input.end = Pos{x, y}
				input.grid[Pos{x, y}] = "."
			default:
				input.grid[Pos{x, y}] = string(ch)
			}
		}
	}

	return input, nil
}

func (queue PQueue) append(dp QueueItem) PQueue {
	pos := 0
	for ; pos < len(queue); pos += 1 {
		if queue[pos].score > dp.score {
			break
		}
	}

	nq := append(PQueue{}, queue[0:pos]...)
	nq = append(nq, dp)
	nq = append(nq, queue[pos:]...)

	return nq
}

func draw(grid Grid, visited []Pos) {
	maxX := 0
	maxY := 0
	for p := range grid {
		maxX = max(maxX, p[0])
		maxY = max(maxY, p[1])
	}

	for y := range maxY + 1 {
		for x := range maxX + 1 {
			p := Pos{x, y}
			if grid[p] == "." && slices.Contains(visited, p) {
				fmt.Print("o")
			} else if grid[p] == "#" {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func walk(input Input) (int, int) {
	pqueue := PQueue{}
	distances := map[DirPos]int{}

	ifBetter := func(dp DirPos, score int, vitem []Pos) {
		if previous, found := distances[dp]; found {
			if previous < score {
				return
			}
		}

		distances[dp] = score

		vcopy := make([]Pos, len(vitem))
		copy(vcopy, vitem)

		pqueue = append(pqueue, QueueItem{
			dirpos:  dp,
			score:   score,
			visited: append(vcopy, dp.pos),
		})
	}

	ifBetter(input.start, 0, []Pos{})
	minScore := math.MaxInt

	bestPaths := map[int][]Pos{}

	for len(pqueue) != 0 {
		item := pqueue[0]
		pqueue = pqueue[1:]

		if item.score > minScore {
			continue
		}

		if item.dirpos.pos.eq(input.end) {
			if item.score <= minScore {
				// fmt.Println("GOT PATH", item.visited[0:19])
				minScore = item.score
				bestPaths[minScore] = append(bestPaths[minScore], item.visited...)
			}
			continue
		}
		np := item.dirpos.move()
		if input.grid[np.pos] == "." {
			ifBetter(np, item.score+1, item.visited)
		}
		np = item.dirpos.turn(1)
		ifBetter(np, item.score+1000, item.visited)
		np = item.dirpos.turn(-1)
		ifBetter(np, item.score+1000, item.visited)
	}

	visited := map[Pos]bool{}
	for _, p := range bestPaths[minScore] {
		visited[p] = true
	}

	// draw(input.grid, slices.Collect(maps.Keys(visited)))

	return minScore, len(visited)
}

func PartOneSolution(input Input) (int, error) {
	sum := 0

	sum, _ = walk(input)

	return sum, nil
}

func PartTwoSolution(input Input) (int, error) {
	_, sum := walk(input)

	// too low: 572
	// want: 583
	return sum, nil
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
	values, err = PartTwoSolution(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 2 (%.2fms): %v\n", float64(time.Since(timeStart).Microseconds())/1000, values)
}
