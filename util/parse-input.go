package util

import (
	"bufio"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
)

func ParseInputLinesToIntSlice(pathToInputFile string) []int {
	_, filename, _, _ := runtime.Caller(1)

	file, err := os.Open(path.Join(path.Dir(filename), pathToInputFile))
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

	return lines
}

func ParseInputLinesToStringSlice(pathToInputFile string) []string {
	_, filename, _, _ := runtime.Caller(1)

	file, err := os.Open(path.Join(path.Dir(filename), pathToInputFile))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
