package main

import "fmt"

type Hotel struct {
	name     string
	distance int
}

var hotels = []Hotel{
	{
		name:     "Hotel-1",
		distance: 11,
	},
	{
		name:     "Hotel-2",
		distance: 5,
	},
	{
		name:     "Hotel-3",
		distance: 11,
	},
	{
		name:     "Hotel-4",
		distance: 5,
	},
	{
		name:     "Hotel-5",
		distance: 0,
	},
}

func main() {
	fmt.Println(calcDistance("Hotel-5", "Hotel-1"))
}

func calcDistance(from, to string) int {
	var distance int
	var start bool = false
	var finished bool = false

	for !finished {
		for _, e := range hotels {
			if e.name == from {
				start = true
			}

			if e.name == to && start {
				distance += e.distance
				finished = true
				break
			}

			if start {
				distance += e.distance
			}
		}
	}

	return distance
}

//
