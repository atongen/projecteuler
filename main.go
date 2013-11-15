package main

import (
  "fmt"
  "flag"
  "log"
)

type probFunc func()
var problems map[string]probFunc = make(map[string]probFunc)

func addProb(key string, prob probFunc) {
  if _, ok := problems[key]; ok {
    log.Fatalf("Problem key exists: %s\n", key)
  }
  problems[key] = prob
}

func main() {
  flag.Parse()

  arg := flag.Arg(0)
  if myProb, ok := problems[arg]; ok {
    myProb()
  } else {
    fmt.Printf("Problem key not found: %s", arg)
  }
}
