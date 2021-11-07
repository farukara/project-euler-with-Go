package main
//project euler - 030. Digit fifth powers

import (
    "fmt"
    "strconv"
    "math"
    "time"
)

func main() {
    start := time.Now()
    const power float64=5
    var grand_total float64

    //pre-calculate the powers
    var lookup_array [10]float64
    for i:=0; i<10; i++ {
        lookup_array[i] = math.Pow(float64(i), power)
    }

    for i:=2; i<=300_000; i++ {
        str_number := string(strconv.Itoa(i))
        var total float64
        for i:=0; i<len(str_number); i++ {
            digit, _ := strconv.Atoi(string(str_number[i]))
            total += lookup_array[digit]
        }
        if total == float64(i) {
            grand_total += total
        }
    }

    fmt.Println("grand_total: ", grand_total)
    fmt.Println("Time: ", time.Since(start).Seconds())
}
