package main

import (
    "fmt"
)

func fibonacciGenerator(c chan int, done chan struct{}) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-done:
            fmt.Println("Stop")
            return
        }
    }
}

func main232323() {
    c := make(chan int)
    done := make(chan struct{})

    go fibonacciGenerator(c, done)

    for i := 0; i < 10; i++ {
        fmt.Println(<-c)
    }

    // Signal the goroutine to stop
    done <- struct{}{}

    // Wait a moment to ensure the goroutine has stopped
    fmt.Println("Main function completed")

}
