package main
//project euler - 076 - Counting summations

import (
    "fmt"
)

func main() {
    row:=1
    for {
        col:=row
        table := make([][]uint, row+1)
        for i:=0;i<=row;i++ {
            table[i] = append(table[i], 1)
        }
        for i:=1; i<=col;i++ {
            table[0] = append(table[0], 0)
        }
        for i:=1;i<=row;i++ {
            for j:=1;j<=col;j++ {
                if j<i {
                    table[i] = append(table[i], table[i-1][j])
                } else {
                    table[i] = append(table[i], table[i-1][j] + table[i][j-i])
                }
            }
        }
        if table[row][col] % 1_000_000 == 0 {
            fmt.Println(table[row][col], row)
            break
        }
        // if row == 1000 {break}
        row +=1
        // fmt.Println(row)
    }
    // fmt.Println(table[row][col])
    // fmt.Println(table)
}
