package main

import (
    "fmt"
    "math"
    "strconv"
)

func main() {
    var palindroms []int  // slice of palindroms
    for i:=999; i>99; i-- { //outer loop
        for j:=999; j>99; j-- { //inner loop
            prod := i*j //find the product of 2 numbers
            str_prod := strconv.Itoa(prod) //string version of production for easy access
            half_way := int(math.Floor(float64(len(str_prod))/2)) //we only need to check first half with second
            for i:=0; i<half_way+1; i++ {
                if str_prod[i] != str_prod[len(str_prod)-1-i] {
                    break //if any inequality no need to check further
                } else if i == half_way { //if made it to the end then it's a palindrome
                palindroms = append(palindroms, prod) // add to candidate list
                }
            }
        }
    }

    largest := palindroms[0] //loop to find the largest
    for _, element := range palindroms {
        if largest < element {largest = element}
    }
    fmt.Println(math.Max(palindroms))
}
