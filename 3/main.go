package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gwpmad/advent-of-code-2021/util"
)

func countBitByPosition(lines []string, bit string) [12]int {
	counts := [12]int{}

	for _, bits := range lines {
		for i, c := range bits {
			if string(c) == bit {
				counts[i]++
			}
		}
	}

	return counts
}

func isGreaterThanOrEqual(n1 float64, n2 float64) bool {
	return n1 >= n2
}

func isLessThanOrEqual(n1 float64, n2 float64) bool {
	return n1 <= n2
}

func findLineMatchingPriority(compareFunc func(float64, float64) bool, priorityBit string, otherBit string, lines []string) string {
	filteredLines := append([]string{}, lines...)

	for i := 0; len(filteredLines) > 1; i++ {
		counts := countBitByPosition(filteredLines, priorityBit)
		priorityBitCount := float64(counts[i])
		halfLinesLen := float64(len(filteredLines)) / 2.0
		var bitToUse string
		if compareFunc(priorityBitCount, halfLinesLen) {
			bitToUse = priorityBit
		} else {
			bitToUse = otherBit
		}

		n := 0
		for _, bits := range filteredLines {
			if string(bits[i]) == bitToUse {
				filteredLines[n] = bits
				n++
			}
		}
		filteredLines = filteredLines[:n]
	}
	return filteredLines[0]
}

func one(lines []string) {
	counts := countBitByPosition(lines, "0")

	halfLinesLen := len(lines) / 2
	gamma := ""
	epsilon := ""
	for _, count := range counts {
		if count > halfLinesLen {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaInt, _ := strconv.ParseInt(gamma, 2, 64) // always use 64 otherwise your number could be too high for the max of whatever bitsize you took
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println("result:", gammaInt*epsilonInt)
}

func two(lines []string) {
	ogc := findLineMatchingPriority(isGreaterThanOrEqual, "1", "0", lines)
	csr := findLineMatchingPriority(isLessThanOrEqual, "0", "1", lines)
	ogcInt, _ := strconv.ParseInt(ogc, 2, 64)
	csrInt, _ := strconv.ParseInt(csr, 2, 64)
	fmt.Println("result:", ogcInt*csrInt)
}

func main() {
	lines := util.ParseInputLinesToStringSlice("./input")

	switch os.Args[1] {
	case "1":
		one(lines)
	case "2":
		two(lines)
	}
}
