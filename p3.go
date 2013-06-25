package main

import (
  "fmt"
)

var p3primeMap = make(map[int]bool)

func p3isPrime(n int) bool {
  if prime, ok := p3primeMap[n]; ok {
    return prime
  } else {
    for i := 2; i < n/2; i++ {
      if n % i == 0 {
        p3primeMap[n] = false
        return false
      }
    }
    p3primeMap[n] = true
    return true
  }
}

func p3getFactors(n int) (factors []int) {
  factors = make([]int, n/2)
  for i := n/2; i >= 1; i-- {
    if n % i == 0 {
      pos := len(factors)
      fmt.Println(pos)
      if pos > cap(factors) {
        factors = append(factors, i)
      } else {
        factors[len(factors)+1] = i
      }
    }
  }
  return
}

func runP3() {
  ultimate := 600851475143
  for factor := range p3getFactors(ultimate) {
    if p3isPrime(factor) {
      fmt.Println(factor)
    }
  }
}
