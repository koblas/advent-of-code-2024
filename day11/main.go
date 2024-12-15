package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Input struct {
	data []int
}

func ParseInput(lines []string) (Input, error) {
	input := Input{}

	for _, line := range lines {
		for _, valus := range strings.Split(line, " ") {
			v, err := strconv.Atoi(valus)
			if err != nil {
				return Input{}, err
			}
			input.data = append(input.data, v)
		}
	}

	return input, nil
}

func do(data []int) []int {
	var output []int

	for _, item := range data {
		if item == 0 {
			output = append(output, 1)
			continue
		}
		sval := strconv.FormatInt(int64(item), 10)
		if len(sval)%2 == 1 {
			output = append(output, item*2024)
			continue
		}
		pow := pows[len(sval)/2]
		output = append(output, item/pow)
		output = append(output, item%pow)
	}

	return output
}

var pows = []int{
	1,
	10,
	100,
	1000,
	10000,
	100000,
	1000000,
	10000000,
	100000000,
	1000000000,
	10000000000,
	100000000000,
	1000000000000,
	10000000000000,
	100000000000000,
	100000000000000,
	1000000000000000,
	10000000000000000,
	100000000000000000,
}

var memo = map[[2]int]int{}

func do2(value int, depth int) int {
	if v, found := memo[[2]int{value, depth}]; found {
		return v
	}

	if depth == 0 {
		return 1
	}

	save := func(v int) int {
		memo[[2]int{value, depth}] = v

		return v
	}

	if value == 0 {
		return save(do2(1, depth-1))
	}

	sval := strconv.FormatInt(int64(value), 10)
	if len(sval)%2 == 1 {
		return save(do2(value*2024, depth-1))
	}
	pow := pows[len(sval)/2]

	return save(do2(value/pow, depth-1) + do2(value%pow, depth-1))
}

func PartOneSolution(input Input) (int, error) {
	stones := input.data
	for range 25 {
		stones = do(stones)
	}
	fmt.Println("GOT           do", len(stones))

	sum := 0
	for _, value := range input.data {
		sum += do2(value, 25)
	}

	// right answer: 190865
	return sum, nil
}

func PartTwoSolution(input Input) (int, error) {
	sum := 0
	for _, value := range input.data {
		sum += do2(value, 75)
	}

	// result: 225404711855335
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
