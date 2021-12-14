package main

import (
	"fmt"
	"os"

	"github.com/gwpmad/advent-of-code-2021/util"
)

type intsMap map[int]int

func main() {
	ints := util.ParseSingleLineToIntSlice("./input")
	intsMap := createIntsMap(ints)
	switch os.Args[1] {
	case "1":
		countFish(intsMap, 80)
	case "2":
		countFish(intsMap, 256)
	}
}

func countFish(intsMap intsMap, days int) {
	for i := 0; i < days; i++ {
		passDay(intsMap)
	}
	result := 0
	for _, count := range intsMap {
		result += count
	}
	fmt.Println("result:", result)
}

func passDay(intsMap intsMap) {
	zeroCount := intsMap[0]
	tmpInt := zeroCount
	for i := 8; i >= 0; i-- {
		tmp2 := intsMap[i]
		intsMap[i] = tmpInt
		tmpInt = tmp2
		if i == 6 {
			intsMap[i] += zeroCount
		}
	}
}

func createIntsMap(ints []int) intsMap {
	newMap := intsMap{}
	for _, int := range ints {
		newMap[int]++
	}
	return newMap
}
