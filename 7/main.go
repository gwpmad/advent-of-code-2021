package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/gwpmad/advent-of-code-2021/util"
)

func main() {
	ints := util.ParseSingleLineToIntSlice("./input")
	sort.Ints(ints)
	switch os.Args[1] {
	case "1":
		one(ints)
	case "2":
		two(ints)
	}
}

func one(ints []int) {
	var median int
	length := len(ints)
	halfLength := length / 2

	if length%2 == 0 {
		median = (ints[halfLength-1] + ints[halfLength]) / 2
	} else {
		median = ints[halfLength]
	}

	fuel := 0
	for i := 0; i < len(ints); i++ {
		fuel += abs(ints[i] - median)
	}
	fmt.Println("result:", fuel)
}

func two(ints []int) {
	floatLength := float64(len(ints))
	floatSum := float64(0)

	for _, n := range ints {
		floatSum += float64(n)
	}

	mean := int(floatSum / floatLength)

	fuel := 0
	for i := 0; i < len(ints); i++ {
		difference := abs(ints[i] - mean)
		triangularNumber := (difference * (difference + 1)) / 2
		fuel += triangularNumber
	}
	fmt.Println("result:", fuel)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
