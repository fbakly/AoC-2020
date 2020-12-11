package main

import (
	"aoc"
	"fmt"
	"strconv"
	"strings"
)

func findOuterBags(lines []string, bag string, validBags map[string]bool) int {
	numBags := 0

	for _, line := range lines {
		if strings.Index(line, bag) > 0 {
			fields := strings.Fields(line)
			tempBag := fields[0] + " " + fields[1]
			if _, ok := validBags[tempBag]; !ok {
				numBags += 1
				validBags[tempBag] = true
			}
		}
	}

	if numBags == 0 {
		return numBags
	}

	for key, _ := range validBags {
		numBags += findOuterBags(lines, key, validBags)
	}
	return numBags
}

func findInnerBags(lines []string, color string, totalBagContent map[string]int) int {
	numBags := 0
	bagsInColor := make(map[string]int)

	for _, line := range lines {
		if strings.Index(line, color) == 0 {
			field := strings.Fields(line)
			for i := 4; i < len(field); i += 4 {
				col := field[i+1] + " " + field[i+2]
				bagsInColor[col], _ = strconv.Atoi(field[i])
			}
			break
		}
	}

	for key, val := range bagsInColor {
		if _, ok := totalBagContent[key]; !ok {
			totalInBag := findInnerBags(lines, key, totalBagContent)
			totalBagContent[key] = totalInBag
			numBags += (val + (val * totalInBag))
		} else {
			numBags += (val + (val * totalBagContent[key]))
		}
	}
	return numBags
}

func part1(lines []string) int {
	validBags := make(map[string]bool)
	return findOuterBags(lines, "shiny gold", validBags)
}

func part2(lines []string) int {
	totalBagContent := make(map[string]int) // Added for memoization
	return findInnerBags(lines, "shiny gold", totalBagContent)
}

func main() {
	lines := aoc.GetStdin()
	fmt.Printf("Part 1 solution: %d\n", part1(lines))
	fmt.Printf("Part 2 solution: %d\n", part2(lines))
}
