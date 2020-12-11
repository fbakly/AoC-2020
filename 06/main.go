package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func part1(lines []string) int {
	questions := ""
	sum := 0

	for _, line := range lines {
		if len(line) > 0 {
			for _, c := range line {
				if !strings.Contains(questions, string(c)) {
					questions += string(c)
				}
			}
		} else {
			sum += len(questions)
			questions = ""
		}
	}
	return sum
}

func part2(lines []string) int {
	groupQuestions := make(map[rune]int)
	var sum, numPeople int

	for _, line := range lines {
		if len(line) > 0 {
			numPeople += 1
			for _, c := range line {
				_, exists := groupQuestions[c]
				if exists {
					groupQuestions[c] += 1
				} else {
					groupQuestions[c] = 1
				}
			}
		} else {
			for _, val := range groupQuestions {
				if val == numPeople {
					sum += 1
				}
			}
			groupQuestions = make(map[rune]int)
			numPeople = 0
		}
	}
	return sum
}

func main() {
	lines := getInput()
	fmt.Printf("Part 1 solution: %d\n", part1(lines))
	fmt.Printf("Part 2 solution: %d\n", part2(lines))
}
