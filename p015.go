package main

import "errors"

/**
 * Problem #15 - Lattice paths
 * https://projecteuler.net/problem=15
 *
 * Starting in the top left corner of a 2×2 grid,
 * and only being able to move to the right and down,
 * there are exactly 6 routes to the bottom right corner.
 *
 * How many such routes are there through a 20×20 grid?
 */

type entry struct {
	x int
	y int
	v int
}

var entries = []entry{}

func p015GetEntry(x int, y int) (int, error) {
	var min, max int
	if x <= y {
		min = x
		max = y
	} else {
		min = y
		max = x
	}

	for _, entry := range entries {
		if entry.x == min && entry.y == max {
			return entry.v, nil
		}
	}

	return 0, errors.New("no entry")
}

func p015SetEntry(x int, y int, v int) {
	var min, max int
	if x <= y {
		min = x
		max = y
	} else {
		min = y
		max = x
	}

	entries = append(entries, entry{min, max, v})
}

/**
 * f(x,y) = f(x-1,y)+f(x,y-1)
 */
func p015(x int, y int) int {
	if x == 1 && y == 1 {
		return 2
	} else if x == 1 {
		return y + 1
	} else if y == 1 {
		return x + 1
	} else {
		v1, err := p015GetEntry(x, y-1)
		if err != nil {
			v1 = p015(x, y-1)
			p015SetEntry(x, y-1, v1)
		}

		v2, err := p015GetEntry(x-1, y)
		if err != nil {
			v2 = p015(x-1, y)
			p015SetEntry(x-1, y, v2)
		}

		return v1 + v2
	}
}

// Solution: 137846528820
func runP15() {
	println(p015(20, 20))
}

func init() {
	addProb("15", func() { runP15() })
}
