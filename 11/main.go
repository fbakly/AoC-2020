package main

import (
	"aoc"
	"fmt"
)

func recurAdjacent(lines []string, row, col, dRow, dCol int, recur bool) int {
	i := row + dRow
	j := col + dCol

	if (i >= 0 && i < len(lines)) && (j >= 0 && j < len(lines[i])) {
		if lines[i][j] == '#' {
			return 1
		} else if lines[i][j] == 'L' {
			return 0
		} else if recur {
			return recurAdjacent(lines, i, j, dRow, dCol, recur)
		}
	}
	return 0
}

func numAdjacent(lines []string, row, col int, recur bool) int {
	seen := 0

	dRow := []int{-1, 0, 1}
	dCol := []int{-1, 0, 1}

	for _, dR := range dRow {
		for _, dC := range dCol {
			if dR == 0 && dC == 0 {
				continue
			}
			seen += recurAdjacent(lines, row, col, dR, dC, recur)
		}
	}
	return seen
}

func part1(seatMap []string) int {
	numSeats := 0
	lines := make([]string, len(seatMap))
	copy(lines, seatMap)

	for true {
		prevNumSeats := numSeats
		linesCopy := make([]string, len(lines))
		copy(linesCopy, lines)
		for i := 0; i < len(linesCopy); i++ {
			for j := 0; j < len(linesCopy[i]); j++ {
				if linesCopy[i][j] == 'L' {
					if numAdjacent(linesCopy, i, j, false) == 0 {
						lines[i] = lines[i][:j] + "#" + lines[i][j+1:]
						numSeats++
					}
				} else if linesCopy[i][j] == '#' {
					if numAdjacent(linesCopy, i, j, false) >= 4 {
						lines[i] = lines[i][:j] + "L" + lines[i][j+1:]
						numSeats--
					}
				}
			}
		}
		if prevNumSeats == numSeats {
			break
		}
	}

	return numSeats
}

func part2(seatMap []string) int {
	numSeats := 0
	lines := make([]string, len(seatMap))
	copy(lines, seatMap)

	for true {
		prevNumSeats := numSeats
		linesCopy := make([]string, len(lines))
		copy(linesCopy, lines)
		for i := 0; i < len(linesCopy); i++ {
			for j := 0; j < len(linesCopy[i]); j++ {
				if linesCopy[i][j] == 'L' {
					if numAdjacent(linesCopy, i, j, true) == 0 {
						lines[i] = lines[i][:j] + "#" + lines[i][j+1:]
						numSeats++
					}
				} else if linesCopy[i][j] == '#' {
					if numAdjacent(linesCopy, i, j, true) >= 5 {
						lines[i] = lines[i][:j] + "L" + lines[i][j+1:]
						numSeats--
					}
				}
			}
		}
		if prevNumSeats == numSeats {
			break
		}
	}

	return numSeats

}

func main() {
	lines := aoc.GetStdin()
	fmt.Printf("Part 1 solution: %d\n", part1(lines))
	fmt.Printf("Part 2 solution: %d\n", part2(lines))
}
