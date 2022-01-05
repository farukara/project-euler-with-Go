package main
//project euler - 024 - Lexicographic permutations

import (
    "fmt"
    "sort"
    "math"
)
func Permutator (n int) []int {
//generates permutations of 0 thorough n digits
    var numbers []int
    var arrows []int //0 left, 1 rigth
    var movable []int
    var perms []int
    var dummy int
    var j float64

    //create numbers
    for i:=1;i<=n;i++ {
        numbers = append(numbers, i)
        arrows = append(arrows, 0)
    }

    for {
        movable = nil
        //find the movable
        for i,number := range numbers {
            if arrows[i] == 0 {
                if i>0 && numbers[i-1] < number { 
                    movable = append(movable, number)
                }
            } else {
                if i<n-1 && numbers[i+1] < number {
                    movable = append(movable, number)
                }
            }
        }

        //append perms to the slice
        dummy = 0
        // for _,number := range numbers {
        j=0
        for i:=len(numbers)-1; i>=0; i-- {
            dummy += int(float64(numbers[i]) * math.Pow(10, j))
            j++
        }
        // dummy = append(dummy, number)
        // }
        // fmt.Println(dummy)
        perms = append(perms, dummy)
        if len(movable) == 0 {return perms}

        //find max movable number
        maxn := 0
        for _,number := range movable {
            if number > maxn {
                maxn = number
            }
        }

        //swap numbers and arrows
        for i := range numbers {
            if numbers[i] == maxn {
                if arrows[i] ==0 {
                    numbers[i], numbers[i-1] = numbers[i-1], numbers[i]
                    arrows[i], arrows[i-1] = arrows[i-1], arrows[i]
                    break
                } else {
                    numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
                    arrows[i], arrows[i+1] = arrows[i+1], arrows[i]
                    break
                }
            }
        }
        //flip arrows direction
        for i,number := range numbers {
            if number>maxn {
                arrows[i] ^= 1
            }
        }
    }
}

func main() {
    perms := Permutator(3)
    sort.Ints(perms)
    fmt.Println(perms)
    fmt.Println(len(perms))
}
