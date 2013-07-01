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

func p6sumOfSquares(n int) uint64 {
  var i uint64 = 0
  for j := 1; j <= n; j++ {
    i += uint64(j*j)
  }
  return i
}

func p6squareOfSum(n int) uint64 {
  var sum uint64 = 0
  for i := 1; i <= n; i++ {
    sum += uint64(i)
  }
  return sum*sum
}

func runP6() {
  r := 100
  //r := 10
  a := p6squareOfSum(r)
  b := p6sumOfSquares(r)
  c := a - b
  fmt.Printf("(%d) %d - %d = %d\n", r, a, b, c)
}
