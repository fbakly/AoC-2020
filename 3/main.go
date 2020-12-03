package main

import (
	"aoc"
	"fmt"
)

func part1(lines []string) int {
	col := 0
	numTrees := 0
	for _, val := range lines {
		if val[col] == '#' {
			numTrees += 1
		}
		col += 3
		col %= len(val)
	}
	return numTrees
}

func part2(lines []string) int {
	res := 1
	dCol := []int{1, 3, 5, 7, 1}
	dRow := []int{1, 1, 1, 1, 2}

	for i := 0; i < len(dRow); i++ {
		numTrees := 0
		var row, col int = 0, 0
		for row < len(lines) {
			if lines[row][col] == '#' {
				numTrees += 1
			}
			col += dCol[i]
			col %= len(lines[row])
			row += dRow[i]
		}
		res *= numTrees
	}
	return res
}

func main() {
	lines := aoc.GetStdin()
	fmt.Printf("Part 1 solution: %d\n", part1(lines))
	fmt.Printf("Part 2 solution: %d\n", part2(lines))
}
