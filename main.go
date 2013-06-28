package main

import (
  "fmt"
  "flag"
  "errors"
)

func runProblem(num int) (err error) {
  switch num {
  case 1:
    runP1()
    return nil
  case 2:
    runP2()
    return nil
  case 3:
    runP3()
    return nil
  case 4:
    runP4()
    return nil
  case 5:
    runP5()
    return nil
  }
  return errors.New("No problem!")
}

var num *int = flag.Int("n", 0, "The problem number to run")

func main() {
  flag.Parse();
  if err := runProblem(*num); err != nil {
    fmt.Println(err)
  }
}
