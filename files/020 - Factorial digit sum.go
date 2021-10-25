package main

import (
    "fmt"
)

func main() {
    fib := 1
    for i:=1;i<=100;i++ {
        fib *= i
    }
    fmt.Println(fib) //shows 0 because of the damn formatting same problem with 16
}
