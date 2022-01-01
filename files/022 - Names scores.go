package main
//project euler - 022 - Names scores

import (
    "fmt"
    "os"
    "bytes"
    "sort"
)

func main() {
    raw, _ := os.ReadFile("p022_names.txt")
    data := bytes.Split(raw, []byte(","))
    score := 0
    var names []string
    for _,name := range data {
        names = append(names, string(name[1:len(name)-1]))
    }
    sort.Strings(names)
    for i := range names {
        letScore := 0
        for _,letter := range names[i] {
            letScore += int(letter) - 64
        }
        score += (letScore * (i+1))
    }
    fmt.Println(score)
}
