/*
	A* algorithm using pseudocode from here: https://en.wikipedia.org/wiki/A*_search_algorithm
	Trying every path would take way too long so the algorithm gives up on paths it finds less efficient

	From Wikipedia, succinctly put:
	Specifically, A* selects the path that minimizes
	f(n) = g(n) + h(n)
	where n is the next node on the path, g(n) is the cost of the path from the start node to n, and h(n) is a heuristic
	function that estimates the cost of the cheapest path from n to the goal.

	So g(n) is the distance to the grid square and h(n) is the estimated distance from the grid square to the goal.
	For each square, we calculate the g score of all its neighbours based on the path we're on. This is the total distance
	to the neighbour plus the estimated distance from the neighbour to the goal.
	If the g score we get for a neighbour is better than any other g score we've had for that neighbour via any other path,
	we enqueue it (otherwise we stop bothering with this path - therein lies the algorithm's optimisation).
	The g score of the goal grid square is our answer. At this point there is no subjectivity (i.e. h(n)) included in the g
	score.

	When we enqueue a neighbour we also store an estimated f(n) score - this is the (known) g score (the total so far +
	the value of the neighbour) plus the heuristic function's output for the neighbour (an estimate of the remaining
	distance to the goal - we don't know what the actual values will be going forward so our best guess is just the
	Manhattan distance, which implicitly assumes every weight/grid square value/weight will be 1 going forward)

	A* is guaranteed to return a least-cost path from start to goal if the heuristic function is admissible, meaning
	that it	never overestimates the actual cost to get to the goal. Our heuristic function fulfils this criterion.

	The number value of each square (node) on the grid is equivalent of the 'weight' of the edge required to get to it
	if it were a graph.

	A priority queue is used for further optimisation, as a concept it is not necessary for the algorithm. The first
	item the priority queue	returns is the one with the lowest f score, so the one that is estimated to be part of the
	cheapest overall path	to the goal. This is vital for A* to return the right answer (otherwise we may get to the
	end of a longer path first). However, the	priority queue structure is not necessarily required - a regular array
	would be fine so long as we sorted it correctly by the stated criteria.

	The priority queue uses a min heap. Golang info/implementation used is from here: https://pkg.go.dev/container/heap

	This was a particularly tricky advent of code puzzle because a brute force graph solution (all paths) worked for part
	1 on the example,	but not on the real data (a much larger grid - the code may have taken days to finish). This is why
	the A* algorithm is required.	Part 2 had another trick - the efficiency of only traversing right or downward (which
	worked for part 1) produces too high an answer (2816). Adding the ability to go left or upward reveals that there are
	six left/upward movements required for the lowest answer (2809).
*/

package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/gwpmad/advent-of-code-2021/util"
)

type pqItem struct {
	coords string
	index  int // The index is needed by update and is maintained by the heap.Interface methods.
}

// A PriorityQueue implements heap.Interface and holds pqItems.
type PriorityQueue []*pqItem

func (pq PriorityQueue) Len() int { return len(pq) }

// Decides the priority of items. Here we want to prioritise those with a lower fScore.
// fScore[n] represents our current best guess as to how short a path from start to
// finish can be if it goes through n.
func (pq PriorityQueue) Less(i, j int) bool {
	return fScore[pq[i].coords] < fScore[pq[j].coords]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// heap.Push() calls this method
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*pqItem) // x is the item. My understanding is that the .(*pqItem) on the end is like writing "as pqItem" in TypeScript - the func doesn't know what the interface is otherwise
	item.index = n
	*pq = append(*pq, item)
	pqSet[item.coords] = struct{}{}
}

// will be used by heap.Pop() to return the 'minimum' item on the queue based on the Less() func defined above, in our case the one with the lowest value
// therefore our graph won't be traversed by DFS or BFS but simply by which 'neighbours' have the lowest value - this is the efficiency improvement that
// the priority queue offers to our A* algorithm
func (pq *PriorityQueue) Pop() interface{} { // it returns an initialised interface of some sort
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	delete(pqSet, item.coords)
	return item
}

var pqSet = map[string]struct{}{}

// For node n, fScore[n] := gScore[n] + h(n). fScore[n] represents our current best guess as to
// how short a path from start to finish can be if it goes through n.
// We use the fScore to decide priority in the queue. As the pseudocode says: current := the node in openSet having the lowest fScore[] value
var fScore = map[string]int{}

func main() {
	lines := util.ParseInputLinesToStringSlice("./input")
	grid := createGrid(lines)

	switch os.Args[1] {
	case "1":
		findLowestPathTotal(grid)
	case "2":
		grid = getExpandedGrid(grid, len(lines[0]))
		findLowestPathTotal(grid)
	}
}

func findLowestPathTotal(grid grid) {
	maxSideIdx := getMaxSideIdxForSquareGrid(grid)
	result := getLowestTotalUsingAStarAlgorithm(grid, maxSideIdx)
	fmt.Println("result:", result)
}

