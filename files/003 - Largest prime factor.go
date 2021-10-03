package main

import (
    "fmt"
    "math"
    )

func is_prime(number int)bool {
    for i:=2; float64(i) < math.Sqrt(float64(number))+1; i++ {
        if number%i==0.0 {return false}
    }
    return true
}

func main() {
    target := 600851475143
    for i:=2; i<target; i++ {
        if target%i == 0 && is_prime(target/i) {
            fmt.Println("Largest factor is:", target/i)
            break
        }
    }
}
