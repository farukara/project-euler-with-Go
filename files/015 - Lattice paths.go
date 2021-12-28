package main
//Project Euler - 015 - Lattice paths

import (
    "fmt"
    "time"
)

func main() {
    start := time.Now()
    table := make([][]int, 21)
    for i:=0; i<21; i++ {
        table[0] = append(table[0], 1)
        table[i] = append(table[i], 1)
    }
    for i:=1; i<21; i++ {
        for j:=1; j<21; j++ {
            table[i] = append(table[i], table[i-1][j] + table[i][j-1])
        }
    }
    fmt.Println(table[20][20])
    fmt.Println("Time:", time.Since(start).Microseconds())

}
