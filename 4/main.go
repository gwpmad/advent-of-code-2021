package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gwpmad/advent-of-code-2021/util"
)

type numberSet map[int]struct{}
type bingoCard struct {
	allNumbers numberSet
	sets       [10]numberSet
}
type cardResult struct {
	remainingNumbers numberSet
	lastNumberDrawn  int
}

func main() {
	lines := util.ParseInputLinesToStringSlice("./input")
	numbersDrawn := getNumbersDrawn(lines[0])
	bingoCards := parseBingoCards(lines[2:])

	w, l := playGame(numbersDrawn, bingoCards)
	switch os.Args[1] {
	case "1":
		one(w)
	case "2":
		two(l)
	}
}

func one(winner cardResult) {
	multpliedResult := getMultipliedResult(winner)
	fmt.Println("result:", multpliedResult)
}

func two(loser cardResult) {
	multpliedResult := getMultipliedResult(loser)
	fmt.Println("result:", multpliedResult)
}

func getNumbersDrawn(numbers string) []int {
	numbersSlice := strings.Split(numbers, ",")
	ints := make([]int, len(numbersSlice))
	for i, s := range numbersSlice {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}

func getMultipliedResult(result cardResult) int {
	sumOfRemaining := 0
	for k := range result.remainingNumbers {
		sumOfRemaining += k
	}
	return result.lastNumberDrawn * sumOfRemaining
}

func playGame(numbersDrawn []int, bingoCards []bingoCard) (winner, loser cardResult) {
	remainingCards := map[int]struct{}{}
	for i := range bingoCards {
		remainingCards[i] = struct{}{}
	}
top:
	for _, numberDrawn := range numbersDrawn {
		for cardIdx, card := range bingoCards {
			delete(card.allNumbers, numberDrawn)
			for _, set := range card.sets {
				delete(set, numberDrawn)
				if len(set) != 0 {
					continue
				}

				if len(remainingCards) == len(bingoCards) {
					winner.remainingNumbers = copyNumberSet(bingoCards[cardIdx].allNumbers)
					winner.lastNumberDrawn = numberDrawn
				}
				if len(remainingCards) == 1 {
					if _, ok := remainingCards[cardIdx]; ok {
						loser.remainingNumbers = copyNumberSet(bingoCards[cardIdx].allNumbers)
						loser.lastNumberDrawn = numberDrawn
						break top
					}
				}
				delete(remainingCards, cardIdx)
			}
		}
	}
	return winner, loser
}

func copyNumberSet(set numberSet) numberSet {
	copy := make(numberSet)
	for key, value := range set {
		copy[key] = value
	}
	return copy
}

func parseBingoCards(lines []string) []bingoCard {
	bingoCards := []bingoCard{}

	wsRegexp := regexp.MustCompile(`\s+`)
	rowNum := 5
	card := getBingoCard()
	for _, line := range lines {
		if len(line) == 0 {
			rowNum = 5
			bingoCards = append(bingoCards, card)
			card = getBingoCard()
			continue
		}
		splitLine := wsRegexp.Split(strings.TrimSpace(line), -1)
		for colNum, n := range splitLine {
			intN, _ := strconv.Atoi(n)
			card.sets[rowNum][intN] = struct{}{}
			card.sets[colNum][intN] = struct{}{}
			card.allNumbers[intN] = struct{}{}
		}
		rowNum++
	}
	bingoCards = append(bingoCards, card)
	return bingoCards
}

func getBingoCard() bingoCard {
	card := bingoCard{}
	card.allNumbers = numberSet{}
	for i := 0; i < 10; i++ {
		card.sets[i] = numberSet{}
	}
	return card
}
