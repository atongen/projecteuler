package main

import (
  "strconv"
  "fmt"
)

func p4reverseStr(s string) string {
  runes := []rune(s)
  for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
  }
  return string(runes)
}

func p4isPalindrome(n int) bool {
  v := strconv.FormatInt(int64(n), 10)
  return v == p4reverseStr(v)
}

func runP4() {
  largest := 0
  largest_i := 0
  largest_j := 0
  for i := 999; i > 0; i-- {
    for j := 999; j >= i; j-- {
      val := i*j
      if p4isPalindrome(val) {
        if val > largest {
          largest = val
          largest_i = i
          largest_j = j
          fmt.Printf("%d, %d, %d\n", largest_i, largest_j, largest)
        }
      }
    }
  }
}
