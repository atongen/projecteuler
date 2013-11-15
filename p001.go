package main

import (
  "fmt"
)

func runP1() {
  sum := 0
  for i := 1; i < 1000; i++ {
    if i % 3 == 0 || i % 5 == 0 {
      sum += i
    }
  }
  fmt.Println(sum)
}

func init() {
  addProb("1", func() { runP1() })
}
