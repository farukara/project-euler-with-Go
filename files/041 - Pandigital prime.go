package main
//project euler - 041 - Pandigital prime

import (
    "fmt"
    "math"
    "time"
    "sort"
)

func is_prime (n int) bool {
    if n%2 == 0 || n%3==0 || n%5==0 {return false}
    for i:=6;i<int(math.Sqrt(float64(n)))+1;i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}


func Permutator (n int) [][]int {
    var numbers []int
    var arrows []int //0 left, 1 rigth
    var mobile []int
    var perms [][]int
    var dummy []int

    //create numbers
    for i:=1;i<=n;i++ {
        numbers = append(numbers, i)
        arrows = append(arrows, 0)
    }

    for {
        mobile = nil
        //find the mobile
        for i,number := range numbers {
            if arrows[i] == 0 {
                if i>0 && numbers[i-1] < number { 
                    mobile = append(mobile, number)
                }
            } else {
                if i<n-1 && numbers[i+1] < number {
                    mobile = append(mobile, number)
                }
            }
        }

        //append perms to the slice
        dummy = nil
        for _,number := range numbers {
            dummy = append(dummy, number)
        }
        perms = append(perms, dummy)
        if len(mobile) == 0 {return perms}

        //find max mobile number
        maxn := 0
        for _,number := range mobile {
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
    start := time.Now()
    found := false
    digits := 9
    for !found {
        combsSlice := Permutator(digits)
        var combs []int
        for _,comb := range combsSlice {
            number := 0
            for i:=len(comb);i>=1;i-- {
                number += comb[i-1] * int(math.Pow10(i-1))
            }
            combs = append(combs, number)
        }
        sort.Sort(sort.Reverse(sort.IntSlice(combs)))
        for i:=0;i<len(combs);i++ {
            if is_prime(combs[i]) {
                fmt.Println(combs[i])
                found = true
                break
            }
        }
        digits--
    }
    fmt.Println("Time:", time.Since(start).Seconds(), "seconds")
}
