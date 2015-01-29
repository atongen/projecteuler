package main

import (
	"math"
)

func nextPrime(n uint64) uint64 {
	var i uint64 = n + 1
	for ; ; i++ {
		if isPrime(i) {
			return i
		}
	}
}

func isPrime(n uint64) bool {
	if n > 2 {
		var i uint64 = 2
		max := uint64(math.Sqrt(float64(n)))
		for ; i <= max; i++ {
			if n%i == 0 {
				return false
			}
		}
	}
	return true
}
