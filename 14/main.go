package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/gwpmad/advent-of-code-2021/util"
)

func main() {
	splitInput := util.ParseFileAndSplitByDelimiter("./input", "\n\n")
	template := splitInput[0]
	rules := strings.Split(splitInput[1], "\n")
	rulesMap := createRulesMap(rules)
	pairCounts := getPairCounts(rulesMap, template)
	letterCounts := getLetterCounts(template)

	switch os.Args[1] {
	case "1":
		one(pairCounts, letterCounts, rulesMap)
	case "2":
		two(pairCounts, letterCounts, rulesMap)
	}
}

type rulesMap map[string]string
type pairCounts map[string]int
type letterCounts map[string]int

func one(pairCounts pairCounts, letterCounts letterCounts, rulesMap rulesMap) {
	for i := 0; i < 10; i++ {
		pairCounts = processPairs(pairCounts, letterCounts, rulesMap)
	}

	maxCount, minCount := getMaxAndMinLetterCounts(letterCounts)
	fmt.Println("result:", maxCount-minCount)
}

func two(pairCounts pairCounts, letterCounts letterCounts, rulesMap rulesMap) {
	for i := 0; i < 40; i++ {
		pairCounts = processPairs(pairCounts, letterCounts, rulesMap)
	}

	maxCount, minCount := getMaxAndMinLetterCounts(letterCounts)
	fmt.Println("result:", maxCount-minCount)
}

func processPairs(pairCounts pairCounts, letterCounts letterCounts, rulesMap rulesMap) pairCounts {
	newPairCounts := getPairCounts(rulesMap, "")

	for pair, count := range pairCounts {
		middleChar := rulesMap[pair]
		newPair1, newPair2 := fmt.Sprintf("%v%v", string(pair[0]), middleChar), fmt.Sprintf("%v%v", middleChar, string(pair[1]))
		newPairCounts[newPair1] += count
		newPairCounts[newPair2] += count
		letterCounts[middleChar] += count
	}
	return newPairCounts
}

func getMaxAndMinLetterCounts(letterCounts letterCounts) (int, int) {
	maxCount, minCount := 0, math.MaxInt64
	for _, count := range letterCounts {
		if count > maxCount {
			maxCount = count
		}
		if count < minCount {
			minCount = count
		}
	}
	return maxCount, minCount
}

func getLetterCounts(str string) letterCounts {
	letterCounts := letterCounts{}
	for _, rune := range str {
		letter := string(rune)
		letterCounts[letter]++
	}
	return letterCounts
}

func createRulesMap(rules []string) rulesMap {
	rulesMap := rulesMap{}
	for _, rule := range rules {
		splitRule := strings.Split(rule, " -> ")
		rulesMap[splitRule[0]] = splitRule[1]
	}
	return rulesMap
}

func getPairCounts(rulesMap rulesMap, str string) pairCounts {
	pairCounts := pairCounts{}
	for pair := range rulesMap {
		pairCounts[pair] = 0
	}
	for i := 0; i < len(str)-1; i++ {
		pair := fmt.Sprintf("%v%v", string(str[i]), string(str[i+1]))
		pairCounts[pair]++
	}
	return pairCounts
}
