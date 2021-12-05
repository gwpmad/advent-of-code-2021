package main

import (
	"fmt"
	"os"

	"github.com/gwpmad/advent-of-code-2021/util"
)

func one(lines []int) {
	count := 0
	for i := 1; i < len(lines); i++ {
		if lines[i] > lines[i-1] {
			count++
		}
	}

	fmt.Println("count:", count)
}

func two(lines []int) {
	count := 0
	for i := 1; i < len(lines)-2; i++ {
		previousGroupSum := lines[i-1] + lines[i] + lines[i+1]
		currentGroupSum := lines[i] + lines[i+1] + lines[i+2]
		if currentGroupSum > previousGroupSum {
			count++
		}
	}

	fmt.Println("count:", count)
}

func main() {
	lines := util.ParseInputLinesToIntSlice("./input")

	switch os.Args[1] {
	case "1":
		one(lines)
	case "2":
		two(lines)
	}
}
