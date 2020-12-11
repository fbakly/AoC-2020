package main

import (
	"aoc"
	"fmt"
	"sort"
)

func partition(pass string, lo, hi int, comp1, comp2 rune) int {
	mid := lo + ((hi - lo)/2)

	for _, c := range pass {
		if c == comp1 {
			hi = mid
			mid = lo + ((hi - lo)/2)
		} else if c == comp2 {
			lo = mid
			mid = lo + ((hi - lo)/2)
		}
	}
	return mid
}

func part1(lines []string) (int, []int) {
	maxVal := 0
	var seatIDs []int

	for _, line := range lines {
		val := partition(line[:len(line) - 3], 0, 128, 'F', 'B') * 8
		val += partition(line[len(line) - 3: ], 0, 8, 'L', 'R')
		seatIDs = append(seatIDs, val)
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal, seatIDs
}

func part2(seatIDs []int) int {
	N := len(seatIDs)

	sort.Ints(seatIDs)

	for i := 0; i < N - 1; i++ {
		if seatIDs[i + 1] - seatIDs[i] > 1 {
			return seatIDs[i] + 1
		}
	}

	return -1
}

func main() {
	lines := aoc.GetStdin()
	maxSeat, seatIDs := part1(lines)
	seat := part2(seatIDs)
	fmt.Printf("Part 1 solution: %d\n", maxSeat)
	fmt.Printf("Part 2 solution: %d\n", seat)
}
