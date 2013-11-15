package main

/*
http://projecteuler.net/problem=9

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a**2 + b**2 = c**2
For example, 3**2 + 4**2 = 9 + 16 = 25 = 5**2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

import "fmt"

func runP9() {
  for c := 1000; c > 0; c-- {
    for b := c-1; b > 0; b-- {
      for a := b-1; a > 0; a-- {
        if a < b && b < c {
          if a + b + c == 1000 {
            if a*a + b*b == c*c {
              fmt.Printf("%d**2 + %d**2 = %d**2\n", a, b, c)
              fmt.Printf("%d * %d * %d = %d\n", a, b, c, a*b*c)
            }
          }
        }
      }
    }
  }
}

func init() {
  addProb("9", func() { runP9() })
}
