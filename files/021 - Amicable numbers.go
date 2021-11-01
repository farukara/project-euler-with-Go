package main
//project euler 021 - Amicable numbers

import (
    "fmt"
    "time"
    "math"
)

//returns sum of proper divisors that are less than n
func divSum (n int) int {
    total := 0
    for i:=2; i<int(math.Sqrt(float64(n))+2); i++ {
        if n%i == 0 {
            total += i
            if i*i != n {
                total += n/i
            }
        }
    }
    return total+1 //adding 1 because iteration started from 2 to avoid i*i logic failure
}

func main() {
    start := time.Now()
    total := 0
    sum := 0 
    var alreadyCounted [10_000] bool
    for i:=1; i<10_000; i++ {
        sum = divSum(i) 
        if !alreadyCounted[i] && divSum(sum) == i && sum != i {
            total += i + sum
            alreadyCounted[i] = true
            alreadyCounted[sum] = true
        }
    }
    println("Total:", total)
    fmt.Printf("Time: %.3f", time.Since(start).Seconds())
}
