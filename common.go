package main

import (
	"math"
	"strconv"
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

func reverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(n int) bool {
	v := strconv.FormatInt(int64(n), 10)
	return v == reverseStr(v)
}

func sumOfSquares(n int) uint64 {
	var i uint64 = 0
	for j := 1; j <= n; j++ {
		i += uint64(j * j)
	}
	return i
}

func squareOfSum(n int) uint64 {
	var sum uint64 = 0
	for i := 1; i <= n; i++ {
		sum += uint64(i)
	}
	return sum * sum
}
