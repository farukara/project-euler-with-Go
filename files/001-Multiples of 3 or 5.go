package main

import "fmt"

func main() {
    total := 0
    for i:=1; i < 1_000; i++ {
        if (i%3==0) || (i%5==0) {
            fmt.Println(i)
            total += i
        }
    }
    fmt.Println("total:", total)
}
