package main

import "testing"

func TestP017NumberInWords(t *testing.T) {
	for _, tt := range []struct {
		in  int
		out string
	}{
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "four"},
		{5, "five"},
		{6, "six"},
		{7, "seven"},
		{8, "eight"},
		{9, "nine"},
		{10, "ten"},
		{342, "three hundred and forty-two"},
		{115, "one hundred and fifteen"},
		{900, "nine hundred"},
		{999, "nine hundred and ninety-nine"},
		{1000, "one thousand"},
		{1001, "one thousand one"},
		{2342, "two thousand three hundred and forty-two"},
		{1234567001890, "one trillion two hundred and thirty-four billion five hundred and sixty-seven million one thousand eight hundred and ninety"},
	} {
		out := p017NumberInWords(tt.in)
		if out != tt.out {
			t.Errorf("p017NumberInWords(%d) => %s, want %s", tt.in, out, tt.out)
		}
	}
}

func TestP017NumberOfLetters(t *testing.T) {
	for _, tt := range []struct {
		in  string
		out int
	}{
		{"one", 3},
		{"two", 3},
		{"three", 5},
		{"four", 4},
		{"five", 4},
		{"three hundred and forty-two", 23},
		{"one hundred and fifteen", 20},
	} {
		out := p017NumberOfLetters(tt.in)
		if out != tt.out {
			t.Errorf("p017NumberOfLetters(%s) => %d, want %d", tt.in, out, tt.out)
		}
	}
}

func TestP017SumOfNumberOfLetters(t *testing.T) {
	for _, tt := range []struct {
		in  int
		out int
	}{
		{1, 3},  // one
		{2, 6},  // one two
		{3, 11}, // one two three
		{4, 15}, // one two three four
		{5, 19}, // one two three four five
	} {
		out := p017SumOfNumberOfLetters(tt.in)
		if out != tt.out {
			t.Errorf("p017SumOfNumberOfLetters(%d) => %d, want %d", tt.in, out, tt.out)
		}
	}
}
