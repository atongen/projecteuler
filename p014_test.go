package main

import "testing"

func TestP014NextN(t *testing.T) {
	var data = []struct {
		in  int
		out int
	}{
		{13, 40},
		{40, 20},
		{20, 10},
		{10, 5},
		{5, 16},
		{16, 8},
		{8, 4},
		{4, 2},
		{2, 1},
	}

	for _, tt := range data {
		s := p014NextN(tt.in)
		if s != tt.out {
			t.Error("Expected", tt.out, "but got", s, "for", tt.in)
		}
	}
}

func TestP014GetSeqLen(t *testing.T) {
	var data = []struct {
		in  int
		out int
	}{
		{13, 10},
		{40, 9},
		{20, 8},
		{10, 7},
		{5, 6},
		{16, 5},
		{8, 4},
		{4, 3},
		{2, 2},
		{1, 1},
	}

	for _, tt := range data {
		s := p014GetSeqLen(tt.in)
		if s != tt.out {
			t.Error("Expected", tt.out, "but got", s, "for", tt.in)
		}
	}
}
