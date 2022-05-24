package main
//project euler - 067 - Maximum path sum II

import (
    "fmt"
    "strings"
    "strconv"
    "time"
    "os"
)
func main() {
    start := time.Now()
    f, err := os.ReadFile("p067_triangle.txt")
    if err != nil {fmt.Println("error reading file")}
    str := strings.Trim(string(f), "\n")
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
    for _,n := range sum[len(sum)-1] { //-2 because some how the last line is empty
        if n>max { max = n}
    }
    fmt.Println("max:", max)
    fmt.Println("time:", time.Since(start).Microseconds())
}
