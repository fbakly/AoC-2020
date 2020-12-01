package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func readFile(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, num)
	}
	return lines, scanner.Err()
}

func main() {
	data, err := readFile("input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1 solution: %d\n", part1(data))
	fmt.Printf("Part 2 solution: %d\n", part2(data))
}
