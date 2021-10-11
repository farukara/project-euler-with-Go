package main

import (
    "fmt"
    "math"
    "time"
)

func isPrime(n int) bool {
    switch {
    case n<2:
        return false
    case n==2:
        return true
    case n==3:
        return true
    case n==5:
        return true
    case n==6:
        return false
    case n%2 == 0:
        return false
    case n%3 == 0:
        return false
    case n%5 == 0:
        return false
    case n >6:
        for i:=7; float64(i)<math.Sqrt(float64(n)) +1; i+=2 {
            if n%i==0 {
                return false
            }
        }
    }
    return true
}

func make_sieve (size int) [2_000_001]bool {
    var sieve[2_000_001]bool
    sieve[0], sieve[1] = true, true
    for i:=1; i<1_000_000; i++ {
        if !sieve[i] {
            if isPrime(i) {
                j:=2*i
                for {
                    sieve[j] = true
                    j +=i
                    if j > size {break}
                }
            }
        }
    }
    return sieve
}
func main() {
    start := time.Now()
    total := 0
    sieve := make_sieve(2_000_000)
    for i,item := range sieve {
        if !item {
            total +=i
        }
    }
    fmt.Println(total)
    fmt.Printf("Time: %.5f\n", time.Since(start).Seconds())
}
