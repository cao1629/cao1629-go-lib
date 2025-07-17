package main

import (
    "fmt"
    "time"
)

func main() {
    var t *time.Ticker

    select {
    case <-t.C:
        fmt.Println("tick")
    }
}
