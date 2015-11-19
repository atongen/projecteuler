package main

import (
	"strconv"
	"strings"
)

func p016SumOfPowerDigit(n int) int {
	if n < 0 {
		panic("This doesn't work for negative numbers!")
	}
	if n == 0 {
		return 1
	} else if n == 1 {
		return 2
	} else {
		return 2 * p016SumOfPowerDigit(n-1)
	}
}

func p016SumOfDigits(n int) int {
	s := strconv.Itoa(n)
	ss := strings.Split(s, "")
	sum := 0
	for _, v := range ss {
		i, _ := strconv.Atoi(v)
		sum += i
	}
	return sum
}

/**
 * Problem #16 - Power digit sum
 * https://projecteuler.net/problem=16
 *
 * 2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
 *
 * What is the sum of the digits of the number 2^1000?
 *
 * Solution:
 */
func runP16() {
	println(p016SumOfPowerDigit(1000))
}

func init() {
	addProb("16", func() { runP16() })
}
