package main

import (
  "fmt"
)

func p3isPrime(n uint64) bool {
  if n > 2 {
    var i uint64 = 2
    for ; i < n/2; i++ {
      if n % i == 0 {
        return false
      }
    }
  }
  return true
}

func runP3() {
  var ultimate uint64 = 600851475143
  fmt.Printf("Largest prime factor of %d\n", ultimate)

  var i uint64 = ultimate / 2
  if i % 2 == 0 {
    i--
  }
  for ; i > 2; i -= 2 {
    if (i-1) % 1000000 == 0 {
      fmt.Printf("%d...\n", i)
    }
    if ultimate % i == 0 {
      if p3isPrime(i) {
        fmt.Println(i)
        break
      }
    }
  }
}
