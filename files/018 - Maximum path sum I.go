package main
//project euler - 018 - Maximum path sum I

import (
    "fmt"
    "strings"
    "strconv"
    "time"
)
func main() {
    start := time.Now()
    str := `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`
    var str1 []string
    str1 = strings.Split(str, "\n")
    var tris [][]string
    for _,s := range str1 {
        tris = append(tris, strings.Split(s, " "))
    }
    var tri [][]int
    var temp []int
    for _,line := range tris {
        temp = nil
        for _,s := range line {
            d,_ := strconv.Atoi(s)
            temp = append(temp, int(d)) 
        }
        tri = append(tri, temp)

    }
    var sum [][]int
    sum = append(sum, tri[0])
    for i:=1;i<len(tri);i++ {
        temp = nil

        //getting first element in each row
        temp = append(temp, tri[i][0] + sum[i-1][0])

        //getting middle elements for each row
        for j:=1;j<len(tri[i])-1;j++ {
            if (tri[i][j] + sum[i-1][j-1]) > (tri[i][j] + sum[i-1][j]) {
                temp = append(temp, tri[i][j] + sum[i-1][j-1])
            } else {
                temp = append(temp, tri[i][j] + sum[i-1][j])
            }
        }

        //getting last element in each row
        temp = append(temp, tri[i][len(tri[i])-1] + sum[i-1][len(sum[i-1])-1])
        sum = append(sum, temp)
    }
    //getting the max of the last row
    var max int
    for _,n := range sum[len(sum)-1] {
        if n>max { max = n}
    }
    fmt.Println(max)
    fmt.Println("time:", time.Since(start).Microseconds())
}
