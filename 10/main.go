package main

import (
	"aoc"
	"fmt"
	"sort"
	"strconv"
)

func part1(nums []int) int {
	diffs := []int{0, 0, 0}

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff == 3 {
			diffs[2]++
		} else if diff == 2 {
			diffs[1]++
		} else if diff == 1 {
			diffs[0]++
		} else {
			fmt.Println("Invalid Joltage difference")
			break
		}
	}

	return diffs[0] * diffs[2]
}

func findArrangements(nums []int, m map[int]uint64, index int) uint64 {
	if index == len(nums)-1 {
		return 1
	}

	if _, ok := m[index]; ok {
		return m[index]
	}

	var numArrangements uint64 = 0

	for j := index + 1; j < len(nums); j++ {
		diff := nums[j] - nums[index]
		if diff >= 1 && diff <= 3 {
			numArrangements += findArrangements(nums, m, j)
		} else {
			break
		}
	}
	m[index] = numArrangements
	return numArrangements
}

func part2(nums []int) uint64 {
	m := make(map[int]uint64)
	return findArrangements(nums, m, 0)
}

func main() {
	lines := aoc.GetStdin()
	nums := []int{}

	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}

	nums = append(nums, 0)
	sort.Ints(nums)
	nums = append(nums, nums[len(nums)-1]+3)

	fmt.Printf("Part 1 solution: %d\n", part1(nums))
	fmt.Printf("Part 2 solution: %d\n", part2(nums))
}
