package main

import (
	"aoc"
	"fmt"
)

type Point3D struct {
	x, y, z int
}

type Point4D struct {
	x, y, z, w int
}

var exists struct{} = struct{}{}

func part1(lines []string) uint64 {
	active := make(map[Point3D]struct{})
	for x, line := range lines {
		for y, state := range line {
			if state == '#' {
				active[Point3D{x, y, 0}] = exists
			}
		}
	}
	for iter := 0; iter < 6; iter++ {
		newActive := make(map[Point3D]struct{})
		for x := -15; x <= 15; x++ {
			for y := -15; y <= 15; y++ {
				for z := -15; z <= 15; z++ {
					numActive := 0
					for dx := -1; dx <= 1; dx++ {
						for dy := -1; dy <= 1; dy++ {
							for dz := -1; dz <= 1; dz++ {
								if dx == 0 && dy == 0 && dz == 0 {
									continue
								}
								if _, ok := active[Point3D{x + dx, y + dy, z + dz}]; ok {
									numActive += 1
								}
							}
						}
					}
					if numActive == 3 {
						newActive[Point3D{x, y, z}] = exists
					}
					if _, ok := active[Point3D{x, y, z}]; ok && numActive == 2 {
						newActive[Point3D{x, y, z}] = exists
					}

				}

			}
		}
		active = newActive
	}
	return uint64(len(active))
}

func part2(lines []string) uint64 {
	active := make(map[Point4D]struct{})
	for x, line := range lines {
		for y, state := range line {
			if state == '#' {
				active[Point4D{x, y, 0, 0}] = exists
			}
		}
	}
	for iter := 0; iter < 6; iter++ {
		newActive := make(map[Point4D]struct{})
		for x := -15; x <= 15; x++ {
			for y := -15; y <= 15; y++ {
				for z := -15; z <= 15; z++ {
					for w := -15; w <= 15; w++ {
						numActive := 0
						for dx := -1; dx <= 1; dx++ {
							for dy := -1; dy <= 1; dy++ {
								for dz := -1; dz <= 1; dz++ {
									for dw := -1; dw <= 1; dw++ {
										if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
											continue
										}
										if _, ok := active[Point4D{x + dx, y + dy, z + dz, w + dw}]; ok {
											numActive += 1
										}
									}
								}
							}
						}
						if numActive == 3 {
							newActive[Point4D{x, y, z, w}] = exists
						}
						if _, ok := active[Point4D{x, y, z, w}]; ok && numActive == 2 {
							newActive[Point4D{x, y, z, w}] = exists
						}
					}
				}

			}
		}
		active = newActive
	}
	return uint64(len(active))
}

func main() {
	lines := aoc.GetStdin()
	fmt.Printf("Part 1 solution: %d\n", part1(lines))
	fmt.Printf("Part 2 solution: %d\n", part2(lines))
}
