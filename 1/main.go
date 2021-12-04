package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
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
	_, filename, _, _ := runtime.Caller(0)

	file, err := os.Open(path.Join(path.Dir(filename), "./input"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]int, 0)

	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case "1":
		one(lines)
	case "2":
		two(lines)
	}
}
