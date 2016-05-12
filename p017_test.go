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
		{342, "three hundred and forty-two"},
		{115, "one hundred and fifteen"},
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
		{"three", 3},
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
		{1, 3},
		{2, 6},
		{3, 9},
		{4, 13},
		{5, 17},
	} {
		out := p017SumOfNumberOfLetters(tt.in)
		if out != tt.out {
			t.Errorf("p017SumOfNumberOfLetters(%d) => %d, want %d", tt.in, out, tt.out)
		}
	}
}
