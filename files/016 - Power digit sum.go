package main
//project euler 016 - Power digit sum

import (
    "fmt"
    "time"
    "math"
    "strconv"
)

func main() {
    start := time.Now()
    number := strconv.FormatFloat(math.Pow(2, 1_000), 'E', 301, 64)
    str := number[:1]+number[2:len(number)-5]
    total := 0
    for _, digit := range str {
        total += int(digit)-int('0')
    }
    fmt.Println("Total:", total)
    fmt.Printf("Time: %.6f", time.Since(start).Seconds())
    
}
