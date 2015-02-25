package main

/*
 * Problem #14 - Longest Collatz sequence
 * https://projecteuler.net/problem=14
 *
 * The following iterative sequence is defined for the set of positive integers:
 *
 * n → n/2 (n is even)
 * n → 3n + 1 (n is odd)
 *
 * Using the rule above and starting with 13, we generate the following sequence:
 *
 * 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
 *
 * It can be seen that this sequence (starting at 13 and finishing at 1)
 * contains 10 terms. Although it has not been proved yet (Collatz Problem),
 * it is thought that all starting numbers finish at 1.
 *
 * Which starting number, under one million, produces the longest chain?
 *
 * NOTE: Once the chain starts the terms are allowed to go above one million.
 */

var p014SeqLen = make(map[int]int)

func p014GetSeqLen(n int) int {
	if val, ok := p014SeqLen[n]; ok {
		return val
	} else if n == 1 {
		return 1
	} else {
		return 1 + p014GetSeqLen(p014NextN(n))
	}
}

func p014NextN(n int) int {
	if n&1 == 1 {
		// n is odd
		return 3*n + 1
	} else {
		// n is even
		return n / 2
	}
}

// Solution: 837799 produces sequence of length 525
func runP14() {
	n := 0
	l := 0
	for i := 1; i < 1000000; i++ {
		v := p014GetSeqLen(i)
		if v > l {
			n = i
			l = v
		}
	}
	println(n, "produces sequence of length", l)
}

func init() {
	addProb("14", func() { runP14() })
}
