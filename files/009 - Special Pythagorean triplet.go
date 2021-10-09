package main

import (
    "fmt"
    "time"
)

func main() {
    prod := 0
    start := time.Now()
    for a:=1; a<=334; a++ { //mathematically a cannot be greater than 1/3 of 1000
        for b:=a+1; b<=501; b++ { //b cannot be greater than 1/2 of 1000
            c := 1000 - a - b
            if a*a+b*b==c*c {
                prod = a*b*c
            }
        }
    }

    fmt.Println("Product is", prod)
    fmt.Printf("Time: %.5f", time.Since(start).Seconds())
}
