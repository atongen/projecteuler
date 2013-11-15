package main

import (
  "fmt"
)

var p2fibMap = make(map[int]int)

func p2getFib(n int) (fib int) {
  if i, ok := p2fibMap[n]; ok {
    fib = i
  } else {
    //fmt.Printf("Evaluating %d\n", n)
    if n == 0 {
      fib = 0
    } else if n == 1 {
      fib = 1
    } else {
      fib = p2getFib(n-1) + p2getFib(n-2)
    }
    p2fibMap[n] = fib
  }
  return
}

func runP2() {
  sum := 0
  for i := 2;; i++ {
    n := p2getFib(i)
    if n <= 4000000 {
      if n % 2 == 0 {
        fmt.Println(n)
        sum += n
      }
    } else {
      break
    }
  }
  fmt.Printf("Sum: %d\n", sum)
}

func init() {
  addProb("2", func() { runP2() })
}
