package main
//project euler - 079 - Passcode derivation

import (
    "fmt"
    "os"
    "bytes"
    "strconv"
)

func main() {
    var result []byte
    raw, _ := os.ReadFile("p079_keylog.txt")
    numbers := bytes.Split(raw, []byte("\n"))
    numbers = numbers[:len(numbers)-1]
    fmt.Println(numbers)
    for k:=0;k<10;k++ {
        for i:=0;i<10;i++ {
            match := true
            for _,number := range numbers {
                // fmt.Println(strconv.Itoa(i), string(number[1]), string(number[2]))
                if strconv.Itoa(i) == string(number[1]) || strconv.Itoa(i) == string(number[2]) {
                    match = false
                    break
                } 
            }
            if match {
                for _,number := range numbers {
                    present := false
                    for j:=0;j<3;j++ {
                        if strconv.Itoa(i) == string(number[j]) {
                            found := false
                            for _,n := range result {
                                if n==number[j] {
                                    found = true
                                    break
                                }
                                if !found {
                                    result = append(result, number[j])
                                    fmt.Println(i, number, result)
                                    present = true
                                    break
                                }
                            }
                        }
                    }
                    if present {
                        break
                    }
                }
            }
        }
    }
    fmt.Println(result)
}
