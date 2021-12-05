package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gwpmad/advent-of-code-2021/util"
)

func one(lines []string) {
	horizontalPosition := 0
	depth := 0

	for _, instruction := range lines {
		splitInstruction := strings.Split(instruction, " ")
		distance, _ := strconv.Atoi(splitInstruction[1])

		switch splitInstruction[0] {
		case "up":
			depth -= distance
		case "down":
			depth += distance
		case "forward":
			horizontalPosition += distance
		}
	}
	fmt.Println("result:", horizontalPosition*depth)
}

func two(lines []string) {
	aim := 0
	horizontalPosition := 0
	depth := 0

	for _, instruction := range lines {
		splitInstruction := strings.Split(instruction, " ")
		units, _ := strconv.Atoi(splitInstruction[1])

		switch splitInstruction[0] {
		case "up":
			aim -= units
		case "down":
			aim += units
		case "forward":
			horizontalPosition += units
			depth += (aim * units)
		}
	}
	fmt.Println("result:", horizontalPosition*depth)
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
