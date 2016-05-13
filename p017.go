package main

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

var (
	p017Groups = []string{
		"",
		"thousand",
		"million",
		"billion",
		"trillion",
		"quadrillion",
		"sextillion",
		"septillion",
		"octillion",
		"nonillion",
		"decillion",
	}

	p017Ones = []string{
		"",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	p017Teens = []string{
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}

	p017Tens = []string{
		"",
		"",
		"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}
)

func p017NumberInWords(n int) string {
	tens := []int{}
	for n > 0 {
		tens = append(tens, n%10)
		n /= 10
	}

	result := []string{}

	// tens in now in reverse order
	for idx := 0; idx <= len(tens)/3; idx++ {
		low := idx * 3
		high := int(math.Min(float64(idx*3+3), float64(len(tens))))
		group := tens[low:high]
		if len(group) > 0 {
			hundreds := 0
			for i := 0; i < len(group); i++ {
				v := group[i]
				hundreds += v * int(math.Pow(10, float64(i)))
			}
			result = append([]string{p017HundredsWords(hundreds), p017Groups[idx]}, result...)
		}
	}

	return p017CleanResult(result)
}

func p017HundredsWords(n int) string {
	switch {
	case 0 <= n && n < 10:
		return p017Ones[n]
	case 10 <= n && n < 20:
		return p017Teens[n%10]
	case 20 <= n && n < 100:
		ones := p017Ones[n%10]
		if ones == "" {
			return p017Tens[n/10]
		} else {
			return p017Tens[n/10] + "-" + ones
		}
	case 100 <= n && n < 1000:
		hund := p017HundredsWords(n % 100)
		if hund == "" {
			return p017Ones[n/100] + " hundred"
		} else {
			return p017Ones[n/100] + " hundred and " + hund
		}
	default:
		panic(fmt.Sprintf("p017HundredsWords(%d)", n))
	}
}

func p017NumberOfLetters(s string) int {
	ns := strings.Replace(s, "-", "", -1)
	ns = strings.Replace(ns, "-", "", -1)
	ns = strings.Replace(ns, " ", "", -1)
	return utf8.RuneCountInString(ns)
}

func p017SumOfNumberOfLetters(n int) int {
	if n < 1 {
		panic("Come on, dude.")
	}

	result := 0
	for i := 1; i <= n; i++ {
		word := p017NumberInWords(i)
		nl := p017NumberOfLetters(word)
		result += nl
	}

	return result
}

func p017CleanResult(result []string) string {
	return strings.Replace(strings.TrimSpace(
		strings.Join(result, " "),
	), "  ", " ", -1)
}

/**
 * Problem #17
 * https://projecteuler.net/problem=17
 *
 * Number letter counts
 *
 * If the numbers 1 to 5 are written out in words: one, two, three, four, five,
 * then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
 *
 * If all the numbers from 1 to 1000 (one thousand) inclusive were written out
 * in words, how many letters would be used?
 *
 * NOTE: Do not count spaces or hyphens.
 * For example, 342 (three hundred and forty-two) contains 23 letters and 115
 * (one hundred and fifteen) contains 20 letters. The use of "and" when writing
 * out numbers is in compliance with British usage.
 */
func runP17() {
	fmt.Println(p017SumOfNumberOfLetters(1000))
}

func init() {
	addProb("17", func() { runP17() })
}
