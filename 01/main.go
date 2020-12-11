package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part1(numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] + numbers[j] == 2020 {
				return numbers[i]*numbers[j]
			}
		}
	}
	return -1
}

func part2(numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			for k := j + 1; k < len(numbers); k++ {
				if numbers[i] + numbers[j] + numbers[k]== 2020 {
					return numbers[i]*numbers[j]*numbers[k]
				}
			}
		}
	}
	return -1
}

func getInput(filePath string) []int {
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	lines := strings.Fields(string(data))
	var nums []int

	for _, val := range lines {
		num, _ := strconv.Atoi(val)
		nums = append(nums, num)
	}
	return nums
}

func main() {
	data := getInput("input")

	fmt.Printf("Part 1 solution: %d\n", part1(data))
	fmt.Printf("Part 2 solution: %d\n", part2(data))
}
