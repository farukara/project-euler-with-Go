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
    for i, _ := range str {
        total += strconv.Atoi(string(str[i]))

        // fmt.Printf("type %v: %T\n", str[i], str[i])
        // total += strconv.Atoi(string(letter))
        // fmt.Println(strconv.ParseInt(letter, 10, 64))
    }
    fmt.Println(number)
    fmt.Printf("%T", number)
    fmt.Println(str)
    fmt.Printf("%T", str)
    fmt.Println(total)
    fmt.Printf("Time: %.3f", time.Since(start).Seconds())
    
}
