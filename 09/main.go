package main

import (
	"aoc"
	"fmt"
	"math"
	"strconv"
)

func hasTwoSum(num int, preamble []int) bool {
	s := make(map[int]int)
	for _, n := range preamble {
		diff := num - n
		s[n] = diff
		if _, ok := s[diff]; ok {
			return true
		}
	}
	return false
}

func minMax(ar []int) (int, int) {
	min := math.MaxInt64
	max := math.MinInt64

	for _, num := range ar {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return min, max
}

func sumArr(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func part1(nums []int, lenPreamble int) int {

	for i, num := range nums {
		if i >= lenPreamble {
			preamble := nums[i-lenPreamble : i]
			if !hasTwoSum(num, preamble) {
				return num
			}
		}
	}
	return -1
}

func part2(invalid int, nums []int) int {
	lenSubset := 1

	for i := 0; i < len(nums)-lenSubset; i++ {
		for true {
			sumSubset := sumArr(nums[i : i+lenSubset+1])
			if sumSubset > invalid {
				lenSubset = 1
				break
			} else if sumSubset < invalid {
				lenSubset++
			} else {
				min, max := minMax(nums[i : i+lenSubset+1])
				return min + max
			}
		}
	}
	return -1
}

func main() {
	lines := aoc.GetStdin()
	var nums []int

	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}

	invalid := part1(nums, 25)
	fmt.Printf("Part 1 solution: %d\n", invalid)
	fmt.Printf("Part 2 solution: %d\n", part2(invalid, nums))
}
