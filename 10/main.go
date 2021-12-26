package main

import (
	"fmt"
	"os"
	"sort"

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

type symbolMetadata struct {
	opposite string
	score    int
}

var closingSymbolMap = map[string]symbolMetadata{
	")": {opposite: "(", score: 3},
	"]": {opposite: "[", score: 57},
	"}": {opposite: "{", score: 1197},
	">": {opposite: "<", score: 25137},
}

var openingSymbolMap = map[string]symbolMetadata{
	"(": {score: 1},
	"[": {score: 2},
	"{": {score: 3},
	"<": {score: 4},
}

func one(lines []string) {
	totalScore := 0
	for _, line := range lines {
		corruptScore, _ := getStackFromLine(line)
		totalScore += corruptScore
	}
	fmt.Println("result:", totalScore)
}

func two(lines []string) {
	scores := []int{}
	for _, line := range lines {
		score := 0
		corruptScore, remainingStack := getStackFromLine(line)
		if corruptScore != 0 {
			continue
		}

		for i := len(remainingStack) - 1; i >= 0; i-- {
			symbolScore := openingSymbolMap[remainingStack[i]].score
			score *= 5
			score += symbolScore
		}
		scores = append(scores, score)
	}
	sort.Slice(scores, func(i, j int) bool { return scores[i] < scores[j] })
	fmt.Println("result:", scores[len(scores)/2])
}

func getStackFromLine(line string) (corruptScore int, remainingStack []string) {
	stack := []string{}

	for _, rune := range line {
		char := string(rune)
		if metadata, ok := closingSymbolMap[char]; ok {
			stackLastIdx := len(stack) - 1
			lastStackElement := stack[stackLastIdx]
			if lastStackElement != metadata.opposite {
				corruptScore = metadata.score
				break
			}
			stack = stack[:stackLastIdx]
		} else {
			stack = append(stack, char)
		}
	}
	return corruptScore, stack
}
