package main
//project euler 014 - longest collatz sequence

import (
    "fmt"
    "time"
)

func find_collatz_len(n int, memo map[int]int) int {
    counter := 1
    initital_value := n
    for n != 1 {
        counter++
        if n < initital_value {
            counter = counter + memo[n]
            break
        }
        if n%2 == 0 {
            n = n/2
        } else {
            n = n*3+1
        }
    }
    memo[initital_value] = counter
    return counter
}

func main() {
    start := time.Now()
    max_length := 0
    number := 0
    current_length := 0
    memo := make(map[int]int, 1_000_000)
    for i:=1; i<1_000_000; i++ {
        current_length = find_collatz_len(i, memo)
        if current_length > max_length {
            max_length = current_length 
            number = i
        }
    }
    fmt.Println(max_length, number)
    fmt.Printf("Time: %.3f", time.Since(start).Seconds())
}
