package main

import (
    "fmt"
)

func main() {
    first, second, total := 0, 0, 0
    for i:=1; i<=100; i++ {
        first += i*i
        total += i
    }
    second = total*total
    fmt.Println("Difference is: ", second-first)
}
