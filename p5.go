package main

import "fmt"

// http://projecteuler.net/problem=5
// 2013-06-28
// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

// numbers
// 20 -
// 19 *
// 18 *
// 17 *
// 16 *
// 15 *
// 14 *
// 13 *
// 12 *
// 11 *
// 10 - (20)
//  9 - (18)
//  8 - (16)
//  7 - (14)
//  6 - (12)
//  5 - (15)
//  4 - (16)
//  3 - (15)
//  2 - (20)
//  1 -

func runP5() {
  /*
  for i := 10;; i += 10 {
    fmt.Println(i)
    divisible := true
    for j := 9; j > 5; j-- {
      if i % j != 0 {
        divisible = false
        break
      }
    }
    if divisible {
      fmt.Println("Found it!")
      break
    }
  }
  */

  for i := 20;; i += 20 {
    if i % 2000000 == 0 {
      fmt.Println(i)
    }
    divisible := true
    for j := 19; j > 10; j-- {
      if i % j != 0 {
        divisible = false
        break
      }
    }
    if divisible {
      fmt.Printf("Found it: %d\n", i)
      break
    }
  }
}
