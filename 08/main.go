package main

import (
	"aoc"
	"fmt"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	accum := 0
	visitedOps := make(map[int]bool)

	for i := 0; i < len(lines); i++ {
		if _, ok := visitedOps[i]; !ok {
			visitedOps[i] = true
			fields := strings.Fields(lines[i])
			command := fields[0]
			value, _ := strconv.Atoi(fields[1])
			if command == "nop" {
				continue
			} else if command == "acc" {
				accum += value
			} else if command == "jmp" {
				i += value - 1
			}
		} else {
			break
		}
	}
	return accum
}

func part2(lines []string) int {
	accum := 0
	visitedOps := make(map[int]bool)
	visitedIndex := []int{}

	for i := 0; i < len(lines); i++ {
		if _, ok := visitedOps[i]; !ok {
			visitedOps[i] = true
			visitedIndex = append(visitedIndex, i)
			fields := strings.Fields(lines[i])
			command := fields[0]
			value, _ := strconv.Atoi(fields[1])
			if command == "nop" {
				continue
			} else if command == "acc" {
				accum += value
			} else if command == "jmp" {
				i += value - 1
			}
		} else {
			for i := 1; i <= len(visitedIndex); i++ {
				index := visitedIndex[len(visitedIndex)-i]
				fields := strings.Fields(lines[index])
				command := fields[0]
				value, _ := strconv.Atoi(fields[1])
				if command == "jmp" && value < 0 {
					lines[index] = strings.Replace(lines[index], "jmp", "nop", -1)
					break
				} else if command == "nop" {
					lines[index] = strings.Replace(lines[index], "nop", "jmp", -1)
					break
				}
			}
			// Fixed the corrupt command, restart the loop for the correct accum value
			visitedOps = make(map[int]bool)
			accum = 0
			i = -1
		}
	}
	return accum
}

func main() {
	lines := aoc.GetStdin()
	fmt.Printf("Part 1 solution: %d\n", part1(lines))
	fmt.Printf("Part 2 solution: %d\n", part2(lines))
}
