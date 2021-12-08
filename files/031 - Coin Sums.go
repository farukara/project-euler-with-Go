package main
//Project Euler - 031 - Coin Sums

import (
    "fmt"
    "time"
)

func main() {
    start := time.Now()
    ultimateTarget := 200
    coins := [8]int{1,2,5,10,20,50,100,200}
    table := make([][]int, len(coins)+1)

    for i:=0; i<len(coins)+1;i++ {
        table[i] = append(table[i], 1)
    }

    for i:=1;i<ultimateTarget+1;i++ {
        table[0] = append(table[0], 0)
    }

    for i:=1;i<len(coins)+1;i++ {
        for currentTarget:=1;currentTarget<ultimateTarget+1; currentTarget++ {
            currentCoin := coins[i-1]

            if currentTarget-currentCoin < 0 {
                table[i] = append(table[i], table[i-1][currentTarget])
            } else {
            table[i] = append(table[i], table[i-1][currentTarget] + table[i][currentTarget-currentCoin])
            }
        }
    }
    fmt.Println("Time:", time.Since(start).Nanoseconds())
    fmt.Println(table[len(coins)][ultimateTarget])
}
