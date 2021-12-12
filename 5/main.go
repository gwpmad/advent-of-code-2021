package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/gwpmad/advent-of-code-2021/util"
)

func main() {
	lines := util.ParseInputLinesToStringSlice("./input")
	switch os.Args[1] {
	case "1":
		one(lines)
	case "2":
		two(lines)
	}
}

func one(lines []string) {
	coords := parseLinesToCoords(lines)
	grid := coordsMap{}
	for _, ints := range coords {
		xStart := ints[0]
		yStart := ints[1]
		xEnd := ints[2]
		yEnd := ints[3]

		if (xStart != xEnd) && (yStart != yEnd) {
			continue
		}

		countWhenOneCoordDoesNotChange(grid, xStart, yStart, xEnd, yEnd)
	}
	result := 0
	for _, count := range grid {
		if count > 1 {
			result++
		}
	}
	fmt.Println("result:", result)
}

func two(lines []string) {
	coords := parseLinesToCoords(lines)
	grid := coordsMap{}
	for _, ints := range coords {
		xStart := ints[0]
		yStart := ints[1]
		xEnd := ints[2]
		yEnd := ints[3]

		if (xStart != xEnd) && (yStart != yEnd) {
			countWhenBothCoordsChange(grid, xStart, yStart, xEnd, yEnd)
		} else {
			countWhenOneCoordDoesNotChange(grid, xStart, yStart, xEnd, yEnd)
		}
	}
	result := 0
	for _, count := range grid {
		if count > 1 {
			result++
		}
	}
	fmt.Println("result:", result)
}

func countWhenOneCoordDoesNotChange(grid coordsMap, xStart int, yStart int, xEnd int, yEnd int) {
	if xEnd < xStart {
		tmp := xStart
		xStart = xEnd
		xEnd = tmp
	}
	if yEnd < yStart {
		tmp := yStart
		yStart = yEnd
		yEnd = tmp
	}

	for i := xStart; i <= xEnd; i++ {
		for j := yStart; j <= yEnd; j++ {
			coord := fmt.Sprintf("%d,%d", i, j)
			grid[coord]++
		}
	}
}

func countWhenBothCoordsChange(grid coordsMap, xStart int, yStart int, xEnd int, yEnd int) {
	var xIncrementer int
	var yIncrementer int
	if xStart < xEnd {
		xIncrementer = 1
	} else {
		xIncrementer = -1
	}
	if yStart < yEnd {
		yIncrementer = 1
	} else {
		yIncrementer = -1
	}

	j := yStart
	for i := xStart; i != (xEnd + xIncrementer); i += xIncrementer {
		coord := fmt.Sprintf("%d,%d", i, j)
		grid[coord]++
		j += yIncrementer
	}
}

type coordsMap map[string]int
type coords []int

func parseLinesToCoords(lines []string) (parsedCoords []coords) {
	digitRegexp := regexp.MustCompile(`\d+`)

	for _, line := range lines {
		digits := digitRegexp.FindAllStringSubmatch(line, -1)
		ints := make([]int, len(digits))
		for i, g := range digits {
			ints[i], _ = strconv.Atoi(g[0])
		}
		parsedCoords = append(parsedCoords, ints)
	}
	return parsedCoords
}
