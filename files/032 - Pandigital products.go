package main

import (
    "fmt"
    "strconv"
    "strings"
    "sort"
)

func main() {
    var candidates []int
    deck := "123456789"
    for i:=1;i<2_000;i++ {
        for j:=1;j<2_000;j++ {
            product:=i*j
            prostr := strconv.Itoa(product)
            istr := strconv.Itoa(i)
            jstr := strconv.Itoa(j)
            wholestr:=prostr+istr+jstr
            if len(wholestr) == 9 {
                includeAll:=true
                for _,digit:= range deck {
                    if !strings.Contains(wholestr, string(digit)) {
                        includeAll = false
                        break
                    }
                }
                if includeAll {
                    candidates = append(candidates, product)
                }
            }
        }
    }
    sort.Ints(candidates)
    //zeroizing repeating numbers with bit manipulation
    for i,number := range candidates {
        for j:=i+1;j<len(candidates);j++ {
            if number^candidates[j] == 0 {
                candidates[j] = 0
            }
        }
    }
    total := 0
    for _,number := range(candidates) {
        total += number
    }
    fmt.Println(total)
}
