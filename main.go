package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	grid := SquareGrid{
		width:  10,
		height: 10,
		walls: []Location{
			{1, 7},
			{1, 8},
			{2, 7},
			{2, 8},
			{3, 7},
			{3, 8},
		},
		forests: []Location{},
	}

	start := Location{7, 8}
	goal := Location{1, 4}
	path := aStarSearch(grid, start, goal)
	fmt.Printf("%d", path)
}

type Location struct {
	x int
	y int
}

func (l Location) Equals(o Location) bool {
	if l.x == o.x && l.y == o.y {
		return true
	}
	return false
}

type GridWithWeights interface {
	Cost(from, to Location) int
	Neighbors(id Location) []Location
}

type SquareGrid struct {
	width   int
	height  int
	walls   []Location
	forests []Location
}

func (sg SquareGrid) InBounds(id Location) bool {
	return inRange(0, id.x, sg.width) && inRange(0, id.y, sg.height)
}

func (sg SquareGrid) Passable(id Location) bool {
	for _, v := range sg.walls {
		if v.x == id.x && v.y == id.y {
			return false
		}
	}
	return true
}

func (sg SquareGrid) Neighbors(id Location) []Location {
	x := id.x
	y := id.y
	results := []Location{
		{x + 1, y},
		{x, y - 1},
		{x - 1, y},
		{x, y + 1},
	}
	if (x+y)%2 == 0 {
		results = reverse(results)
	}
	results = filter(results, sg.InBounds)
	results = filter(results, sg.Passable)
	return results
}

// For the case of the example, cost function is very simple.
// Cost of stepping into the forest is 5, any other step is 1.
func (sg SquareGrid) Cost(from, to Location) int {
	if contains(sg.forests, to) {
		return 5
	}
	return 1
}

func contains(slice []Location, i Location) bool {
	for _, v := range slice {
		if i == v {
			return true
		}
	}
	return false
}

func filter(a []Location, f func(x Location) bool) []Location {
	b := a[:0]
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}

func reverse(a []Location) []Location {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a
}

func inRange(min, target, max int) bool {
	if (min <= target) && (target < max) {
		return true
	}
	return false
}

//PriorityQueue puts lowest priority Items on the head of the Queue
type Item struct {
	value    interface{} // The value of an item; arbitrary.
	priority int         // The priority of the item in the queue.
	index    int         // The index of the item in the heap.
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

// We want smallest cost first; i should be sorted first
// only if its priority < j.priority
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// aStarSearch performs A* search on Weighed Grid, looking for the least costly path from start to goal location.
// Return least costly path from start to goal in slice of Locations.
func aStarSearch(grid GridWithWeights, start, goal Location) []Location {
	frontier := make(PriorityQueue, 0)
	heap.Init(&frontier)

	heap.Push(&frontier, &Item{
		value:    start,
		priority: 0,
	})

	cameFrom := map[Location]Location{start: start}
	costSoFar := map[Location]int{start: 0}

	for frontier.Len() > 0 {
		current := heap.Pop(&frontier).(*Item).value.(Location)
		if current.Equals(goal) {
			break
		}
		for _, next := range grid.Neighbors(current) {
			newCost := costSoFar[current] + grid.Cost(current, next)
			if _, exists := costSoFar[next]; !exists || newCost < costSoFar[next] {
				costSoFar[next] = newCost
				priority := newCost + heuristic(next, goal)
				heap.Push(&frontier, &Item{
					value:    next,
					priority: priority,
				})
				cameFrom[next] = current
			}
		}
	}

	//Now we are done, we found the path. Go backwards from the goal node to beggining and return nice list
	current := goal
	path := make([]Location, 0)
	for !current.Equals(start) {
		path = append(path, current)
		current = cameFrom[current]
	}
	path = append(path, start)
	path = reverse(path)
	return path
}

func heuristic(a, b Location) int {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	h := math.Abs(dx) + math.Abs(dy)
	return int(h)
}
