package main

import "testing"

func TestReverseStr(t *testing.T) {
	val := reverseStr("andrew")
	if val != "werdna" {
		t.Errorf("The reverse of 'andrew' is not %s", val)
	}
}

func TestIsPalindrome(t *testing.T) {
	val := isPalindrome(200)
	if val {
		t.Errorf("200 isn't a palindrome")
	}

	val = isPalindrome(2002)
	if !val {
		t.Errorf("2002 is a palindrome")
	}
}

func TestJoinInts(t *testing.T) {
	val := joinInts([]int{8, 4, 3, 2, 7})
	if val != "84327" {
		t.Errorf("Joined version of slice is not %s", val)
	}
}
