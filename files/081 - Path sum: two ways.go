package main
//project euler - 081 - Path sum: two ways

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
    // routes := [][]int{{131,673,234,103,18}, {201,96,342,965,150}, {630,803,746,422,111}, {537,699,497,121,956}, {805,732,524,37,331}}
    routes := make([][]int, 80)
    file,err := os.ReadFile("p081_matrix.txt")
    if err !=nil {
        fmt.Println("Error reading file")
    }

    lines := strings.Split(strings.Trim(string(file), "\n"), "\n")//[:80]
    for i,line := range lines {
        words := strings.Split(line, ",")
        for _,word := range words {
            number,_ := strconv.Atoi(word)
            routes[i] = append(routes[i], number)
        }
    }

    table :=make([][]int, len(routes))
    table[0] = append(table[0], routes[0][0])
    for i:=1;i<len(routes);i++ {
        table[i] = append(table[i], table[i-1][0]+routes[i][0])
    }

    for j:=1;j<len(routes[0]);j++ {
        table[0] = append(table[0], table[0][j-1]+routes[0][j])
    }

    for i:=1;i<len(routes);i++ {
        for j:=1;j<len(routes[0]);j++ {
            if table[i-1][j] < table[i][j-1] {
                table[i] = append(table[i], table[i-1][j] + routes[i][j])
            } else {
                table[i] = append(table[i], table[i][j-1] + routes[i][j])
            }
        }
    }
    fmt.Println(table[len(routes)-1][len(routes[0])-1])
}
