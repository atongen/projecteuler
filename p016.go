package main

import (
	"math/big"
	"strconv"
	"strings"
)

func p016BigPowerOf2(n int) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	} else if n == 1 {
		return big.NewInt(2)
	} else {
		r := new(big.Int)
		return r.Mul(big.NewInt(2), p016BigPowerOf2(n-1))
	}
}

func p016SumOfPowerDigit(n int) int {
	str := p016BigPowerOf2(n).String()
	digits := strings.Split(str, "")
	result := 0
	for _, d := range digits {
		i, _ := strconv.Atoi(d)
		result += i
	}
	return result
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
