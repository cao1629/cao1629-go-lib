package main

import "fmt"

func main544() {
    x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    y := x[2:7]

    x[5] = 100
    fmt.Println(y)
}
