package main
//project euler - 039 - Integer right triangles

import (
    "fmt"
    "time"
)

func main() {
    start := time.Now()
    counter := 0
    max := 0
    maxp := 0 
    for p:=2;p<=1_000;p=p+2 {
        counter = 0
        for i:=1;i<=p/3+1;i++{
            for j:=i;j<=(p-i)/2+1;j++ {
                k:= (p-i)-j
                if k*k == i*i + j*j {
                    counter +=1
                }
            }
        }
        if counter>max {
            max = counter
            maxp = p
        }
    }
    fmt.Println(maxp, max)
    fmt.Println("Time:", time.Since(start).Milliseconds(), "milliseconds")
}
