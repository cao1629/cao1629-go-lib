package main

import (
    "fmt"
    "sort"
)

func main1() {
    x := []int{3, 1, 2}
    y := x[:]
    sort.Ints(x)
    fmt.Println(y)
}
