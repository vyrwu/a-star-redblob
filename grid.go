package main

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
