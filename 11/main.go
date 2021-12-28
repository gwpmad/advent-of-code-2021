package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gwpmad/advent-of-code-2021/util"
)

func main() {
	lines := util.ParseInputLinesToStringSlice("./input")
	grid := createGrid(lines)

	switch os.Args[1] {
	case "1":
		one(grid)
	case "2":
		two(grid)
	}
}

type grid map[string]struct {
	value      int
	neighbours []string
}
type queueItem struct {
	coords      string
	instruction string
}

func one(grid grid) {
	totalFlashes := 0
	for i := 1; i < 101; i++ {
		totalFlashes += doStep(grid)
	}
	fmt.Println("result:", totalFlashes)
}

func two(grid grid) {
	var synchronisedStep int
	for i := 1; synchronisedStep == 0; i++ {
		if doStep(grid) == len(grid) {
			synchronisedStep = i
		}
	}
	fmt.Println("result:", synchronisedStep)
}

func doStep(grid grid) int {
	queue := make([]queueItem, len(grid))
	flashedCoords := map[string]struct{}{}

	i := 0
	for key := range grid {
		queue[i] = queueItem{coords: key, instruction: "increment"}
		i++
	}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if _, ok := flashedCoords[item.coords]; ok {
			continue
		}

		entry := grid[item.coords]
		if item.instruction == "increment" {
			entry.value++
			if entry.value == 10 {
				queue = append(queue, queueItem{coords: item.coords, instruction: "flash"})
			}
		}
		if item.instruction == "flash" {
			entry.value = 0
			for _, neighbour := range entry.neighbours {
				queue = append(queue, queueItem{coords: neighbour, instruction: "increment"})
			}
			flashedCoords[item.coords] = struct{}{}
		}
		grid[item.coords] = entry
	}

	return len(flashedCoords)
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
				if y > 0 {
					entry.neighbours = append(entry.neighbours, getCoords(x-1, y-1))
				}
				if y < (len(lines) - 1) {
					entry.neighbours = append(entry.neighbours, getCoords(x-1, y+1))
				}
			}
			if x < (len(numbers) - 1) {
				entry.neighbours = append(entry.neighbours, getCoords(x+1, y))
				if y > 0 {
					entry.neighbours = append(entry.neighbours, getCoords(x+1, y-1))
				}
				if y < (len(lines) - 1) {
					entry.neighbours = append(entry.neighbours, getCoords(x+1, y+1))
				}
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
