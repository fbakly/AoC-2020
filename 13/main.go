package main

import (
	"aoc"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func part1(earliestTime int, busTimes []int) int {
	minTime := math.MaxInt64
	minID := math.MaxInt64

	for _, busTime := range busTimes {
		timeDiff := busTime - (earliestTime % busTime)
		if timeDiff < minTime {
			minTime = timeDiff
			minID = busTime
		}
	}
	return minID * minTime
}

func part2(busTimes, timeDiffs []int) int64 {
	multiplier := 1
	var earliestTime int64 = math.MaxInt64
	var times []int64

OUTER:
	for true {
		times = times[:0]
		for _, val := range busTimes {
			times = append(times, int64(multiplier*val))
		}
		for i := 1; i < len(times); i++ {
			if int(times[i]-times[i-1]) != timeDiffs[i] {
				multiplier += 1
				continue OUTER
			}
		}
		for _, val := range times {
			if val < earliestTime {
				earliestTime = val
			}
		}
		break

	}
	return earliestTime
}

func main() {
	lines := aoc.GetStdin()
	earliestTime, _ := strconv.Atoi(lines[0])
	temp := strings.Split(lines[1], ",")
	var busTimes []int
	var timeDiffs []int

	for i, val := range temp {
		time, err := strconv.Atoi(val)
		if err == nil {
			busTimes = append(busTimes, time)
			timeDiffs = append(timeDiffs, i)
		}
	}
	fmt.Println(busTimes)
	fmt.Printf("Part 1 solution: %d\n", part1(earliestTime, busTimes))
	fmt.Printf("Part 2 solution: %d\n", part2(busTimes, timeDiffs))
}
