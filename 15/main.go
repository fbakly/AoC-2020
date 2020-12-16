package main

import (
	"aoc"
	"fmt"
	"strconv"
	"strings"
)

func solve(nums []int, target int) int {
	m := make(map[int][]int)
	lastNum := 0
	for i := 0; i < target; i++ {
		if i < len(nums) {
			m[nums[i]] = []int{i}
			lastNum = nums[i]
		} else {
			if len(m[lastNum]) > 1 {
				lastNum = m[lastNum][len(m[lastNum])-1] - m[lastNum][len(m[lastNum])-2]
				if _, ok := m[lastNum]; ok {
					m[lastNum] = append(m[lastNum], i)
				} else {
					m[lastNum] = []int{i}
				}
			} else {
				lastNum = 0
				m[lastNum] = append(m[lastNum], i)
			}
		}
	}
	return lastNum
}

func main() {
	lines := strings.Split(aoc.GetStdin()[0], ",")
	var nums []int
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic("Error converting to int")
		}
		nums = append(nums, num)
	}
	fmt.Printf("Part 1 solution: %d\n", solve(nums, 2020))
	fmt.Printf("Part 2 solution: %d\n", solve(nums, 30000000))
}
