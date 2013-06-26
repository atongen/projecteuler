package main

import (
  "fmt"
)

func p3nextPrime(n uint64) uint64 {
  var i uint64 = n+1
  for ;; i++ {
    if p3isPrime(i) {
      return i
    }
  }
}

func p3isPrime(n uint64) bool {
  if n > 2 {
    var i uint64 = 2
    for ; i < n/2; i++ {
      if n % i == 0 {
        return false
      }
    }
  }
  return true
}

//require 'mathn'
//#num, factor = 317_584_931_803, 0
//num, factor = 1000, 0
//primes = Prime.new
//while num > 1
//  puts({ num: num, factor: factor }.inspect)
//  factor = primes.next
//  num /= factor while (num % factor).zero?
//end
//
//puts "Largest factor is #{ factor }."

func runP3() {
  var ultimate uint64 = 600851475143
  fmt.Printf("Largest prime factor of %d\n", ultimate)

  var factor uint64 = 1

  for ; ultimate > 1 ; {
    factor = p3nextPrime(factor)
    for ;; {
      if ultimate % factor == 0 {
        ultimate /= factor
      } else {
        break
      }
    }
  }

  fmt.Println(factor)

  //var i uint64 = ultimate / 2
  //if i % 2 == 0 {
  //  i--
  //}
  //for ; i > 2; i -= 2 {
  //  if (i-1) % 1000000 == 0 {
  //    fmt.Printf("%d...\n", i)
  //  }
  //  if ultimate % i == 0 {
  //    if p3isPrime(i) {
  //      fmt.Println(i)
  //      break
  //    }
  //  }
  //}
}
