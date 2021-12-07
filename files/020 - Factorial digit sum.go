package main
//Project Euler - 020 - Factorial digit sum

import (
    "fmt"
    "math/big"
    "strconv"
)

func main() {
    fib := big.NewInt(1)
    limit := big.NewInt(100)
    step := big.NewInt(1)
    
    for i:=big.NewInt(1);i.Cmp(limit)<0;i.Add(i, step) {
        fib = fib.Mul(fib, i)
    }

    total := 0
    fibs := fib.String()
    for _, digit := range fibs {
        str, _ := strconv.Atoi(string(digit))
        total += int(str)
    }

    fmt.Println(total)
}
