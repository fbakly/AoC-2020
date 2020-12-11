package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func part1(lines []string) int {
	numValid := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		minMax := strings.Split(fields[0], string('-'))
		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])
		count := strings.Count(fields[2], string(fields[1][0]))
		if count >= min && count <= max {
			numValid += 1
		}
	}
	return numValid
}

func part2(lines []string) int {
	numValid := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		positions := strings.Split(fields[0], string('-'))
		pos1, _ := strconv.Atoi(positions[0])
		pos2, _ := strconv.Atoi(positions[1])
		if fields[2][pos1-1] != fields[2][pos2-1] &&
			(fields[2][pos1-1] == fields[1][0] ||
				fields[2][pos2-1] == fields[1][0]) {
			numValid += 1
		}
	}
	return numValid
}

func main() {
	lines := getInput()
	fmt.Printf("Part 1 solution: %d\n", part1(lines))
	fmt.Printf("Part 2 solution: %d\n", part2(lines))
}
