package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(lines []string) ([]string, int) {
	numValid := 0
	var valid []string

	for _, line := range lines {
		data := strings.Fields(line)
		if len(data) == 8 {
			numValid += 1
			valid = append(valid, line)
		} else if len(data) == 7 && !strings.Contains(line, "cid") {
			numValid += 1
			valid = append(valid, line)
		}
	}
	return valid, numValid
}

func part2(lines []string) int {
	numValid := 0

	for _, line := range lines {
		entries := strings.Fields(line)
		isValid := true
		for _, entry := range entries {
			data := strings.Split(entry, ":")
			switch data[0] {
			case "byr":
				if byr, _ := strconv.Atoi(data[1]); byr < 1920 || byr > 2002 {
					isValid = false
				}
			case "iyr":
				if iyr, _ := strconv.Atoi(data[1]); iyr < 2010 || iyr > 2020 {
					isValid = false
				}
			case "eyr":
				if eyr, _ := strconv.Atoi(data[1]); eyr < 2020 || eyr > 2030 {
					isValid = false
				}
			case "hgt":
				unit := data[1][len(data[1])-2:]
				if unit == "cm" {
					if hgt, _ := strconv.Atoi(data[1][:len(data[1])-2]); hgt < 150 || hgt > 193 {
						isValid = false
					}
				} else if unit == "in" {
					if hgt, _ := strconv.Atoi(data[1][:len(data[1])-2]); hgt < 59 || hgt > 76 {
						isValid = false
					}
				} else {
					isValid = false
				}
			case "hcl":
				if data[1][0] == '#' {
					_, err := hex.DecodeString(data[1][1:])
					if err != nil {
						isValid = false
					}
				} else {
					isValid = false
				}
			case "ecl":
				colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
				inColors := false
				for _, color := range colors {
					if data[1] == color {
						inColors = true
					}
				}
				if !inColors {
					isValid = false
				}
			case "pid":
				if len(data[1]) != 9 {
					isValid = false
				} else if _, err := strconv.Atoi(data[1]); err != nil {
					isValid = false
				}
			default:
				continue
			}
		}
		if isValid {
			numValid += 1
		}
	}

	return numValid
}

func getInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	var data string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			data += " "
			data += line
		} else {
			lines = append(lines, data)
			data = ""
		}
	}
	return lines
}

func main() {
	lines := getInput()
	valid, numValid := part1(lines)
	fmt.Printf("Part 1 solution: %d\n", numValid)
	fmt.Printf("Part 2 solution: %d\n", part2(valid))
}
