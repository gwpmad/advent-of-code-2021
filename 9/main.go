package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/gwpmad/advent-of-code-2021/util"
)

func main() {
	lines := util.ParseInputLinesToStringSlice("./input")
	grid := createGrid(lines)
	lowPoints := getLowPoints(grid)

	switch os.Args[1] {
	case "1":
		one(grid, lowPoints)
	case "2":
		two(grid, lowPoints)
	}
}

func one(grid grid, lowPoints []string) {
	result := 0

	for _, coords := range lowPoints {
		entry := grid[coords]
		lowPoint := true
		for _, coord := range entry.neighbours {
			if grid[coord].value <= entry.value {
				lowPoint = false
			}
		}

		if lowPoint {
			result += (entry.value + 1)
		}
	}
	fmt.Println("result:", result)
}

func two(grid grid, lowPoints []string) {
	basinSizes := []int{}

	for _, lowPointCoords := range lowPoints {
		size := getValleySize(grid, lowPointCoords)
		basinSizes = append(basinSizes, size)
	}

	sort.Slice(basinSizes, func(i, j int) bool { return basinSizes[i] > basinSizes[j] })
	result := basinSizes[0] * basinSizes[1] * basinSizes[2]
	fmt.Println("result:", result)
}

func getValleySize(grid grid, lowPointCoords string) int {
	start := grid[lowPointCoords]

	queue := make([]string, len(start.neighbours))
	visitedCoords := map[string]struct{}{lowPointCoords: {}}
	copy(queue, start.neighbours)

	size := 1
	for len(queue) > 0 {
		coords := queue[0]
		queue = queue[1:]
		entry := grid[coords]

		if _, ok := visitedCoords[coords]; ok || entry.value == 9 {
			continue
		}

		size += 1
		queue = prepend(queue, entry.neighbours) // depth first search
		visitedCoords[coords] = struct{}{}
	}
	return size
}

func prepend(slice []string, sliceToPrepend []string) []string {
	prependSliceCopy := make([]string, len(sliceToPrepend))
	copy(prependSliceCopy, sliceToPrepend)
	result := append(prependSliceCopy, slice...)
	return result
}

func getLowPoints(grid grid) []string {
	lowPoints := []string{}
	for coords, entry := range grid {
		lowPoint := true
		for _, coord := range entry.neighbours {
			if grid[coord].value <= entry.value {
				lowPoint = false
			}
		}

		if lowPoint {
			lowPoints = append(lowPoints, coords)
		}
	}
	return lowPoints
}

type grid map[string]struct {
	value      int
	neighbours []string
}

func createGrid(lines []string) grid {
	grid := grid{}

	for y, line := range lines {
		numbers := strings.Split(line, "")
		for x, num := range numbers {
			coords := getCoords(x, y)
			entry := grid[coords]
			entry.value, _ = strconv.Atoi(num)

			if x > 0 {
				entry.neighbours = append(entry.neighbours, getCoords(x-1, y))
			}
			if x < (len(numbers) - 1) {
				entry.neighbours = append(entry.neighbours, getCoords(x+1, y))
			}
			if y > 0 {
				entry.neighbours = append(entry.neighbours, getCoords(x, y-1))
			}
			if y < (len(lines) - 1) {
				entry.neighbours = append(entry.neighbours, getCoords(x, y+1))
			}
			grid[coords] = entry
		}
	}

	return grid
}

func getCoords(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
