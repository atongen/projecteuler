package main

import (
  "fmt"
)

/*
http://projecteuler.net/problem=10
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
Find the sum of all the primes below two million.
*/

func runP10() {
  var sum uint64 = 0
  var max uint64 = 2000000
  //var max uint64 = 10
  var n uint64 = 2
  for ; n < max ; n = p3nextPrime(uint64(n)) {
    fmt.Println(n)
    sum += n
  }
  fmt.Println(sum)
}
