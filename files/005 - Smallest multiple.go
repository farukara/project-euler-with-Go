package main

import (
    "fmt"
)

func main() {
    i, target := 20, 1
    found := false
    for !found {
        for j:=2; j<=20; j++ {
            if i%j == 0 {
                if j==20 {
                    target = i
                    found = true
                }
                continue
            } else {
                break
            }
        }
        i += 20
    }
    fmt.Println(target)
}
