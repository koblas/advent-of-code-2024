package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"
)

type Run struct {
	id     int
	length int
}

type Input struct {
	disk []Run
}

func ParseInput(lines []string) (Input, error) {
	input := Input{}

	id := 0
	for _, line := range lines {
		for idx, ch := range line {
			runlen, err := strconv.Atoi(string(ch))
			if err != nil {
				return Input{}, err
			}

			curId := id
			if idx%2 == 0 {
				id++
			} else {
				curId = -1
			}
			input.disk = append(input.disk, Run{
				id:     curId,
				length: runlen,
			})
		}
	}

	return input, nil
}

type Free struct {
	id     int
	start  int
	length int
}

func (f *Free) move(start int) {
	f.start = start
}

type Freelist struct {
	data []Free
}

func (flist *Freelist) release(start, length int) {
	if length == 0 {
		return
	}

	for idx, item := range flist.data {
		if start < item.start {
			if start+length == item.start {
				// grow this item
				item.start = start
				return
			}
			// insert before
			data := slices.Clone(flist.data[0:idx])
			data = append(data, Free{id: -1, start: start, length: length})
			flist.data = append(data, flist.data[idx:]...)

			return
		}
	}

	flist.data = append(flist.data, Free{id: -1, start: start, length: length})
}

func (flist *Freelist) alloc(length int, maxstart int) (Free, bool) {
	for idx, item := range flist.data {
		// fmt.Println("  CHECK", item.start, item.length, " WANT ", length)
		if item.start >= maxstart {
			return Free{}, false
		}
		if length <= item.length {
			item := flist.data[idx]
			flist.data = append(flist.data[:idx], flist.data[idx+1:]...)

			return item, true
		}
	}

	return Free{}, false
}

func packOne(input Input) []int {
	size := 0
	for _, item := range input.disk {
		size += item.length
	}

	data := make([]int, size)
	offset := 0
	for _, item := range input.disk {
		for range item.length {
			data[offset] = item.id
			offset += 1
		}
	}

	left := 0
	right := size - 1
	for {
		for ; data[left] != -1; left += 1 {
			if left >= right {
				return data
			}
		}
		for ; data[right] == -1; right -= 1 {
			if left >= right {
				return data
			}
		}
		data[left], data[right] = data[right], data[left]
	}
}

func packTwo(input Input) []int {
	freelist := Freelist{}

	var used []Free
	offset := 0
	for _, item := range input.disk {
		if item.id == -1 {
			freelist.data = append(freelist.data, Free{id: -1, start: offset, length: item.length})
		} else {
			used = append(used, Free{id: item.id, start: offset, length: item.length})
		}
		offset += item.length
	}

	for didWork := true; didWork; {
		didWork = false
		for idx := len(used) - 1; idx > 0; idx -= 1 {
			item := used[idx]
			// fmt.Println("LOOKNG FOR ", item.id, item.length, item.start)
			entry, ok := freelist.alloc(item.length, item.start)
			if ok {
				// fmt.Println("  MOVE ", item.id, " FROM", item.start, "TO", entry.start)
				item.start = entry.start
				remain := entry.length - item.length
				freelist.release(entry.start+item.length, remain)
				used = append(used[:idx], used[idx+1:]...)
				pos := 0
				for ; pos < len(used) && used[pos].start < item.start; pos++ {
					// loop
				}
				tmp := append(slices.Clone(used[0:pos]), item)
				used = append(tmp, used[pos:]...)

				didWork = true

				// fmt.Println("  STATE", used)

				break
			}
		}
	}

	result := make([]int, offset)
	for idx := range offset {
		result[idx] = -1
	}
	for _, item := range used {
		for idx := range item.length {
			result[item.start+idx] = item.id
		}
	}

	// fmt.Println(result)
	return result
}

func checksum(data []int) int {
	sum := 0
	for idx, item := range data {
		if item != -1 {
			sum += idx * item
		}
	}

	return sum
}

func PartOneSolution(input Input) (int, error) {
	result := packOne(input)

	return checksum(result), nil
}

func PartTwoSolution(input Input) (int, error) {
	result := packTwo(input)

	return checksum(result), nil
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
