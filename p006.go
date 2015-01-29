package main

// Problem #6
// http://projecteuler.net/problem=6
// 2013-06-30
//
// The sum of the squares of the first ten natural numbers is,
// 12 + 22 + ... + 102 = 385
//
// The square of the sum of the first ten natural numbers is,
// (1 + 2 + ... + 10)2 = 552 = 3025
//
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025  385 = 2640.
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

import "fmt"

func runP6() {
	r := 100
	//r := 10
	a := squareOfSum(r)
	b := sumOfSquares(r)
	c := a - b
	fmt.Printf("(%d) %d - %d = %d\n", r, a, b, c)
}

func init() {
	addProb("6", func() { runP6() })
}
