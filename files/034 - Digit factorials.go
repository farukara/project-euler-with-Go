package main
//Project Euler - 034 - Digit factorials

import (
    "fmt"
    "time"
)

func main() {
    start := time.Now()
    facts := make(map[int]int)
    facts [0] = 1
    for i:=1;i<=9;i++ {
        facts[i] = facts[i-1] * i
    }

    fmt.Println("upper limit:", 7*facts[9]) //theoretical upper limit
    grandTotal := 0
    for i:=3;i<2_540_160;i++ {
        n := i
        total := 0
        for n>0 {
            digit := n%10
            total += facts[digit]
            n = n/10
        }
        if total == i {
            fmt.Println(i)
            grandTotal += total
        }
    }
    fmt.Println("Grand Total:", grandTotal)
    fmt.Println("Time:", time.Since(start).Milliseconds(), "milliseconds")
}
