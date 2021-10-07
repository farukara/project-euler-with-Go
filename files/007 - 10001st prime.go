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
    case n%2 == 0:
        return false
    case n%3 == 0:
        return false
    case n%5 == 0:
        return false
    case n >5:
        for i:=6; float64(i)<math.Sqrt(float64(n)) +1; i++ {
            if n%i==0 {
                return false
            }
        }
    }
    return true
}

func main() {
    start := time.Now()
    i, counter :=1, 0
    for counter!=10_001{
        if isPrime(i) {
            counter++
        }
        i++
    }
    fmt.Println(i-1) //-1 for the extraa i++ in the last loop
    fmt.Printf("Time: %.3f", time.Since(start).Seconds())
}
