package main

var p016map = map[int]int{
	1:  1,
	2:  2,
	3:  4,
	4:  -1,
	5:  -2,
	6:  5,
	7:  3,
	8:  -5,
	9:  -1,
	10: 7,
	11: 5,
}

func p016SumOfPowerDigit(n int) int {
	if n < 0 {
		panic("This doesn't work for negative numbers!")
	}
	if n == 0 {
		return 1
	}
	return p016map[(n-1)%12] + p016SumOfPowerDigit(n-1)
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
