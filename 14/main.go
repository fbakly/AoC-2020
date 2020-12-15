package main

import (
	"aoc"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func part1(lines []string) uint64 {
	var mask string
	mem := make(map[int]uint64)

	for _, line := range lines {
		if strings.Contains(line, "mask") {
			fields := strings.Fields(line)
			mask = strings.TrimSpace(fields[2])
		} else {
			fields := strings.Fields(line)
			val, _ := strconv.Atoi(fields[2])
			startIndex := strings.Index(fields[0], "[")
			stopIndex := strings.Index(fields[0], "]")
			memLoc, _ := strconv.Atoi(fields[0][startIndex+1 : stopIndex])
			var newVal uint64 = 0
			N := len(mask) - 1
			for i := N; i >= 0; i-- {
				if mask[i] == '1' {
					newVal += uint64(math.Pow(2.0, float64(N-i)))
				} else if mask[i] == 'X' {
					currBit := uint64(val) & uint64(math.Pow(2.0, float64(N-i)))
					newVal += currBit
				}
			}
			mem[memLoc] = newVal
		}
	}

	var sum uint64

	for _, value := range mem {
		sum += value
	}
	return sum
}

func findLocs(newIndex uint64, floating []uint64) []uint64 {
	if len(floating) == 0 {
		return []uint64{newIndex}
	}
	res := []uint64{}
	b0 := floating[0]
	rest := floating[1:]
	// Find all the possible locations for this X = 0
	res = append(res, findLocs(newIndex, rest)...)
	// Find all the possible locations for this X = 1
	temp := uint64(math.Pow(2.0, float64(b0)))
	res = append(res, findLocs(newIndex+temp, rest)...)
	return res
}

func changeLocs(memLoc uint64, mask string) []uint64 {
	var newIndex uint64 = 0
	var floating []uint64
	N := len(mask) - 1
	for i := N; i >= 0; i-- {
		if mask[i] == 'X' {
			floating = append(floating, uint64(N-i))
		} else if mask[i] == '1' {
			newIndex += uint64(math.Pow(2.0, float64(N-i)))
		} else {
			newIndex += memLoc & uint64(math.Pow(2.0, float64(N-i)))
		}
	}
	return findLocs(newIndex, floating)
}

func part2(lines []string) uint64 {
	var mask string
	mem := make(map[uint64]uint64)

	for _, line := range lines {
		if strings.Contains(line, "mask") {
			fields := strings.Fields(line)
			mask = strings.TrimSpace(fields[2])
		} else {
			fields := strings.Fields(line)
			val, _ := strconv.Atoi(fields[2])
			startIndex := strings.Index(fields[0], "[")
			stopIndex := strings.Index(fields[0], "]")
			memLoc, _ := strconv.Atoi(fields[0][startIndex+1 : stopIndex])
			I := changeLocs(uint64(memLoc), mask)
			for _, index := range I {
				mem[index] = uint64(val)
			}
		}
	}

	var sum uint64

	for _, value := range mem {
		sum += value
	}
	return sum
}

func main() {
	lines := aoc.GetStdin()
	fmt.Printf("Part 1 solution: %d\n", part1(lines))
	fmt.Printf("Part 2 solution: %d\n", part2(lines))
}
