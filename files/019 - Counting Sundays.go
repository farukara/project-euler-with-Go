package main
//project euler - 019 - Counting Sundays

import (
    "fmt"
    "time"
)

func main() {
    start := time.Now()
    date := time.Date(1901, 1, 1, 9, 0, 0, 0, time.UTC)
    final := time.Date(2000, 12, 31, 8, 0, 0, 0, time.UTC)
    i := 0
    for date.Before(final) {
        if date.Weekday() == 0 && date.Day() == 1 {i++}
        date = date.AddDate(0,0,1)
    }
    fmt.Println(i)
    fmt.Println("Time:", time.Since(start).Microseconds(), "microseconds")
}
