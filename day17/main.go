package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
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
	reg     [3]int
	program []int
	pval    uint64
}

func ParseInput(lines []string) (Input, error) {
	input := Input{
		reg: [3]int{},
	}

	for _, line := range lines {
		if strings.HasPrefix(line, "Register") {
			p := strings.Split(line, " ")
			v, err := strconv.Atoi(p[2])
			if err != nil {
				return Input{}, err
			}
			switch p[1] {
			case "A:":
				input.reg[0] = v
			case "B:":
				input.reg[1] = v
			case "C:":
				input.reg[2] = v
			}
		} else if strings.HasPrefix(line, "Program: ") {
			line = strings.TrimPrefix(line, "Program: ")

			for _, item := range strings.Split(line, ",") {
				v, err := strconv.Atoi(item)
				if err != nil {
					return Input{}, err
				}
				input.program = append(input.program, v)
			}
		}
	}

	for i, v := range input.program {
		input.pval = input.pval | uint64((v << (3 * i)))
	}

	return input, nil
}

func runFast(ra int, input Input) []int {
	var output []int

	rb := input.reg[1]
	rc := input.reg[2]

	const mask = 077

	maxIdx := len(input.program)
	prog := input.pval
	for pc := 0; pc < maxIdx; {
		pair := (prog >> (3 * pc)) & mask
		literal, opcode := int((pair>>3)&0x7), pair&0x7
		pc += 2

		if opcode == 3 {
			if ra != 0 {
				pc = literal
			}
			continue
		}

		combo := literal
		switch literal {
		case 4:
			combo = ra
		case 5:
			combo = rb
		case 6:
			combo = rc
		}

		switch opcode {
		case 0: //  adv
			ra = ra >> combo
		case 1: //  bxl
			rb = rb ^ literal
		case 2: //  bst
			rb = combo % 8
		case 3: //  jnz
			// handled
		case 4: //  bxc
			rb = rb ^ rc
		case 5: //  out
			output = append(output, combo%8)
		case 6: //  bdv
			rb = ra >> combo
		case 7: //  cdv
			rc = ra >> combo
		}
	}

	return output
}

func PartOneSolution(input Input) (string, error) {
	strs := []string{}
	for _, item := range runFast(input.reg[0], input) {
		strs = append(strs, strconv.FormatInt(int64(item), 10))
	}

	return strings.Join(strs, ","), nil
}

func PartTwoSolution(input Input) (int, error) {
	a := 0 // the initial value of register A doesnt matter here, so we can reset it
	for pos := len(input.program) - 1; pos >= 0; pos-- {
		a <<= 3 // shift left by 3 bits
		for !slices.Equal(runFast(a, input), input.program[pos:]) {
			a++
		}
	}

	return a, nil
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
	valuesT, err := PartTwoSolution(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 2 (%.2fms): %v\n", float64(time.Since(timeStart).Microseconds())/1000, valuesT)
}
