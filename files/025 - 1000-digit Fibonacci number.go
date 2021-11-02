package main

import (
    "fmt"
    "time"
    "math/big"
)

func main() {
    start := time.Now()
    fib2 := big.NewInt(1)
    fib1 := big.NewInt(1)
    fib := big.NewInt(2)
    var limit big.Int
    limit.Exp(big.NewInt(10), big.NewInt(999), nil)
    for i:=4;; i++ {
        fib2.Set(fib)
        fib.Add(fib, fib1)
        fib1.Set(fib2)
        if fib.Cmp(&limit) > 0 {
            fmt.Println(i)
            break
        }
    }
    fmt.Printf("Time: %.04f", time.Since(start).Seconds())
}
