package main

import (
	"container/heap"
	"math"
)

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
