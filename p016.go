package main

import "math"

func p016SumOfPowerDigit(n int) int {
	if n < 0 {
		panic("p016SumOfPowerDigit does not work for negative numbers!")
	}
	if n < 4 {
		return int(math.Pow(float64(2.0), float64(n)))
	} else if isEven(n) {
		// even
		return 2 * p016SumOfPowerDigit(n/2)
	} else {
		// odd
		return p016SumOfPowerDigit(1) + p016SumOfPowerDigit(n-1)
	}
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
