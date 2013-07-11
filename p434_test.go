package main

import "testing"

func TestP434countRigid(t *testing.T) {
  val := p434countRigid(2,3)
  if val != 19 {
    t.Errorf("countRigid result for (2,3) is %d, not 19", val)
  }
  val = p434countRigid(5,5)
  if val != 23679901 {
    t.Errorf("countRigid result for (5,5) is %d, not 23679901", val)
  }
}