// Based on 'Pseudocode' function from here https://en.wikipedia.org/wiki/A*_search_algorithm
func getLowestTotalUsingAStarAlgorithm(grid grid, maxSideIdx int) (result int) {
	goalCoords := getCoords(maxSideIdx, maxSideIdx)

	openSet := PriorityQueue{}
	heap.Push(&openSet, &pqItem{
		coords: "0,0",
		index:  0,
	})

	fScore["0,0"] = heuristic("0,0", maxSideIdx, grid)
	// For node n, gScore[n] is the cost of the cheapest path from start to n currently known.
	gScore := map[string]int{"0,0": 0}

	for openSet.Len() > 0 {
		current := heap.Pop(&openSet).(*pqItem)
		if current.coords == goalCoords {
			return gScore[current.coords]
		}

		for _, neighbour := range grid[current.coords].neighbours {
			// d(current,neighbor) is the weight of the edge from current to neighbour - in this case the number value of the neighbour
			// tentative_gScore is the distance from start to the neighbor through current
			tentativeGScore := gScore[current.coords] + grid[neighbour].value
			if neighbourGScore, ok := gScore[neighbour]; !ok || tentativeGScore < neighbourGScore {
				// This path to neighbour is better than any previous one. Record it!
				gScore[neighbour] = tentativeGScore
				fScore[neighbour] = tentativeGScore + heuristic(neighbour, maxSideIdx, grid)

				// We don't push to queue if the item is already there. This is never the case in part one, but in part two we avoid 14841
				// queue pushes because of this (though the priority queue means some items may never be processed, partly making up for it).
				// This efficiency can save us up to 50ms in part two.
				if _, ok := pqSet[neighbour]; !ok {
					heap.Push(&openSet, &pqItem{
						coords: neighbour,
					})
				}
			}
		}
	}
	panic("Never reached the target")
}

// heuristic should return the estimated cost of the path from n to the goal. Given we do not know the values of future grid squares,
// in a square grid with 4 directions of movement the best estimate is the Manhattan distance.
func heuristic(coords string, maxSideIdx int, grid grid) int {
	x, y := splitCoords(coords)
	return (maxSideIdx - x) + (maxSideIdx - y)
}

type gridEntry struct {
	value      int
	neighbours []string
}
type grid map[string]gridEntry

func createGrid(lines []string) grid {
	grid := grid{}

	for y, line := range lines {
		numbers := strings.Split(line, "")
		for x, num := range numbers {
			coords := getCoords(x, y)
			entry := grid[coords]
			entry.value, _ = strconv.Atoi(num)
			addNeighboursAndSave(x, y, len(numbers), len(lines), entry, coords, grid)
		}
	}

	return grid
}

func getExpandedGrid(oldGrid grid, oldSideLen int) grid {
	tileCount := 5
	sideLen := oldSideLen * tileCount
	grid := grid{}
	for key := range oldGrid {
		oldX, oldY := splitCoords(key)
		for j := 0; j < tileCount; j++ {
			for k, l := j, 0; l <= j; l++ {
				x, y := oldX+(l*oldSideLen), oldY+(k*oldSideLen)
				coords := getCoords(x, y)
				entry := grid[coords]
				entry.value = incrementValue(oldGrid[key].value, k+l)
				addNeighboursAndSave(x, y, sideLen, sideLen, entry, coords, grid)
				if k == l {
					continue
				}

				x, y = oldX+(k*oldSideLen), oldY+(l*oldSideLen)
				coords = getCoords(x, y)
				entry = grid[coords]
				entry.value = incrementValue(oldGrid[key].value, k+l)
				addNeighboursAndSave(x, y, sideLen, sideLen, entry, coords, grid)
			}
		}
	}

	return grid
}

func addNeighboursAndSave(x int, y int, maxX int, maxY int, entry gridEntry, coords string, grid grid) {
	if x < (maxX - 1) {
		entry.neighbours = append(entry.neighbours, getCoords(x+1, y))
	}

	if x > 0 {
		entry.neighbours = append(entry.neighbours, getCoords(x-1, y))
	}

	if y < (maxY - 1) {
		entry.neighbours = append(entry.neighbours, getCoords(x, y+1))
	}

	if y > 0 {
		entry.neighbours = append(entry.neighbours, getCoords(x, y-1))
	}
	grid[coords] = entry
}

func getCoords(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func splitCoords(coords string) (x, y int) {
	splitCoords := strings.Split(coords, ",")
	x, _ = strconv.Atoi(splitCoords[0])
	y, _ = strconv.Atoi(splitCoords[1])
	return x, y
}

func incrementValue(value int, incrementBy int) int {
	newValue := value + incrementBy
	if newValue > 9 {
		newValue = newValue - 9
	}
	return newValue
}

func getMaxSideIdxForSquareGrid(grid grid) int {
	sqrt := math.Sqrt(float64(len(grid)))
	return int(sqrt) - 1
}
