package util

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
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

func ParseSingleLineToIntSlice(pathToInputFile string) []int {
	_, filename, _, _ := runtime.Caller(1)

	file, err := os.Open(path.Join(path.Dir(filename), pathToInputFile))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fileContent := strings.TrimSpace(string(bytes))

	stringInts := strings.Split(fileContent, ",")
	ints := make([]int, len(stringInts))

	for i := range stringInts {
		ints[i], _ = strconv.Atoi(stringInts[i])
	}
	return ints
}
