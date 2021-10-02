package main

import "fmt"

func main() {
    total := 0
    prev := 1
    current := 2
    temp := 0
    for current<4_000_000 {
        if current%2 == 0 {
            total += current
        }
        temp = prev
        prev = current
        current = prev + temp
    }
    fmt.Println(total)
}
