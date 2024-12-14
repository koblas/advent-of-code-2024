package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Pos [2]int

type Problem struct {
	moveA Pos
	moveB Pos
	prize Pos
}

type Input struct {
	data []Problem
}

var buttonRE = regexp.MustCompile(`Button ([AB]): X\+(\d+), Y\+(\d+)`)
var prizeRE = regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

func ParseInput(lines []string) (Input, error) {
	input := Input{}

	prob := Problem{}
	for _, line := range lines {
		if m := buttonRE.FindStringSubmatch(line); m != nil {
			x, _ := strconv.Atoi(m[2])
			y, _ := strconv.Atoi(m[3])
			if m[1] == "A" {
				prob.moveA = Pos{x, y}
			} else {
				prob.moveB = Pos{x, y}
			}
		} else if m := prizeRE.FindStringSubmatch(line); m != nil {
			x, _ := strconv.Atoi(m[1])
			y, _ := strconv.Atoi(m[2])
			prob.prize = Pos{x, y}
			input.data = append(input.data, prob)
			prob = Problem{}
		}
	}

	return input, nil
}

func check(prob Problem) int {
	// A = (8400 - B(22)) / 94
	// B = (94 * 5400 - 34 * 8400) / (-34 * 22 + 94 * 67)

	bn := (prob.moveA[0]*prob.prize[1] - prob.moveA[1]*prob.prize[0])
	bd := (-prob.moveA[1]*prob.moveB[0] + prob.moveA[0]*prob.moveB[1])

	if bn%bd != 0 {
		return 0
	}

	b := bn / bd
	an := prob.prize[0] - b*prob.moveB[0]
	ad := prob.moveA[0]
	if an%ad != 0 {
		return 0
	}

	a := an / ad

	return a*3 + b*1
}

func checkBig(prob Problem) int {
	// A = (8400 - B(22)) / 94
	// B = (94 * 5400 - 34 * 8400) / (-34 * 22 + 94 * 67)
	prizeX := big.NewInt(0).Add(big.NewInt(int64(prob.prize[0])), big.NewInt(10000000000000))
	prizeY := big.NewInt(0).Add(big.NewInt(int64(prob.prize[1])), big.NewInt(10000000000000))

	// bn := (prob.moveA[0]*prob.prize[1] - prob.moveA[1]*prob.prize[0])
	bnA := big.NewInt(0).Mul(big.NewInt(int64(prob.moveA[0])), prizeY)
	bnB := big.NewInt(0).Mul(big.NewInt(int64(prob.moveA[1])), prizeX)
	bn := big.NewInt(0).Sub(bnA, bnB)

	// bd := (-prob.moveA[1]*prob.moveB[0] + prob.moveA[0]*prob.moveB[1])
	bdA := big.NewInt(0).Mul(big.NewInt(int64(prob.moveA[1])), big.NewInt(int64(prob.moveB[0])))
	bdB := big.NewInt(0).Mul(big.NewInt(int64(prob.moveA[0])), big.NewInt(int64(prob.moveB[1])))
	bd := big.NewInt(0).Sub(bdB, bdA)

	if big.NewInt(0).Mod(bn, bd).Cmp(big.NewInt(0)) != 0 {
		return 0
	}

	// b := bn / bd
	b := big.NewInt(0).Div(bn, bd)

	// an := prob.prize[0] - b*prob.moveB[0]
	// ad := prob.moveA[0]
	anB := big.NewInt(0).Mul(b, big.NewInt(int64(prob.moveB[0])))
	an := big.NewInt(0).Sub(prizeX, anB)
	ad := big.NewInt(int64(prob.moveA[0]))
	if big.NewInt(0).Mod(an, ad).Cmp(big.NewInt(0)) != 0 {
		return 0
	}

	// a := an / ad
	a := big.NewInt(0).Div(an, ad)

	ac := big.NewInt(0).Add(big.NewInt(0).Mul(a, big.NewInt(3)), b)

	cost, _ := strconv.Atoi(ac.String())

	return cost
}

func PartOneSolution(input Input) (int, error) {
	sum := 0

	for _, item := range input.data {
		sum += check(item)
	}

	return sum, nil
}

func PartTwoSolution(input Input) (int, error) {
	sum := 0

	for _, item := range input.data {
		sum += checkBig(item)
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
