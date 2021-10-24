package main

import (
    "fmt"
)

func main() {
    below20 := "onetwothreefourfivesixseveneightnineteneleventwelvethirteenfourteenfifteensixteenseventeeneighteennineteen"
    below100 := "twentythirtyfourtyfiftysixtyseventyeightyninety"
    below10 := "onetwothreefourfivesixsevennine"
    hundredands := "hundredand"
    thousand := "thousand"
    belowhundred := len(below20) + len(below10)*8 + len(below100)
    total := len(hundredands)*900 + belowhundred*9 + len(below10)*9 + len(thousand)

    fmt.Println(total)
}
