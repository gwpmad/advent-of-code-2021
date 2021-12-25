package main

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/gwpmad/advent-of-code-2021/util"
)

func main() {
	lines := util.ParseInputLinesToStringSlice("./input")
	switch os.Args[1] {
	case "1":
		one(lines)
	case "2":
		two(lines)
	}
}

func one(lines []string) {
	uniques := 0
	for _, line := range lines {
		splitLine := strings.Split(line, " | ")
		outputValues := strings.Split(splitLine[1], " ")
		for _, str := range outputValues {
			length := len(string(str))
			if length == 2 || length == 3 || length == 4 || length == 7 {
				uniques++
			}
		}
	}
	fmt.Println("result:", uniques)
}

func two(lines []string) {
	charRegexp := regexp.MustCompile(`[a-zA-Z]+`)

	total := 0
	for _, line := range lines {
		patterns := charRegexp.FindAllStringSubmatch(line, -1)

		patternsMap := getPatternsMap(patterns)
		valuePatternsMap := determinePatternNumbers(patternsMap)

		outputPatterns := charRegexp.FindAllStringSubmatch(strings.Split(line, "|")[1], -1)

		outputValueString := ""
		for _, outputPatternSlice := range outputPatterns {
			outputPattern := stringToCharMap(outputPatternSlice[0])
			for value, pattern := range valuePatternsMap {
				if reflect.DeepEqual(outputPattern, pattern) {
					outputValueString = fmt.Sprintf("%s%s", outputValueString, strconv.Itoa(value))
					break
				}
			}
		}

		outputValue, _ := strconv.Atoi(outputValueString)
		total += outputValue
	}

	fmt.Println("result:", total)
}

type charMap map[string]struct{}
type lengthToPatternsMap map[int][]charMap
type valueToPatternMap map[int]charMap

func getPatternsMap(patterns [][]string) lengthToPatternsMap {
	patternsMap := lengthToPatternsMap{}

	for _, patternSlice := range patterns {
		pattern := patternSlice[0]
		length := len(pattern)
		patternsMap[length] = append(patternsMap[length], stringToCharMap(pattern))
	}
	return patternsMap
}

func stringToCharMap(str string) charMap {
	charMap := charMap{}
	for _, rune := range str {
		charMap[string(rune)] = struct{}{}
	}
	return charMap
}

func determinePatternNumbers(patternsMap lengthToPatternsMap) valueToPatternMap {
	result := valueToPatternMap{}
	result[1] = patternsMap[2][0]
	result[4] = patternsMap[4][0]
	result[7] = patternsMap[3][0]
	result[8] = patternsMap[7][0]

	processSixCharItems(patternsMap, result)
	processFiveCharItems(patternsMap, result)
	return result
}

func processSixCharItems(patternsMap lengthToPatternsMap, result valueToPatternMap) {
	for _, charMap := range patternsMap[6] {
		var zeroFound, sixFound, nineFound bool
		if charMapContainsAnother(charMap, result[4]) {
			result[9] = charMap
			nineFound = true
		} else if charMapContainsAnother(charMap, result[1]) {
			result[0] = charMap
			zeroFound = true
		} else {
			result[6] = charMap
			sixFound = true
		}
		if zeroFound && nineFound && sixFound {
			break
		}
	}
}

func processFiveCharItems(patternsMap lengthToPatternsMap, result valueToPatternMap) {
	for _, charMap := range patternsMap[5] {
		var twoFound, threeFound, fiveFound bool
		if charMapContainsAnother(charMap, result[1]) {
			result[3] = charMap
			threeFound = true
		} else if charMapContainsAnother(result[6], charMap) {
			result[5] = charMap
			fiveFound = true
		} else {
			result[2] = charMap
			twoFound = true
		}
		if twoFound && threeFound && fiveFound {
			break
		}
	}
}

func charMapContainsAnother(mapA charMap, mapB charMap) bool {
	for key, _ := range mapB {
		_, ok := mapA[key]
		if !ok {
			return false
		}
	}
	return true
}
