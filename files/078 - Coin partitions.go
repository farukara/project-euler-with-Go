package main
//project euler - 078 - Coin partitions

import (
    "fmt"
)

func main() {
    // table := make([][]int, 1)
    // table[0] = append(table[0],0)
    table := []int{1,2,3,5,7,11}
    coin := 6
    for  {
        table = append(table, (table[coin-1] + table[coin-3])%1_000_000)
        /* if table[coin] % 1_000_000 == 0 {
            fmt.Println(table[coin], coin)
            break
        } */
        if coin==10 {
            fmt.Print(table)
            break
        }
        coin += 1
    }
        /* table[0] = append(table[0], 0)
        table = append (table,[]int{1})
        for i:=1;i<row;i++ {
            table[i] = append(table[i], table[i-1][row] + table[i][row-1])
            table[row] = append(table[row], table[row-1][i] + table[row][i-1])
        }
        table[row] = append(table[row], table[row-1][row] + table[row][row-1])
        // if table[row][row] == 7 {
        if row == 5 {
            fmt.Println(table, table[row][row])
            break
        }
        row +=1
    } */
}

/* x  0  1  2  3  4  5
0  0  0  0  0  0  0 
1  1  1
2        2
3           3
4              4
5                 7
6                    10 */
