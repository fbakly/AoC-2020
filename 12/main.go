package main

import (
	"aoc"
	"fmt"
	"math"
	"strconv"
)

func move(instruction byte, value int, facing *int, pos []int) {

	switch instruction {
	case 'N':
		pos[1] += value
	case 'S':
		pos[1] -= value
	case 'E':
		pos[0] += value
	case 'W':
		pos[0] -= value
	case 'L':
		*facing += value
		*facing %= 360
	case 'R':
		*facing -= value
		if *facing < 0 {
			*facing += 360
		}
	case 'F':
		switch *facing {
		case 0:
			move('E', value, facing, pos)
		case 90:
			move('N', value, facing, pos)
		case 180:
			move('W', value, facing, pos)
		case 270:
			move('S', value, facing, pos)
		default:
			fmt.Println(*facing)
		}
	default:
		return
	}

}

func moveWaypoint(instruction byte, value int, waypointPos, shipPos []int) {
	switch instruction {
	case 'N':
		waypointPos[1] += value
	case 'S':
		waypointPos[1] -= value
	case 'E':
		waypointPos[0] += value
	case 'W':
		waypointPos[0] -= value
	case 'F':
		shipPos[0] += (waypointPos[0] * value)
		shipPos[1] += (waypointPos[1] * value)
	case 'L':
		rotations := value / 90
		for i := 0; i < rotations; i++ {
			waypointPos[0], waypointPos[1] = -waypointPos[1], waypointPos[0]
		}
	case 'R':
		rotations := value / 90
		for i := 0; i < rotations; i++ {
			waypointPos[0], waypointPos[1] = waypointPos[1], -waypointPos[0]
		}
	default:
		fmt.Println(instruction)
	}
}

func part1(lines []string) int {
	facing := 0
	pos := []int{0, 0}

	for _, line := range lines {
		instruction := line[0]
		value, _ := strconv.Atoi(line[1:])
		move(instruction, value, &facing, pos)
	}
	return int(math.Abs(float64(pos[0])) + math.Abs(float64(pos[1])))
}

func part2(lines []string) int {
	shipPos := []int{0, 0}
	waypointPos := []int{10, 1}

	for _, line := range lines {
		instruction := line[0]
		value, _ := strconv.Atoi(line[1:])
		moveWaypoint(instruction, value, waypointPos, shipPos)
	}
	return int(math.Abs(float64(shipPos[0])) + math.Abs(float64(shipPos[1])))
}

func main() {
	lines := aoc.GetStdin()
	fmt.Printf("Part1 solution: %d\n", part1(lines))
	fmt.Printf("Part2 solution: %d\n", part2(lines))
}
