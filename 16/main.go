package main

import (
	"aoc"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Boundry struct {
	firstMin, firstMax, secondMin, secondMax int
}

func checkValid(num int, bounds Boundry) bool {
	if !((num >= bounds.firstMin && num <= bounds.firstMax) ||
		(num >= bounds.secondMin && num <= bounds.secondMax)) {
		return false
	}
	return true
}

func getBoundaries(lines []string, m map[string]Boundry) Boundry {
	genBounds := Boundry{math.MaxInt32, math.MinInt32,
		math.MaxInt32, math.MinInt32}
	for i := 0; len(lines[i]) > 0; i++ {
		j := strings.Index(lines[i], ":")
		line := lines[i][j+1:]
		line = strings.Replace(line, "or", "", -1)
		fields := strings.Fields(line)
		tempFirst := strings.Split(fields[0], "-")
		tempFirstMin, _ := strconv.Atoi(tempFirst[0])
		tempFirstMax, _ := strconv.Atoi(tempFirst[1])

		if tempFirstMax > genBounds.firstMax {
			genBounds.firstMax = tempFirstMax
		}

		if tempFirstMin < genBounds.firstMin {
			genBounds.firstMin = tempFirstMin
		}

		tempSecond := strings.Split(fields[1], "-")
		tempSecondMin, _ := strconv.Atoi(tempSecond[0])
		tempSecondMax, _ := strconv.Atoi(tempSecond[1])

		if tempSecondMax >= genBounds.secondMax {
			if tempSecondMax >= genBounds.firstMax {
				genBounds.secondMax = tempSecondMax
			} else {
				genBounds.secondMax = genBounds.firstMax
			}
		}

		if tempSecondMin <= genBounds.secondMin {
			if tempSecondMin >= genBounds.firstMax {
				genBounds.secondMin = tempSecondMin
			} else {
				genBounds.secondMin = genBounds.firstMax
			}
		}
		m[lines[i][:j]] = Boundry{tempFirstMin, tempFirstMax,
			tempSecondMin, tempSecondMax}
	}
	return genBounds
}

func part1(lines []string, i int, bounds Boundry) (int, map[int]bool) {
	sumInvalid := 0
	invalidLines := make(map[int]bool)

	for ; i < len(lines); i++ {
		line := lines[i]
		fields := strings.Split(line, ",")
		for _, f := range fields {
			num, _ := strconv.Atoi(f)
			if !checkValid(num, bounds) {
				sumInvalid += num
				invalidLines[i] = true
			}
		}

	}
	return sumInvalid, invalidLines
}

func findMissing(all map[string]Boundry, invalids []string) string {
	for key, _ := range all {
		present := false
		for _, s := range invalids {
			if s == key {
				present = true
				break
			}
		}
		if !present {
			return key
		}
	}
	return ""
}

func solve(n int, all map[string]Boundry, solved map[int]string, invalids map[int][]string) {
	for pos, invalid := range invalids {
		if len(invalid) == n-1 {
			missing := findMissing(all, invalid)
			solved[pos] = missing
			delete(invalids, pos)
			for p, _ := range invalids {
				invalids[p] = append(invalids[p], missing)
			}
			solve(n, all, solved, invalids)
		}
	}
}

func multDeparture(lines []string, solved map[int]string) uint64 {
	var res uint64 = 1

	for i := 0; i < len(lines); i++ {
		if strings.Index(lines[i], "your") != -1 {
			i++
			fields := strings.Split(lines[i], ",")
			for i, f := range fields {
				num, _ := strconv.Atoi(f)
				if strings.Contains(solved[i], "departure") {
					res *= uint64(num)
				}
			}
		}
	}
	return res
}

func part2(lines []string, i int, invalidLines map[int]bool, all map[string]Boundry) uint64 {
	invalids := make(map[int][]string)
	N := 0
	for ; i < len(lines); i++ {
		if _, ok := invalidLines[i]; ok {
			continue
		}
		fields := strings.Split(lines[i], ",")
		N = len(fields)
		for pos, f := range fields {
			num, _ := strconv.Atoi(f)
			for key, bound := range all {
				if !checkValid(num, bound) {
					if _, ok := invalids[pos]; ok {
						invalids[pos] = append(invalids[pos], key)
					} else {
						invalids[pos] = []string{key}
					}
				}
			}
		}
	}
	solved := make(map[int]string)
	solve(N, all, solved, invalids)
	return multDeparture(lines, solved)
}

func main() {
	lines := aoc.GetStdin()
	genBounds := make(map[string]Boundry)
	bounds := getBoundaries(lines, genBounds)

	i := 0
	for index, line := range lines {
		if strings.Index(line, "nearby") != -1 {
			i = index + 1
			break
		}
	}

	p1, invalidLines := part1(lines, i, bounds)
	fmt.Printf("Part 1 solution: %d\n", p1)
	fmt.Printf("Part 2 solution: %d\n", part2(lines, i, invalidLines, genBounds))
}
