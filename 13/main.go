package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gwpmad/advent-of-code-2021/util"
)

func main() {
	splitInput := util.ParseFileAndSplitByDelimiter("./input", "\n\n")
	points := strings.Split(splitInput[0], "\n")
	grid, maxX, maxY := createGrid(points)
	folds := parseFolds(strings.Split(splitInput[1], "\n"))

	switch os.Args[1] {
	case "1":
		one(grid, maxX, maxY, folds)
	case "2":
		two(grid, maxX, maxY, folds)
	}
}

func one(grid grid, maxX int, maxY int, folds []foldInstructions) {
	fold := folds[0]
	switch fold.axis {
	case "x":
		foldByX(grid, maxX, maxY, fold.line)
	case "y":
		foldByY(grid, maxX, maxY, fold.line)
	}
	fmt.Println("result:", len(grid))
}

func two(grid grid, maxX int, maxY int, folds []foldInstructions) {
	for _, fold := range folds {
		switch fold.axis {
		case "x":
			foldByX(grid, maxX, maxY, fold.line)
			maxX = fold.line - 1
		case "y":
			foldByY(grid, maxX, maxY, fold.line)
			maxY = fold.line - 1
		}
	}
	printCodeString(grid, maxX, maxY)
}

func printCodeString(grid grid, maxX int, maxY int) {
	str := ""

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if _, ok := grid[getCoords(x, y)]; ok {
				str = fmt.Sprintf("%v#", str)
			} else {
				str = fmt.Sprintf("%v.", str)
			}
		}
		str = fmt.Sprintf("%v\n", str)
	}
	fmt.Print(str)
}

func foldByX(grid grid, maxX int, maxY int, foldLine int) {
	for x := foldLine + 1; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			if _, ok := grid[getCoords(x, y)]; ok {
				delete(grid, getCoords(x, y))
				newCoords := getCoords(foldLine-(x-foldLine), y)
				grid[newCoords] = struct{}{}
			}
		}
	}
}

func foldByY(grid grid, maxX int, maxY int, foldLine int) {
	for x := 0; x <= maxX; x++ {
		for y := foldLine + 1; y <= maxY; y++ {
			if _, ok := grid[getCoords(x, y)]; ok {
				delete(grid, getCoords(x, y))
				newCoords := getCoords(x, foldLine-(y-foldLine))
				grid[newCoords] = struct{}{}
			}
		}
	}
}

type grid map[string]struct{}
type foldInstructions struct {
	axis string
	line int
}

func createGrid(points []string) (grid, int, int) {
	grid := make(grid)
	maxX := 0
	maxY := 0
	for _, point := range points {
		splitPoint := strings.Split(point, ",")
		if x, _ := strconv.Atoi(splitPoint[0]); x > maxX {
			maxX = x
		}
		if y, _ := strconv.Atoi(splitPoint[1]); y > maxY {
			maxY = y
		}
		grid[point] = struct{}{}
	}

	return grid, maxX, maxY
}

func parseFolds(folds []string) []foldInstructions {
	result := make([]foldInstructions, 0)
	foldRegex := regexp.MustCompile(`[xy]|\d+`)
	for _, fold := range folds {
		parsedFold := foldRegex.FindAllStringSubmatch(fold, -1)
		line, _ := strconv.Atoi(parsedFold[1][0])
		result = append(result, foldInstructions{axis: parsedFold[0][0], line: line})
	}
	return result
}

func getCoords(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
