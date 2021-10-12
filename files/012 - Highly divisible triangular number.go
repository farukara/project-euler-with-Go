package main

import (
    "fmt"
    "math"
    "time"
)

func main() {
    start := time.Now()
    current, divisors := 0, 0
    j:=1
    found := false
    for !found {
        current = current + j
        divisors = 0
        for i:=1; i<int(math.Ceil(math.Pow(float64(current), 0.5))); i++ {
            if current%i == 0 {
                divisors = divisors + 2
                if divisors > 500 {
                    found = true
                    break
                }
            }
        }
        j++
    }
    fmt.Println(divisors, current)
    fmt.Println("Time:", time.Since(start).Seconds())
}
