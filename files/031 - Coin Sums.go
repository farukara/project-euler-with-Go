//In the United Kingdom the currency is made up of pound (£) and pence (p). There are eight coins in general circulation:

//1p, 2p, 5p, 10p, 20p, 50p, £1 (100p), and £2 (200p).
//It is possible to make £2 in the following way:

//1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
//How many different ways can £2 be made using any number of coins?

package main

import (
    "fmt"
    "time"
)

func main() {
    start := time.Now()
    target := 200
    coins := []int{1,2,5,10,20,50,100,200}
    table := make([][]int, len(coins)+1)
    for i:=0; i<len(coins)+1;i++ {
        table[i] = append(table[i], 1)
    }
    for i:=1;i<target+1;i++ {
        table[0] = append(table[0], 0)
    }
    for i:=1;i<len(coins)+1;i++ {
        for j:=1;j<target+1; j++ {
            if j >= coins[i-1] {
                table[i] = append(table[i], table[i-1][j]+table[i][j-coins[i-1]])
            } else {
                if j-coins[i-1] < 0 {
                    table[i] = append(table[i], table[i-1][j])
                } else {
                table[i] = append(table[i], table[i][j-coins[i-1]] + table[i-1][j])
                }
            }
        }
    }
    fmt.Println("Time:", time.Since(start).Nanoseconds())
    fmt.Println(table[len(coins)][target])
    // fmt.Println(table)
}
