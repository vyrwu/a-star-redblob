package main

import (
	"fmt"
)

func main() {
	maxX := 10
	maxY := 10
	grid := SquareGrid{
		width:  maxX,
		height: maxY,
		walls: []Location{
			{1, 7},
			{1, 8},
			{2, 7},
			{2, 8},
			{3, 7},
			{3, 8},
		},
		forests: []Location{
			{4, 6},
			{4, 7},
			{5, 6},
			{5, 7},
			{6, 6},
			{6, 7},
			{4, 4},
			{4, 5},
			{5, 4},
			{5, 5},
			{6, 4},
			{6, 5},
			{4, 2},
			{4, 3},
			{5, 2},
			{5, 3},
			{6, 2},
			{6, 3},
		},
	}

	start := Location{7, 8}
	goal := Location{1, 4}
	path := aStarSearch(grid, start, goal)

	// now draw a map based on grid and shortest path found by A*
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			l := Location{
				x,
				y,
			}
			if l.y == start.y && l.x == start.x {
				fmt.Print(" S ")
			} else if l.y == goal.y && l.x == goal.x {
				fmt.Print(" G ")
			} else if contains(grid.walls, l) {
				fmt.Print(" X ")
			} else if contains(grid.forests, l) {
				fmt.Print(" ^ ")
			} else if contains(path, l) {
				fmt.Print(" * ")
			} else {
				fmt.Print(" . ")
			}
		}
		fmt.Println()
	}
}
