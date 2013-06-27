package main

import "testing"

func TestP4reverseStr(t *testing.T) {
  val := p4reverseStr("andrew")
  if val != "werdna" {
    t.Errorf("The reverse of 'andrew' is not %s", val)
  }
}

func TestP4isPalindrome(t *testing.T) {
  val := p4isPalindrome(200)
  if val {
    t.Errorf("200 isn't a palindrome")
  }

  val = p4isPalindrome(2002)
  if !val {
    t.Errorf("2002 is a palindrome")
  }
}
