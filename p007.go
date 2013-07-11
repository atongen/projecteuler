package main

/*
http://projecteuler.net/problem=7
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
What is the 10 001st prime number?
*/

import "fmt"

func runP7() {
  var p uint64 = 1
  //t := 6
  t := 10001
  for i := 1; i <= t; i++ {
    p = p3nextPrime(uint64(p))
    fmt.Printf("%d %d\n", i, p)
  }
}
