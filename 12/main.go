package main

import (
	"log"
	"os"
	"strings"

	"github.com/gwpmad/advent-of-code-2021/util"
)

func main() {
	lines := util.ParseInputLinesToStringSlice("./input")
	graph := createGraph(lines)

	switch os.Args[1] {
	case "1":
		one(graph)
	case "2":
		two(graph)
	}
}

func one(graph caveGraph) {
	totalPaths := 0
	queue := make([]queueItem, len(graph["start"].neighbours))

	i := 0
	for _, caveName := range graph["start"].neighbours {
		queue[i] = queueItem{
			caveName:          caveName,
			visitedSmallCaves: map[string]struct{}{"start": {}},
		}
		i++
	}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		cave := graph[item.caveName]
		newVisitedSmallCaves := copyMap(item.visitedSmallCaves)
		if cave.small {
			newVisitedSmallCaves[item.caveName] = struct{}{}
		}

		for _, neighbourCaveName := range graph[item.caveName].neighbours {
			if neighbourCaveName == "end" {
				totalPaths++
				continue
			}

			if _, ok := item.visitedSmallCaves[neighbourCaveName]; ok {
				continue
			}
			queue = append(queue, queueItem{
				caveName:          neighbourCaveName,
				visitedSmallCaves: newVisitedSmallCaves,
			})
		}
	}
	log.Printf("result: %v\n", totalPaths)
}

func two(graph caveGraph) {
	totalPaths := 0
	queue := make([]queueItem, len(graph["start"].neighbours))

	i := 0
	for _, caveName := range graph["start"].neighbours {
		queue[i] = queueItem{
			caveName:          caveName,
			visitedSmallCaves: map[string]struct{}{"start": {}},
		}
		i++
	}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		cave := graph[item.caveName]
		newVisitedSmallCaves := copyMap(item.visitedSmallCaves)
		if cave.small {
			if _, ok := newVisitedSmallCaves[item.caveName]; ok {
				item.smallCaveVisitedTwice = true
			} else {
				newVisitedSmallCaves[item.caveName] = struct{}{}
			}
		}

		for _, neighbourCaveName := range graph[item.caveName].neighbours {
			if neighbourCaveName == "end" {
				totalPaths++
				continue
			}
			if _, ok := newVisitedSmallCaves[neighbourCaveName]; ok && (neighbourCaveName == "start" || item.smallCaveVisitedTwice) {
				continue
			}
			queue = append(queue, queueItem{
				caveName:              neighbourCaveName,
				visitedSmallCaves:     newVisitedSmallCaves,
				smallCaveVisitedTwice: item.smallCaveVisitedTwice,
			})
		}
	}
	log.Printf("result: %v\n", totalPaths)
}

type cave struct {
	neighbours []string
	small      bool
}
type caveGraph map[string]cave
type queueItem struct {
	caveName              string
	visitedSmallCaves     map[string]struct{}
	smallCaveVisitedTwice bool
}

func createGraph(lines []string) caveGraph {
	graph := caveGraph{}

	for _, line := range lines {
		splitLine := strings.Split(line, "-")
		associateCaves(graph, splitLine[0], splitLine[1])
	}
	return graph
}

func associateCaves(graph caveGraph, cave1, cave2 string) {
	cave1Entry, cave1Ok := graph[cave1]
	cave2Entry, cave2Ok := graph[cave2]
	if !cave1Ok {
		cave1Entry.small = isSmall(cave1)
	}
	if !cave2Ok {
		cave2Entry.small = isSmall(cave2)
	}

	cave1Entry.neighbours = append(cave1Entry.neighbours, cave2)
	cave2Entry.neighbours = append(cave2Entry.neighbours, cave1)
	graph[cave1] = cave1Entry
	graph[cave2] = cave2Entry
}

func isSmall(caveName string) bool {
	return caveName == strings.ToLower(caveName)
}

func copyMap(originalMap map[string]struct{}) map[string]struct{} {
	copyMap := map[string]struct{}{}
	for key, value := range originalMap {
		copyMap[key] = value
	}
	return copyMap
}
