package main

import (
	"fmt"
)

func runP4() {
	largest := 0
	largest_i := 0
	largest_j := 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= i; j-- {
			val := i * j
			if isPalindrome(val) {
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

func init() {
	addProb("4", func() { runP4() })
}
